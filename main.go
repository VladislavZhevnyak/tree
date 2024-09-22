package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func tree(path string, indent string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, file := range files {
		fmt.Print(indent)

		if i == len(files)-1 {
			fmt.Println("└───" + file.Name())
			if file.IsDir() {
				tree(filepath.Join(path, file.Name()), indent+"    ")
			}
		} else {
			fmt.Println("├───" + file.Name())
			if file.IsDir() {
				tree(filepath.Join(path, file.Name()), indent+"│   ")
			}
		}
	}
}

func main() {

	tree(".", "")
}


