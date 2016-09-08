package postDao;

import (
	"domains" 
	"daos/userDao"
)



var posts [10]domains.Post

func init() {
	post1 := domains.Post{Id: 1,  Text: "Hello from my new IPad", User: userDao.GetUserById(1)}
	post2 := domains.Post{Id: 2,  Text: "I'm Happy!!", User: userDao.GetUserById(1)}
	post3 := domains.Post{Id: 3,  Text: "Go USA #Rio2016", User: userDao.GetUserById(2)}
	post4 := domains.Post{Id: 4,  Text: "Phillips <3 #Rio2016", User: userDao.GetUserById(2)}
	post5 := domains.Post{Id: 5,  Text: "Can't Sleep :(", User: userDao.GetUserById(3)}
	post6 := domains.Post{Id: 6,  Text: "2:00am and still awake", User: userDao.GetUserById(3)}
	post7 := domains.Post{Id: 7,  Text: "In my way to Spain", User: userDao.GetUserById(4)}
	post8 := domains.Post{Id: 8,  Text: "Spain!!, here we go", User: userDao.GetUserById(4)}
	post9 := domains.Post{Id: 9,  Text: "if u can dream it you can do it", User: userDao.GetUserById(5)}
	post10 := domains.Post{Id: 10,  Text: "Never give up", User: userDao.GetUserById(5)}
	posts[0] = post1
	posts[1] = post2
	posts[2] = post3
	posts[3] = post4
	posts[4] = post5
	posts[5] = post6
	posts[6] = post7
	posts[7] = post8
	posts[8] = post9
	posts[9] = post10
}

func GetPostById(id int) domains.Post {
	return posts[id - 1]
}

func GetPosts() [len(posts)]domains.Post {
	return posts
}

func GetPostsByUserId(userId int) []domains.Post {
	var userPosts []domains.Post
	for i:= range posts {
		if (posts[i].User.Id == userId) {
			userPosts = append(userPosts, posts[i])
		}
	}
	return userPosts
}
