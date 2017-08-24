package ctftime

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

/* Test */

func GetTestData(fname string) []byte {
	fpath := path.Join(test_dir, fname)

	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

/* HTTP */

func HttpResponseToStruct(r *http.Response, v interface{}) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("[ReadAll] ")
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		log.Fatal("[Unmarshal] ")
	}

	if valid, ok := v.(interface {
		OK() error
	}); ok {
		err = valid.OK()
		if err != nil {
			log.Fatal("[Validation] ")
		}
	}
}
