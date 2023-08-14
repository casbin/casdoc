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

package main

import (
	"fmt"
	"path"
	"strings"
	"testing"

	"github.com/casbin/casdoc/config"
	. "github.com/casbin/casdoc/logger"
	"github.com/casbin/casdoc/utils"
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
