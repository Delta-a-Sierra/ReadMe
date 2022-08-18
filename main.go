/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"os"

	"github.com/Delta-a-Sierra/ReadMe/cmd"
)

var AppFolderPath string = "~/.readme"

func createAppFolder() error {
	if _, err := os.Stat(AppFolderPath); os.IsNotExist(err) {
		err := os.MkdirAll("~/.readme", os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := createAppFolder()
	if err != nil {
		panic(err)
	}

	cmd.Execute()
}
