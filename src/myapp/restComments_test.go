package main 

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"domains"
	"github.com/gorilla/mux"
	"strings"
)

// Mock the DAO
type CommentTestDao struct{} 

func (d CommentTestDao) GetCommentsByPostId (postId int) [] domains.Comment {
	var postComments [] domains.Comment
	user := domains.User{Id: 1, Name: "John Swift", Address: "1000 N 4th St", Email: "john@xyz.com"}
	postComments = append(postComments, domains.Comment{Id: 1, Text:"(Y)", PostId: postId, User: user})	
	postComments = append(postComments, domains.Comment{Id: 2, Text:"(Y)", PostId: postId, User: user})
	return postComments;
}


// Test Cases

func TestGetCommentByPostHandlerWithValidId(t *testing.T) {
	//Create test router
	router := getCommentTestRouter()
    
    req, err := http.NewRequest("GET", "/posts/1/comments", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
   	router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `[{"Object":{"Id":1,"Text":"(Y)","PostId":1,"User":{"Id":1,"Name":"JohnSwift","Address":"1000N4thSt","Email":"john@xyz.com"}},"Links":{}},{"Object":{"Id":2,"Text":"(Y)","PostId":1,"User":{"Id":1,"Name":"JohnSwift","Address":"1000N4thSt","Email":"john@xyz.com"}},"Links":{}}]`
   	actual :=  rr.Body.String()
   	actual = strings.Replace(actual, "\n","",-1)
   	actual = strings.Replace(actual, " ","",-1)
    if  actual != expected {
        t.Errorf("handler returned unexpected body: got %s want %s",
          actual, expected)
    }
}

func TestGetCommentByPostHandlerWithInValidId(t *testing.T) {
	//Create test router
	router := getCommentTestRouter()
    
    req, err := http.NewRequest("GET", "/posts/111ew/comments", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
   	router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := ``
   	actual :=  rr.Body.String()
    if  actual != expected {
        t.Errorf("handler returned unexpected body: got %s want %s", actual, expected)
    }
}

func getCommentTestRouter() *mux.Router {
	mockCommentDao := CommentTestDao{}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/posts/{postId}/comments", getCommetsByPost(mockCommentDao)).Methods("GET")
	return router
}