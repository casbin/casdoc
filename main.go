package main

import (
	"casdoc/config"
	. "casdoc/logger"
	"casdoc/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"path"
	"strings"
)

func PolishDocs() {
	Q := utils.WorkQueue{}
	err := Q.GetFileList(path.Join(config.RepoPath, "/docs/"))
	Q.PrintFileList()
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
		err := utils.Polish(p)

		if err != nil {
			Q.AddToFailedList(p)
			Logger.Errorf("error: %v\n", err)
		}

		if Q.Empty() {
			break
		}
	}
}

func main() {
	PolishDocs()
}
