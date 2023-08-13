package utils

import (
	"casdoc/config"
	. "casdoc/logger"
	"casdoc/prompt"
	"errors"
	"fmt"
	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// i18n path
const i18nPathPrefix = "/i18n/"
const i18nPathSuffix = "/docusaurus-plugin-content-docs/current/"

// returns the context of a doc
func getDocContext(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Errorf("Failed to get context of doc: %s", path)
		return "", nil
	}
	res := string(b)
	return res, nil
}

// return the number of tokens of a text
func getTokenNum(text string) (int, error) {
	tkm, err := tiktoken.EncodingForModel("gpt-3.5-turbo")
	if err != nil {
		Logger.Errorf("Failed to get token num")
		return 0, err
	}
	tokens := tkm.Encode(text, nil, nil)
	return len(tokens), nil
}

func SplitDoc(docContext string) []string {
	tokenNumber, err := getTokenNum(docContext)
	if err != nil {
		return nil
	}
	/*
	*	https://platform.openai.com/docs/models/gpt-3-5
	*	gpt-3.5-turbo model's maximum context length is 4096 tokens.
	**/
	var result []string
	if tokenNumber > 2048 {
		// split doc context by h2
		strArr := strings.Split(docContext, "\n## ")
		if len(strArr) < 2 {
			Logger.Error("Failed to split doc by h2")
			return nil
		}
		strArr[0] = strArr[0] + "\n## " + strArr[1]
		result = append(result, strArr[0])
		for i := 2; i < len(strArr); i++ {
			strArr[i] = "## " + strArr[i]
			result = append(result, strArr[i])
		}
		Logger.Info("Split doc to ", len(result), " parts")
	} else {
		result = append(result, docContext)
	}
	return result
}

// process a doc
func processDoc(p string, req openai.ChatCompletionRequest) (processedDoc string, err error) {

	docContext, err := getDocContext(p)
	if err != nil {
		return
	}

	contexts := SplitDoc(docContext)

	for i, c := range contexts {
		if len(contexts) != 1 {
			Logger.Info("Now process part ", i+1)
		}

		tokenNumber, err := getTokenNum(c)
		if err != nil {
			return "", err
		}

		err = Wait(tokenNumber * 2)
		if err != nil {
			Logger.Error("Failed to Wait for enough tokens")
			return "", err
		}

		ans, err := gpt(req, &c)
		if err != nil {
			return "", err
		}

		processedDoc += *ans + "\n"
	}

	return
}

// Polish a doc
func Polish(path string) error {
	polishedDoc, err := processDoc(path, prompt.PolishRequest)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, []byte(polishedDoc), 0644)
	if err != nil {
		Logger.Errorf("Unable to write file: %s \n", path)
		return err
	}
	return nil
}

// Translate a doc to another language
func Translate(docPath string, lang string) error {
	var req openai.ChatCompletionRequest
	switch lang {
	case "zh": // 中文
		req = prompt.ChineseRequest
	case "fr": // Français
		req = prompt.FrenchRequest
	case "de": // Deutsch
		req = prompt.GermanRequest
	case "ko": // 한국어
		req = prompt.KoreanRequest
	case "ru": // Русский
		req = prompt.RussianRequest
	case "ja": // 日本語
		req = prompt.JapaneseRequest
	default:
		return errors.New(fmt.Sprint("unknown language: ", lang))
	}

	translatedDoc, err := processDoc(docPath, req)
	if err != nil {
		return err
	}

	relativePath := strings.Split(docPath, path.Join(config.RepoPath, "/docs/"))[1]
	translatedDocPath := path.Join(config.RepoPath, i18nPathPrefix, lang, i18nPathSuffix,
		relativePath)

	_ = os.MkdirAll(filepath.Dir(translatedDocPath), 0755)
	err = os.WriteFile(translatedDocPath, []byte(translatedDoc), 0644)
	if err != nil {
		Logger.Errorf("Unable to write file: %s", docPath)
		return err
	}
	return nil
}
