package main

import (
	"fmt"
	"strings"
)

type node struct {
	name     string
	path     string
	children []node
}

func main() {

	testData := [...]string{"/root/folder_1/folder_1_1/file_1_1_1", "/root/folder_1/folder_1_1/file_1_1_2", "/root/folder_1/folder_1_2/file_1_2_1", "/root/folder_1/folder_1_2/folder_1_2_1/file_1_2_1_1", "/root/folder_2/folder_2_1/file_2_1_1", "/root/folder_2/folder_2_1/file_2_1_2", "/root/folder_2/folder_2_2/file_2_2_1", "/root/folder_2/folder_2_2/folder_2_2_1/file_2_2_1_1"}

	m := make(map[string]node)

	m["/root"] = node{"root", "/root", []node{}}

	for _, path := range testData {

		fmt.Println(path)

		tokens := strings.Split(path, "/")

		for index, token := range tokens {

			if token == "root" {
				continue
			}

			parentPath := strings.Join(tokens[0:index], "/")

			parentNode, parentNodePresent := m[parentPath]

			if !parentNodePresent {
				fmt.Println(index)
				parentNode := node{tokens[index], parentPath, []node{}}
				m[parentPath] = parentNode
			}

			fmt.Println("ParentNode: " + parentNode.path)

			fmt.Println("Token: " + token)
			fmt.Println("Path:  " + parentPath)
		}
	}

	fmt.Println(m)
}
