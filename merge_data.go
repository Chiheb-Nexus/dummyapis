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
	formated := prettyFormat(users, posts, todos)
	for userID, data := range formated {
		fmt.Printf("########\n")
		for key, value := range data {
			fmt.Printf("<UserId: %d> %s: %v\n", userID, key, value)
		}
	}

}

func prettyFormat(users structs.Users, posts structs.Posts, todos structs.ToDos) map[int64]map[string]interface{} {
	prettifier := make(map[int64]map[string]interface{})
	for _, user := range users {
		prettifier[user.ID] = make(map[string]interface{})
		prettifier[user.ID]["username"] = user.Username
		prettifier[user.ID]["name"] = user.Name
	}
	for _, post := range posts {
		if value, ok := prettifier[post.UserID]; ok {
			value["title"] = post.Title
			value["body"] = post.Body
		}
	}
	for _, todo := range todos {
		if value, ok := prettifier[todo.UserID]; ok {
			value["todoTitle"] = todo.Title
			value["completed"] = todo.Completed
		}
	}
	return prettifier
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
