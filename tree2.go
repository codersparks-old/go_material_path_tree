package main

import (
	"fmt"
	"github.com/disiqueira/gotree"
	"strings"
)

func main() {

	const SEPERATOR = "/"

	testData := [...]string{"/root/folder_1/folder_1_1/file_1_1_1", "/root/folder_1/folder_1_1/file_1_1_2", "/root/folder_1/folder_1_2/file_1_2_1", "/root/folder_1/folder_1_2/folder_1_2_1/file_1_2_1_1", "/root/folder_2/folder_2_1/file_2_1_1", "/root/folder_2/folder_2_1/file_2_1_2", "/root/folder_2/folder_2_2/file_2_2_1", "/root/folder_2/folder_2_2/folder_2_2_1/file_2_2_1_1"}

	rootTree := gotree.New("root")

	for _, path := range testData {

		currentTree := rootTree

		tokens := strings.Split(path, SEPERATOR)

		for _, token := range tokens {

			if token == "root" || len(token) == 0 {
				continue
			}

			var foundTree gotree.Tree
			foundTree = nil

			for _, tree := range currentTree.Items() {

				if tree.Text() == token {
					foundTree = tree
				}
			}

			if foundTree != nil {
				currentTree = foundTree
			} else {
				currentTree = currentTree.Add(token)
			}

		}

	}

	fmt.Print(rootTree.Print())
}
