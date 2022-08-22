package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	AppFolderPath string = os.Getenv("HOME") + "/.readme"
	DataFilePath  string = AppFolderPath + "/data.json"
	Data          TemplateData
)

type TemplateData struct {
	Tags          []string
	TemplatesInfo []TemplateInfo
}

type TemplateInfo struct {
	Name       string
	Filepath   string
	Tags       []string
	LastUsed   time.Time
	Created    time.Time
	UsageCount int
}

func LoadData() error {

	file, err := ioutil.ReadFile(DataFilePath)
	if err != nil {
		return fmt.Errorf("unable read datafile: %s", err)
	}

	err = json.Unmarshal([]byte(file), &Data)
	if err != nil {
		return fmt.Errorf("failed unmarshalling datafiles json's: %s", err)
	}

	return nil
}

func WriteData() error {
	dataAsBytes, err := json.MarshalIndent(Data, "", " ")
	if err != nil {
		return errors.New("failed marshalling data back to json")
	}
	if err := ioutil.WriteFile(DataFilePath, dataAsBytes, 0644); err != nil {
		return errors.New("failed writing json data back to data file")
	}
	return nil
}
