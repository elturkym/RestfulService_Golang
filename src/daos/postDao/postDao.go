package postDao;

import (
	"domains" 
	"daos/userDao"
)



var posts []domains.Post

func init() {
	// create dummy data of posts
	myUserDao := userDao.MyDao{}
	posts = append(posts, domains.Post{Id: 1,  Text: "Hello from my new IPad", User: myUserDao.GetUserById(1)})
	posts = append(posts, domains.Post{Id: 2,  Text: "I'm Happy!!", User: myUserDao.GetUserById(1)})
	posts = append(posts, domains.Post{Id: 3,  Text: "Go USA #Rio2016", User: myUserDao.GetUserById(2)})
	posts = append(posts, domains.Post{Id: 4,  Text: "Phillips <3 #Rio2016", User: myUserDao.GetUserById(2)})
	posts = append(posts, domains.Post{Id: 5,  Text: "Can't Sleep :(", User: myUserDao.GetUserById(3)})
	posts = append(posts, domains.Post{Id: 6,  Text: "2:00am and still awake", User: myUserDao.GetUserById(3)})
	posts = append(posts, domains.Post{Id: 7,  Text: "In my way to Spain", User: myUserDao.GetUserById(4)})
	posts = append(posts, domains.Post{Id: 8,  Text: "Spain!!, here we go", User: myUserDao.GetUserById(4)})
	posts = append(posts, domains.Post{Id: 9,  Text: "if u can dream it you can do it", User: myUserDao.GetUserById(5)})
	posts = append(posts,  domains.Post{Id: 10,  Text: "Never give up", User: myUserDao.GetUserById(5)})
}

type IPostDao interface {
	GetPostById(id int) domains.Post
	GetPosts() []domains.Post 
	GetPostsByUserId(userId int) []domains.Post 
}

type MyDao struct {}


func (d MyDao) GetPostById(id int) domains.Post {
	return posts[id - 1]
}

func (d MyDao) GetPosts() []domains.Post {
	return posts
}

func (d MyDao) GetPostsByUserId(userId int) []domains.Post {
	var userPosts []domains.Post
	for i:= range posts {
		if (posts[i].User.Id == userId) {
			userPosts = append(userPosts, posts[i])
		}
	}
	return userPosts
}
