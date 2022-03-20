package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var fruits []string

	root := "public/img/products"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".png" {
			fruits = append(fruits, strings.TrimLeft(strings.TrimRight(path, ".png"), "public\\img\\products\\"))
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fruits {
		fmt.Println(file)
	}
}
