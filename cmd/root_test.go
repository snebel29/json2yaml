package cmd

import (
	"testing"
	"strings"
)

func TestJson2yaml(t *testing.T) {
	json := []byte(`{"a": 1}`)
	if r := strings.Trim(json2yaml(json), "\n"); r != `a: 1` {
		t.Errorf("Returned [%s]", r)
	}
}

func TestIsJSON(t *testing.T) {
	json := []byte(`{"a": 1}`)
	if isJSON(json) == false {
		t.Error("Should be considered a valid json")
	}

	json = []byte(`{"a": 1`)
	if isJSON(json) == true {
		t.Error("Should be considered a not valid json")
	}

}
