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
	err := filepath.Walk(docDir, func(path string, info fs.FileInfo, err error) error {
		var name = info.Name()
		var size = len(name)
		if !info.IsDir() && (name[size-3:] == "mdx" || name[size-2:] == "md") {
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
