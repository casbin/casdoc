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
	"path"
	"strings"
	"testing"

	"github.com/casbin/casdoc/config"
	. "github.com/casbin/casdoc/logger"
	"github.com/casbin/casdoc/utils"
)

func TestPolishDocs(t *testing.T) {
	Q := utils.WorkQueue{}
	err := Q.GetFileList(path.Join(config.RepoPath, "/docs/"))
	if err != nil {
		t.Fatalf("Failed to get file list, err: %v", err)
	}

	totalItems := len(Q.Item)

	for counter := 1; !Q.Empty(); counter++ {
		p := Q.Pop()
		t.Logf("rate: %d/%d, now polish: %s", counter, totalItems, strings.TrimPrefix(p, config.RepoPath))
		if err := utils.Polish(p); err != nil {
			Q.AddToFailedList(p)
			Logger.Errorf("error: %v\n", err)
		}
	}
}

func TestTranslateDocs(t *testing.T) {
	languages := []string{"zh", "fr", "de", "ko", "ru", "ja"}

	Q := utils.WorkQueue{}
	err := Q.GetFileList(path.Join(config.RepoPath, "/docs/"))
	if err != nil {
		Logger.Fatalf("Failed to get file list")
	}

	totalItems := len(Q.Item)

	for _, lang := range languages {
		for counter := 1; !Q.Empty(); counter++ {
			p := Q.Pop()
			t.Logf("lang: %s, rate: %d/%d, now translate: %s", lang, counter, totalItems, strings.TrimPrefix(p, config.RepoPath))
			if err := utils.Translate(p, lang); err != nil {
				Q.AddToFailedList(p)
				Logger.Errorf("error: %v\n", err)
			}
		}
	}
}
