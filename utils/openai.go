package utils

import (
	"casdoc/config"
	. "casdoc/logger"
	"context"
	"github.com/sashabaranov/go-openai"
	"golang.org/x/time/rate"
	"strings"
	"time"
)

var OpenAIClient *openai.Client
var RequestLimiter *rate.Limiter
var ToeknLimiter *rate.Limiter

func init() {
	OpenAIClient = openai.NewClient(config.AuthToken)
	requestLimit := rate.Every(time.Minute / time.Duration(config.RPM))
	tokenLimit := rate.Every(time.Minute / time.Duration(config.TPM))
	RequestLimiter = rate.NewLimiter(requestLimit, config.RPM)
	ToeknLimiter = rate.NewLimiter(tokenLimit, config.TPM)
}

func Wait(tokenNumber int) (err error) {
	err = RequestLimiter.Wait(context.Background())
	if err != nil {
		return
	}
	err = ToeknLimiter.WaitN(context.Background(), tokenNumber)
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
