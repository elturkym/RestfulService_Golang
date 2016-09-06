package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"daos"
	"domains"
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

	http.ListenAndServe(":8080", router)
}

func createHateoasObject(subject interface{}, host string) domains.Hateoas {
	h := domains.Hateoas{subject, make(map[string]string)}
	switch s := subject.(type) {
	case domains.User:
		h.Links["posts"] = fmt.Sprintf("http://%s/users/%d/posts", host, s.Id)
	case domains.Post:
		h.Links["comments"] = fmt.Sprintf("http://%s/posts/%d/comments", host, s.Id)
	}
	return h
}

func writeJsonResponse(subject interface{}, w http.ResponseWriter) {
		response, err := json.MarshalIndent(subject, "", "  ")
		if err != nil {
			panic(err)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, string(response))
}

//Users
func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	userIdInt , err := strconv.Atoi(userId)
    if err == nil {
		 h := createHateoasObject(daos.GetUserById(userIdInt), r.Host)
		 writeJsonResponse(h, w)
    }
}


// Comments
func getCommetsByPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]

	postIdInt , err := strconv.Atoi(postId)
    if err == nil {
		comments := daos.GetCommentsByPostId(postIdInt)
		var hateoas [] domains.Hateoas
		for i := range comments {
			hateoas = append(hateoas, createHateoasObject(comments[i], r.Host));
		}
		writeJsonResponse(hateoas, w)
    }
}

// Posts

func getPostsByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId := vars["userId"]
	userIdInt , err := strconv.Atoi(userId)
    if err == nil {
    	posts := daos.GetPostsByUserId(userIdInt)
		writeJsonResponse(buildPostsHateoas(posts, r.Host), w)
    }
}



func getPosts(w http.ResponseWriter, r *http.Request) {
	posts := daos.GetPosts()
	writeJsonResponse(buildPostsHateoas(posts[:], r.Host), w)
}


func buildPostsHateoas (posts []domains.Post, host string) [] domains.Hateoas {
	var hateoas [] domains.Hateoas
	for i := range posts {
		hateoas = append(hateoas, createHateoasObject(posts[i], host));
	}
	return hateoas;
}

func getPostById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postId := vars["postId"]
	postIdInt , err := strconv.Atoi(postId)
    if err == nil {
    	post := daos.GetPostById(postIdInt);
		writeJsonResponse(createHateoasObject(post, r.Host), w)
    }
}


//End