/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"os"

	"github.com/Delta-a-Sierra/ReadMe/cmd"
)

var appFolderPath string = "~/.readme"

func createAppFolder(appFolderPath string) error {
	if _, err := os.Stat(appFolderPath); os.IsNotExist(err) {
		err := os.MkdirAll(appFolderPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := createAppFolder(appFolderPath)
	if err != nil {
		panic(err)
	}

	cmd.Execute()
}
