package ctftime

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
)

/* Test */

// テストに使うjsonファイルのデータを読み込んでくる
func getTestData(fname string) []byte {
	fpath := path.Join(test_dir, fname)

	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

/* HTTP */

// HTTPレスポンスのボディから構造体へ変換するための関数
func httpResponseToStruct(r *http.Response, v interface{}) {
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

// 年度をキーにしているなど、今後キーが変更される場合、mapで取れるようにするための関数
func httpResponseToMap(r *http.Response, v interface{}) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("[ReadAll ")
	}

	decoder := json.NewDecoder(strings.NewReader(string(body)))
	err = decoder.Decode(&v)
	if err != nil && err != io.EOF {
		log.Fatal(err)
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
