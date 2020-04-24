package structs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"time"
)

var (
	// BaseURL is the Base URL
	BaseURL    string = "https://jsonplaceholder.typicode.com/"
	userSample string = "samples/users_sample.json"
	// UsersURL is the users data URL
	UsersURL   string = URLJoin(BaseURL, "users")
	toDoSample string = "samples/todo_sample.json"
	// ToDoURL is todo users data URL
	ToDoURL     string = URLJoin(BaseURL, "todos")
	postsSample string = "samples/posts_sample.json"
	// PostsURL is users posts data URL
	PostsURL string = URLJoin(BaseURL, "posts")
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
	client := &http.Client{Timeout: 100 * time.Second} // Timeout
	resp, err := client.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
