package common

import (
	"io/ioutil"
	"log"
	"path"
)

const (
	test_dir = "./test_data/"
)

func GetTestData(fname string) []byte {
	fpath := path.Join(test_dir, fname)

	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
