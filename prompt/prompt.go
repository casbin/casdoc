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

package prompt

import "github.com/sashabaranov/go-openai"

// the prompt for polish
const polishPrompt = `You are given an mdx document written in English, and you are tasked with polishing it by correcting typos, improving sentence fluency, and so on, without changing the structure or style of the document. Please only return the polished document.`

var PolishRequest = openai.ChatCompletionRequest{
	Model: openai.GPT3Dot5Turbo,
	Messages: []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: polishPrompt,
		},
	},
}

// the prompt for translate
// 中文
const chinesePrompt = `Translate the given document in mdx format into Chinese. Please be careful not to modify the structure of the document and do not translate technical terms such as Casbin, Casdoor, SSO, Swagger, URL, etc. Only provide the final result.`

var ChineseRequest = openai.ChatCompletionRequest{
	Model: openai.GPT3Dot5Turbo,
	Messages: []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: chinesePrompt,
		},
	},
}

// Français
const frenchPrompt = ``

var FrenchRequest = openai.ChatCompletionRequest{
	Model: openai.GPT3Dot5Turbo,
	Messages: []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: frenchPrompt,
		},
	},
}

// Deutsch
const germanPrompt = ``

var GermanRequest = openai.ChatCompletionRequest{
	Model: openai.GPT3Dot5Turbo,
	Messages: []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: germanPrompt,
		},
	},
}

// 한국어
const koreanPrompt = ``

var KoreanRequest = openai.ChatCompletionRequest{
	Model: openai.GPT3Dot5Turbo,
	Messages: []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: koreanPrompt,
		},
	},
}

// Русский
const russianPrompt = ``

var RussianRequest = openai.ChatCompletionRequest{
	Model: openai.GPT3Dot5Turbo,
	Messages: []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: russianPrompt,
		},
	},
}

// 日本語
const japanesePrompt = ``

var JapaneseRequest = openai.ChatCompletionRequest{
	Model: openai.GPT3Dot5Turbo,
	Messages: []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: japanesePrompt,
		},
	},
}
