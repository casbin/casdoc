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
	"fmt"
	"io/fs"
	"path/filepath"
)

type WorkQueue struct {
	Item   []string
	failed []string
}

// GetFileList get all doc in the docs directory
func (q *WorkQueue) GetFileList(docDir string) error {
	isMarkdownFile := func(path string) bool {
		ext := filepath.Ext(path)
		return ext == ".md" || ext == ".mdx"
	}

	err := filepath.Walk(docDir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && isMarkdownFile(path) {
			q.Push(path)
		}
		return err
	})
	return err
}

// Push a doc to the work queue
func (q *WorkQueue) Push(path string) {
	q.Item = append(q.Item, path)
}

// Pop a doc from the work queue
func (q *WorkQueue) Pop() string {
	removed := q.Item[0]
	q.Item = q.Item[1:]
	return removed
}

// Empty judge whether the queue is Empty
func (q *WorkQueue) Empty() bool {
	return len(q.Item) == 0
}

// return the front of the queue
func (q *WorkQueue) front() string {
	return q.Item[0]
}

// return the back of the queue
func (q *WorkQueue) back() string {
	return q.Item[len(q.Item)-1]
}

// AddToFailedList add a doc to the failed list
func (q *WorkQueue) AddToFailedList(path string) {
	q.failed = append(q.failed, path)
}

func (q *WorkQueue) PrintFileList() {
	for _, s := range q.Item {
		fmt.Println(s)
	}
}
