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
type UserTestDao struct{} 

func (d UserTestDao) GetUserById(id int) domains.User {
   return domains.User{Id: id, Name: "John Swift", Address: "1000 N 4th St", Email: "john@xyz.com"}
}

// Test Cases

func TestGetUserHandlerWithValidId(t *testing.T) {
	//Create test router
	router := getUserTestRouter()
    
    req, err := http.NewRequest("GET", "/users/2", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
   	router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := `{"Object":{"Id":2,"Name":"JohnSwift","Address":"1000N4thSt","Email":"john@xyz.com"},"Links":{"posts":"http:///users/2/posts"}}`
   	actual :=  rr.Body.String()
   	actual = strings.Replace(actual, "\n","",-1)
   	actual = strings.Replace(actual, " ","",-1)
    if  actual != expected {
        t.Errorf("handler returned unexpected body: got %s want %s",
          actual, expected)
    }
}

func TestGetUserHandlerWithInValidId(t *testing.T) {
	//Create test router
	router := getUserTestRouter()
    
    req, err := http.NewRequest("GET", "/users/1qw", nil)
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

func getUserTestRouter() *mux.Router {
	mockUserDao := UserTestDao{}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users/{userId}", getUser(mockUserDao)).Methods("GET")
	return router
}