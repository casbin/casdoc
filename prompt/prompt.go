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
