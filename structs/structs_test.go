package dummyapis

/*
Visit: https://jsonplaceholder.typicode.com/ for more testing URLS
*/

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	baseURL     string = "https://jsonplaceholder.typicode.com/"
	userSample  string = "samples/users_sample.json"
	usersURL    string = URLJoin(baseURL, "users")
	toDoSample  string = "samples/todo_sample.json"
	toDoURL     string = URLJoin(baseURL, "todos")
	postsSample string = "samples/posts_sample.json"
	postsURL    string = URLJoin(baseURL, "posts")
)

func TestUserStructSample(t *testing.T) {
	data := ReadFile(userSample)
	users := make(Users, 0)
	if err := json.Unmarshal(data, &users); err != nil {
		msg := fmt.Sprintf("Error while marshalling Users!: %v", err)
		t.Fatal(msg)
	} else {
		t.Log(users[:1]) // log the first element
	}
}

func TestUserStructHTTP(t *testing.T) {
	users := make(Users, 0)
	if err := GetJSONDecoded(usersURL, &users); err != nil {
		msg := fmt.Sprintf("Error while requesting Users URL: %v", err)
		t.Fatal(msg)
	} else {
		t.Log(users[:1])
	}
}

func TestToDoStructSample(t *testing.T) {
	data := ReadFile(toDoSample)
	todos := make(ToDos, 0)
	if err := json.Unmarshal(data, &todos); err != nil {
		msg := fmt.Sprintf("Error while marshalling ToDos!: %v", err)
		t.Fatal(msg)
	} else {
		t.Log(todos[:1])
	}
}

func TestToDoStructHTTP(t *testing.T) {
	todos := make(ToDos, 0)
	if err := GetJSONDecoded(toDoURL, &todos); err != nil {
		msg := fmt.Sprintf("Error while requesting ToDo URL: %v", err)
		t.Fatal(msg)
	} else {
		t.Log(todos[:1])
	}
}

func TestPostStructSample(t *testing.T) {
	data := ReadFile(postsSample)
	posts := make(Posts, 0)
	if err := json.Unmarshal(data, &posts); err != nil {
		msg := fmt.Sprintf("Error while marshalling Posts!: %v", err)
		t.Fatal(msg)
	} else {
		t.Log(posts[:1])
	}
}

func TestPostStructHTTP(t *testing.T) {
	posts := make(Posts, 0)
	if err := GetJSONDecoded(postsURL, &posts); err != nil {
		msg := fmt.Sprintf("Error while requesting Posts URL: %v", err)
		t.Fatal(msg)
	} else {
		t.Log(posts[:1])
	}
}
