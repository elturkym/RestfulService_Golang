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
type PostTestDao struct{} 

func (d PostTestDao) GetPostById(id int) domains.Post {
  user := domains.User{Id: id, Name: "John Swift", Address: "1000 N 4th St", Email: "john@xyz.com"}
  return domains.Post{Id: 1,  Text: "Spain!!, here we go", User: user}
}

func (d PostTestDao) GetPostsByUserId(userId int) []domains.Post {
  var userPosts []domains.Post
  user := domains.User{Id: userId, Name: "John Swift", Address: "1000 N 4th St", Email: "john@xyz.com"}
  userPosts = append(userPosts, domains.Post{Id: 1,  Text: "if u can dream it you can do it", User: user})
  return userPosts
}

func (d PostTestDao) GetPosts() []domains.Post {
  return d.GetPostsByUserId(1)
}


// Test Cases

func TestGetPostsByUserIdHandlerWithValidId(t *testing.T) {
  //Create test router
  router := getPostTestRouter()
    
    req, err := http.NewRequest("GET", "/users/1/posts", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `[{"Object":{"Id":1,"Text":"ifucandreamityoucandoit","User":{"Id":1,"Name":"JohnSwift","Address":"1000N4thSt","Email":"john@xyz.com"}},"Links":{"comments":"http:///posts/1/comments"}}]`
    actual :=  rr.Body.String()
    actual = strings.Replace(actual, "\n","",-1)
    actual = strings.Replace(actual, " ","",-1)
    if  actual != expected {
        t.Errorf("handler returned unexpected body: got %s want %s",
          actual, expected)
    }
}

func TestGetPostsByUserIdHandlerWithInvalidId(t *testing.T) {
  //Create test router
  router := getPostTestRouter()
    
    req, err := http.NewRequest("GET", "/users/2asa/posts", nil)
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
        t.Errorf("handler returned unexpected body: got %s want %s",
          actual, expected)
    }
}

func TestGetPostsHandler(t *testing.T) {
  //Create test router
  router := getPostTestRouter()
    
    req, err := http.NewRequest("GET", "/posts", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `[{"Object":{"Id":1,"Text":"ifucandreamityoucandoit","User":{"Id":1,"Name":"JohnSwift","Address":"1000N4thSt","Email":"john@xyz.com"}},"Links":{"comments":"http:///posts/1/comments"}}]`
    actual :=  rr.Body.String()
    actual = strings.Replace(actual, "\n","",-1)
    actual = strings.Replace(actual, " ","",-1)
    if  actual != expected {
        t.Errorf("handler returned unexpected body: got %s want %s",
          actual, expected)
    }
}

func TestGetPostsByIdHandlerWithValidId(t *testing.T) {
  //Create test router
  router := getPostTestRouter()
    
    req, err := http.NewRequest("GET", "/posts/1", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `{"Object":{"Id":1,"Text":"Spain!!,herewego","User":{"Id":1,"Name":"JohnSwift","Address":"1000N4thSt","Email":"john@xyz.com"}},"Links":{"comments":"http:///posts/1/comments"}}`
    actual :=  rr.Body.String()
    actual = strings.Replace(actual, "\n","",-1)
    actual = strings.Replace(actual, " ","",-1)
    if  actual != expected {
        t.Errorf("handler returned unexpected body: got %s want %s",
          actual, expected)
    }
}

func TestGetPostsByIdHandlerWithInvalidId(t *testing.T) {
  //Create test router
  router := getPostTestRouter()
    
    req, err := http.NewRequest("GET", "/posts/12sss", nil)
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
        t.Errorf("handler returned unexpected body: got %s want %s",
          actual, expected)
    }
}

func getPostTestRouter() *mux.Router {
  mockPostDao := PostTestDao{}
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/users/{userId}/posts", getPostsByUserId(mockPostDao)).Methods("GET")
  router.HandleFunc("/posts", getPosts(mockPostDao)).Methods("GET")
  router.HandleFunc("/posts/{postId}", getPostById(mockPostDao)).Methods("GET")
  return router
}