package daos;

import ("domains")



var posts [10]domains.Post


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
