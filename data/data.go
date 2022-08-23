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

func (t TemplateData) SortUsage(index int) {
	if index >= len(t.TemplatesInfo)-1 {
		return
	}
	key := t.TemplatesInfo[index+1]
	if key.UsageCount > t.TemplatesInfo[index].UsageCount {
		t.TemplatesInfo[index+1], t.TemplatesInfo[index] = t.TemplatesInfo[index], key
		if index >= 1 {
			t.SortUsage(index - 1)
		}
	}
	t.SortUsage(index + 1)
}

func (t TemplateData) SortRecent(index int) {
	if index >= len(t.TemplatesInfo)-1 {
		return
	}
	key := t.TemplatesInfo[index+1]
	if key.LastUsed.Before(t.TemplatesInfo[index].LastUsed) {
		t.TemplatesInfo[index+1], t.TemplatesInfo[index] = t.TemplatesInfo[index], key
		if index >= 1 {
			t.SortRecent(index - 1)
		}
	}
	t.SortRecent(index + 1)
}

func (t TemplateData) SortAgeAcsending(index int) {
	if index >= len(t.TemplatesInfo)-1 {
		return
	}
	key := t.TemplatesInfo[index+1]
	if key.Created.Before(t.TemplatesInfo[index].Created) {
		t.TemplatesInfo[index+1], t.TemplatesInfo[index] = t.TemplatesInfo[index], key
		if index >= 1 {
			t.SortAgeAcsending(index - 1)
		}
	}
	t.SortAgeAcsending(index + 1)
}

func (t TemplateData) SortAgeDescending(index int) {
	if index >= len(t.TemplatesInfo)-1 {
		return
	}
	key := t.TemplatesInfo[index+1]
	if key.Created.After(t.TemplatesInfo[index].Created) {
		t.TemplatesInfo[index+1], t.TemplatesInfo[index] = t.TemplatesInfo[index], key
		if index >= 1 {
			t.SortAgeDescending(index - 1)
		}
	}
	t.SortAgeDescending(index + 1)
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
