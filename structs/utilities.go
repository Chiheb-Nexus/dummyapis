package dummyapis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"time"
)

// URLJoin a handler funcs to join URLs
func URLJoin(base string, paths ...string) string {
	path := path.Join(paths...)
	return fmt.Sprintf("%s/%s", strings.TrimRight(base, "/"), strings.TrimLeft(path, "/"))

}

// ReadFile reads a file
func ReadFile(file string) []byte {
	if data, err := ioutil.ReadFile(file); err != nil {
		msg := fmt.Sprintf("Error while reading file <%s>: %v", file, err)
		panic(msg)
	} else {
		return data
	}
}

// GetJSONDecoded reads & decodes JSON form URL
func GetJSONDecoded(url string, target interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second} // Timeout
	resp, err := client.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
