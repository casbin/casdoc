package main

import (
	"casdoc/config"
	. "casdoc/logger"
	"casdoc/utils"
	"fmt"
	"path"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestPolishDocs(t *testing.T) {
	Q := utils.WorkQueue{}
	err := Q.GetFileList(path.Join(config.RepoPath, "/docs/"))
	if err != nil {
		Logger.Errorf("Failed to get file list")
		panic(err)
	}

	counter := 0
	totalItems := len(Q.Item)

	for {
		p := Q.Pop()
		counter++
		Logger = log.WithField("rate", fmt.Sprintf("%d/%d", counter, totalItems))

		Logger.Info("Now polish: ", strings.TrimPrefix(p, config.RepoPath))
		err = utils.Polish(p)

		if err != nil {
			Q.AddToFailedList(p)
			Logger.Errorf("error: %v\n", err)
		}

		if Q.Empty() {
			break
		}
	}
}

func TestTranslateDocs(t *testing.T) {
	langs := []string{"zh", "fr", "de", "ko", "ru", "ja"}

	Q := utils.WorkQueue{}
	err := Q.GetFileList(path.Join(config.RepoPath, "/docs/"))
	if err != nil {
		Logger.Errorf("Failed to get file list")
		panic(err)
	}

	counter := 0
	totalItems := len(Q.Item)

	for _, lang := range langs {
		Logger = log.WithField("lang", lang)
		for {
			p := Q.Pop()
			counter++
			Logger = Logger.WithField("rate", fmt.Sprintf("%d/%d", counter, totalItems))

			Logger.Info("Now translate: ", strings.TrimPrefix(p, config.RepoPath))
			err = utils.Translate(p, lang)

			if err != nil {
				Q.AddToFailedList(p)
				Logger.Errorf("error: %v\n", err)
			}

			if Q.Empty() {
				break
			}
		}
	}
}
