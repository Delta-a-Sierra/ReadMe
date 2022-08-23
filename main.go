/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"os"

	"github.com/Delta-a-Sierra/ReadMe/cmd"
	"github.com/Delta-a-Sierra/ReadMe/data"
)

func createAppFolder(appFolderPath string) error {
	if _, err := os.Stat(appFolderPath); os.IsNotExist(err) {
		err := os.MkdirAll(appFolderPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func createAppFile(dataFilePath string) error {
	if _, err := os.Stat(dataFilePath); os.IsNotExist(err) {
		_, err := os.Create(dataFilePath)
		if err != nil {
			return err
		}
		data.Data = data.TemplateData{}
		if err := data.WriteData(); err != nil {
			return err
		}
	}
	return nil
}

func createTemplateFolder(appFolderPath string) error {
	if _, err := os.Stat(appFolderPath + "/templates"); os.IsNotExist(err) {
		err := os.MkdirAll(appFolderPath+"/templates", os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := createAppFolder(data.AppFolderPath)
	if err != nil {
		panic(err)
	}
	err = createTemplateFolder(data.AppFolderPath)
	if err != nil {
		panic(err)
	}
	err = createAppFile(data.DataFilePath)
	if err != nil {
		panic(err)
	}

	cmd.Execute()
}
