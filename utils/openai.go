// Copyright 2023 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"context"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
	"golang.org/x/time/rate"

	"github.com/casbin/casdoc/config"
	. "github.com/casbin/casdoc/logger"
)

var (
	OpenAIClient   *openai.Client
	RequestLimiter *rate.Limiter
	TokenLimiter   *rate.Limiter
)

func init() {
	OpenAIClient = openai.NewClient(config.AuthToken)
	requestLimit := rate.Every(time.Minute / time.Duration(config.RPM))
	tokenLimit := rate.Every(time.Minute / time.Duration(config.TPM))
	RequestLimiter = rate.NewLimiter(requestLimit, config.RPM)
	TokenLimiter = rate.NewLimiter(tokenLimit, config.TPM)
}

func Wait(tokenNumber int) (err error) {
	err = RequestLimiter.Wait(context.Background())
	if err != nil {
		return
	}
	err = TokenLimiter.WaitN(context.Background(), tokenNumber)
	if err != nil {
		return
	}
	return nil
}

// submit a request to gpt
func gpt(req openai.ChatCompletionRequest, c *string) (ans *string, err error) {
	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: *c,
	})
	resp, err := OpenAIClient.CreateChatCompletion(context.Background(), req)

	if err != nil {
		if strings.Contains(err.Error(), "status code: 429") {
			time.Sleep(30 * time.Second)
			resp, err = OpenAIClient.CreateChatCompletion(context.Background(), req)
			if err != nil {
				Logger.Errorf("Failed to get response from gpt: %v", err)
				return
			}
		} else {
			Logger.Errorf("Failed to get response from gpt: %v", err)
		}
		return
	}

	Logger.Info("Prompt tokens: ", resp.Usage.PromptTokens)
	Logger.Info("Completion tokens: ", resp.Usage.CompletionTokens)
	Logger.Info("Total tokens: ", resp.Usage.TotalTokens)
	ans = &resp.Choices[0].Message.Content
	return
}
