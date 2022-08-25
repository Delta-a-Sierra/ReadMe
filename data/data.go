package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

var (
	AppFolderPath string = os.Getenv("HOME") + "/.readme"
	DataFilePath  string = AppFolderPath + "/data.json"
	Data          TemplateData
)

type TemplateInfo struct {
	Filepath   string
	Tags       map[string]string
	LastUsed   time.Time
	Created    time.Time
	UsageCount int
}
type TemplateData struct {
	Tags          map[string][]string
	TemplatesInfo map[string]TemplateInfo
	templateNames []string
}

func (t TemplateData) TemplateNames() []string {
	return t.templateNames
}

func (t TemplateData) FillTemplateNames() {
	if t.templateNames == nil {
		for k := range t.TemplatesInfo {
			t.templateNames = append(t.templateNames, k)
		}
		sort.Strings(t.templateNames)
		Data.templateNames = t.templateNames
	}
}

func (t TemplateData) sortTemplates(sortAlg func(t1, t2 TemplateInfo) bool, index ...int) {
	if len(index) == 0 {
		index = []int{0}
	}
	if t.templateNames == nil {
		for k := range t.TemplatesInfo {
			t.templateNames = append(t.templateNames, k)
		}
	}
	if index[0] >= len(t.templateNames)-1 {
		Data.templateNames = t.templateNames
		return
	}

	t1, t2 := t.TemplatesInfo[t.templateNames[index[0]]], t.TemplatesInfo[t.templateNames[index[0]+1]]
	if sortAlg(t1, t2) {
		t.templateNames[index[0]+1], t.templateNames[index[0]] =
			t.templateNames[index[0]], t.templateNames[index[0]+1]
		if index[0] >= 1 {
			t.sortTemplates(sortAlg, index[0]-1)
		}
	}
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
