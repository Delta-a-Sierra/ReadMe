package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	AppFolderPath string = os.Getenv("HOME") + "/.readme"
	DataFilePath  string = AppFolderPath + "/data.json"
	Data          TemplateData
)

type TemplateData struct {
	Tags          []string
	TemplatesInfo []templateInfo
}

type templateInfo struct {
	Name       string
	Filepath   string
	Tags       []string
	LastUsed   string
	UsageCount int
}

func LoadData() error {

	file, err := ioutil.ReadFile(DataFilePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(file), &Data)
	if err != nil {
		return err
	}

	return nil
}

func WriteData() {
	dataAsBytes, _ := json.MarshalIndent(Data, "", " ")
	ioutil.WriteFile(DataFilePath, dataAsBytes, 0644)
}
