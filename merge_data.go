package main

import (
	"fmt"

	"github.com/chiheb-nexus/dummyapis/structs"
)

func main() {
	var (
		usersChan = make(chan structs.Users)
		postsChan = make(chan structs.Posts)
		todosChan = make(chan structs.ToDos)
	)
	go GetUsers(usersChan)
	go GetPosts(postsChan)
	go GetToDos(todosChan)
	users := <-usersChan
	posts := <-postsChan
	todos := <-todosChan
	fmt.Println(users[:1], posts[:1], todos[:1])

}

// GetUsers return all data from Users URL
func GetUsers(usersChan chan structs.Users) {
	users := make(structs.Users, 0)
	if err := structs.GetJSONDecoded(structs.UsersURL, &users); err != nil {
		usersChan <- nil
	} else {
		usersChan <- users
	}
}

// GetPosts return all data from Posts URL
func GetPosts(postsChan chan structs.Posts) {
	posts := make(structs.Posts, 0)
	if err := structs.GetJSONDecoded(structs.PostsURL, &posts); err != nil {
		postsChan <- nil
	} else {
		postsChan <- posts
	}
}

// GetToDos return all data from ToDo URL
func GetToDos(todosChan chan structs.ToDos) {
	todos := make(structs.ToDos, 0)
	if err := structs.GetJSONDecoded(structs.ToDoURL, &todos); err != nil {
		todosChan <- nil
	} else {
		todosChan <- todos
	}
}
