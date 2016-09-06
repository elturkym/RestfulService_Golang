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

	http.ListenAndServe("0.0.0.0:8080", router)
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

//Users
func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	userIdInt , err := strconv.Atoi(userId)
    if err == nil {
		 h := createHateoasObject(daos.GetUserById(userIdInt), r.Host)
		 response, err := json.MarshalIndent(h, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(response))
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
		response, err := json.MarshalIndent(hateoas, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(response))
    }
}

// Posts

func getPostsByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId := vars["userId"]
	userIdInt , err := strconv.Atoi(userId)
    if err == nil {
    	posts := daos.GetPostsByUserId(userIdInt)
		response, err := json.MarshalIndent(buildPostsHateoas(posts, r.Host), "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(response))
    }
}



func getPosts(w http.ResponseWriter, r *http.Request) {
	posts := daos.GetPosts()
	response, err := json.MarshalIndent(buildPostsHateoas(posts[:], r.Host), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(response))
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
		response, err := json.MarshalIndent(createHateoasObject(post, r.Host), "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(response))
    }
}


//End