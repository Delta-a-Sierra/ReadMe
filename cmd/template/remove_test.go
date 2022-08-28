package template

import (
	"fmt"
	"testing"

	"github.com/Delta-a-Sierra/ReadMe/data"
)

func TestRemoveCmd(t *testing.T) {
	var testData data.TemplateData
	if err := testData.LoadData("../test_data/data.json"); err != nil {
		t.Fatal(err)
	}
	for k := range testData.TemplatesInfo {
		fmt.Printf("test: removeTemplate %s\n", k)
		removeTemplates(testData.TemplatesInfo, k)
		if _, prs := testData.TemplatesInfo[k]; prs {
			t.Fatalf("expected template: %s to be removed however it is still present", k)
		}
	}
}
