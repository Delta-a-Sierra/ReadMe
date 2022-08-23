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

type TemplateInfo struct {
	Name       string
	Filepath   string
	Tags       []string
	LastUsed   time.Time
	Created    time.Time
	UsageCount int
}
type TemplateData struct {
	Tags          []string
	TemplatesInfo []TemplateInfo
}

func (t TemplateData) sortTemplates(sortAlg func(t1, t2 TemplateInfo) bool, index ...int) {
	if len(index) == 0 {
		index = []int{0}
	}
	if index[0] >= len(t.TemplatesInfo)-1 {
		fmt.Println("final recusion")
		return
	}
	key := t.TemplatesInfo[index[0]+1]
	if sortAlg(t.TemplatesInfo[index[0]], key) {
		fmt.Println("sort alg true")
		t.TemplatesInfo[index[0]+1], t.TemplatesInfo[index[0]] = t.TemplatesInfo[index[0]], key
		if index[0] >= 1 {
			fmt.Println("recusion down", index[0])
			t.sortTemplates(sortAlg, index[0]-1)
		}
	}
	fmt.Println("recusion up", index[0])
	t.sortTemplates(sortAlg, index[0]+1)
}

func (t TemplateData) SortRecent() {
	t.sortTemplates(
		func(t1, t2 TemplateInfo) bool {
			if t1.LastUsed.Before(t2.LastUsed) {
				return true
			}
			return false
		},
	)
}

func (t TemplateData) SortUsage() {
	t.sortTemplates(
		func(t1, t2 TemplateInfo) bool {
			if t1.UsageCount < t2.UsageCount {
				return true
			}
			return false
		},
	)
}

func (t TemplateData) SortAgeAcsending() {
	t.sortTemplates(
		func(t1, t2 TemplateInfo) bool {
			if t1.Created.After(t2.Created) {
				return true
			}
			return false
		},
	)
}

func (t TemplateData) SortAgeDescending() {
	t.sortTemplates(
		func(t1, t2 TemplateInfo) bool {
			if t1.Created.Before(t2.Created) {
				return true
			}
			return false
		},
	)
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
