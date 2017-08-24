package common

import (
	"io/ioutil"
	"log"
	"path"
	"testing"
)

const (
	testFile = "event_1.json"
)

func TestGetTestData(t *testing.T) {
	buf := GetTestData(testFile)

	data, err := ioutil.ReadFile(path.Join("./test_data", testFile))
	if err != nil {
		log.Fatal(err)
	}

	if string(buf) != string(data) {
		t.Errorf("Data doesn't match %v != %v\n", buf, data)
	}
}
