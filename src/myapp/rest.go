package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"daos/commentDao"
	"daos/postDao"
	"daos/userDao"
	"domains"
	"github.com/gorilla/mux"
	"strconv"
)


func main() {
	myCommentDao := commentDao.MyDao{}
	myUserDao := userDao.MyDao{}
	myPostDao := postDao.MyDao{}
	
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users/{userId}", getUser(myUserDao)).Methods("GET")
	router.HandleFunc("/posts/{postId}/comments", getCommetsByPost(myCommentDao)).Methods("GET")
	router.HandleFunc("/users/{userId}/posts", getPostsByUserId(myPostDao)).Methods("GET")
	router.HandleFunc("/posts", getPosts(myPostDao)).Methods("GET")
	router.HandleFunc("/posts/{postId}", getPostById(myPostDao)).Methods("GET")

	http.ListenAndServe("0.0.0.0:8080", router)
}

//Users
func getUser(iuserDao userDao.IUserDao) func(http.ResponseWriter, *http.Request) {

	return func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId := vars["userId"]

		userIdInt , err := strconv.Atoi(userId)
	    if err == nil {
			 h := createHateoasObject(iuserDao.GetUserById(userIdInt), r.Host)
			 writeJsonResponse(h, w)
	    }
	}
}


// Comments
func getCommetsByPost(icommentDao commentDao.ICommentDao) func(http.ResponseWriter, *http.Request) {
	return  func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		postId := vars["postId"]

		postIdInt , err := strconv.Atoi(postId)
	    if err == nil {
			comments := icommentDao.GetCommentsByPostId(postIdInt)
			var hateoas [] domains.Hateoas
			for i := range comments {
				hateoas = append(hateoas, createHateoasObject(comments[i], r.Host));
			}
			writeJsonResponse(hateoas, w)
	    }
	}
}

// Posts

func getPostsByUserId(ipostDao postDao.IPostDao) func(http.ResponseWriter, *http.Request) { 
	return func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		userId := vars["userId"]
		userIdInt , err := strconv.Atoi(userId)
	    if err == nil {
	    	posts := ipostDao.GetPostsByUserId(userIdInt)
			writeJsonResponse(buildPostsHateoas(posts, r.Host), w)
	    }
	}
}



func getPosts(ipostDao postDao.IPostDao) func(http.ResponseWriter, *http.Request) { 
	return func (w http.ResponseWriter, r *http.Request) {
		posts := ipostDao.GetPosts()
		writeJsonResponse(buildPostsHateoas(posts[:], r.Host), w)
	}
}


func buildPostsHateoas (posts []domains.Post, host string) [] domains.Hateoas {
	var hateoas [] domains.Hateoas
	for i := range posts {
		hateoas = append(hateoas, createHateoasObject(posts[i], host));
	}
	return hateoas;
}

func getPostById(ipostDao postDao.IPostDao) func(http.ResponseWriter, *http.Request) { 
	return func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		postId := vars["postId"]
		postIdInt , err := strconv.Atoi(postId)
	    if err == nil {
	    	post := ipostDao.GetPostById(postIdInt);
			writeJsonResponse(createHateoasObject(post, r.Host), w)
	    }
	}
}

// Hateoas Handling

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
//End