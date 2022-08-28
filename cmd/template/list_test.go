package template

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Delta-a-Sierra/ReadMe/data"
)

type filterResults struct {
	filterArg string
	results   any
}

func TestFilterTemplate(t *testing.T) {
	var testData data.TemplateData
	filterTests := []filterResults{
		{"tag=favourite", []string{"alpha1"}},
		{"tag=fish", "no tag fish in tag collection"},
	}

	for _, test := range filterTests {
		fmt.Printf("testing: filterTemplates %s\n", test.filterArg)
		filterArgArr := strings.Split(test.filterArg, "=")

		if err := testData.LoadData("../test_data/data.json"); err != nil {
			t.Fatal(err)
		}

		if err := filterTemplates(&testData, test.filterArg); err != nil {
			if string(err.Error()) != test.results {
				t.Fatal("unexpected error:", err)
			}
			return
		}

		switch filterArgArr[0] {
		case "tag":
			for name, temps := range testData.TemplatesInfo {
				if _, prs := temps.Tags[filterArgArr[1]]; !prs {
					t.Fatalf("tag %s not found in template: %s", filterArgArr[1], name)
				}
			}
		}
	}
}
