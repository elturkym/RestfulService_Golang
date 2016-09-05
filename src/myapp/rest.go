package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"daos"
	"github.com/gorilla/mux"
	"strconv"
)


func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users/{userId}", getUser).Methods("GET")

	router.HandleFunc("/posts/{postId}/comments", getCommetsByPost).Methods("GET")

	router.HandleFunc("/users/{userId}/posts", getPostsByUserId).Methods("GET")
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts/{postId}", getPostById).Methods("GET")

	http.ListenAndServe("0.0.0.0:8080", router)
}

//Users
func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	userIdInt , err := strconv.Atoi(userId)
    if err == nil {
		response, err := getUserJsonResponse(userIdInt)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(response))
    }
}

func getUserJsonResponse(userId int) ([]byte, error) {
	return json.MarshalIndent(daos.GetUserById(userId), "", "  ")
}

// Comments
func getCommetsByPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]

	postIdInt , err := strconv.Atoi(postId)
    if err == nil {
		response, err := getCommentsJsonResponse(postIdInt)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(response))
    }
}

func getCommentsJsonResponse(postId int) ([]byte, error) {
	return json.MarshalIndent(daos.GetCommentsByPostId(postId), "", "  ")
}


// Posts

func getPostsByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId := vars["userId"]
	userIdInt , err := strconv.Atoi(userId)
    if err == nil {
		response, err := json.MarshalIndent(daos.GetPostsByUserId(userIdInt), "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(response))
    }
}



func getPosts(w http.ResponseWriter, r *http.Request) {
	response, err := json.MarshalIndent(daos.GetPosts(), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}


func getPostById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postId := vars["postId"]
	postIdInt , err := strconv.Atoi(postId)
    if err == nil {
		response, err := json.MarshalIndent(daos.GetPostById(postIdInt), "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(response))
    }
}


//End

