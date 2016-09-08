package commentDao;

import (
	"domains"
	"daos/userDao"
)

var comments [14]domains.Comment

func init() {
	comment1 := domains.Comment{Id: 1, Text:"Congrats", PostId: 1, User: userDao.GetUserById(2)}
	comment2 := domains.Comment{Id: 2, Text:"Good choice", PostId: 1, User: userDao.GetUserById(3)}
	comment3 := domains.Comment{Id: 3, Text:"(Y)", PostId: 2, User: userDao.GetUserById(4)}
	comment4 := domains.Comment{Id: 4, Text:"me too, lol!", PostId: 2, User: userDao.GetUserById(3)}
	comment5 := domains.Comment{Id: 5, Text:"GO GO :)", PostId: 3, User: userDao.GetUserById(1)}
	comment6 := domains.Comment{Id: 6, Text:"What does Rio2016 mean? :P", PostId: 3, User: userDao.GetUserById(4)}
	comment7 := domains.Comment{Id: 7, Text:"Incredible!", PostId: 4, User: userDao.GetUserById(4)}
	comment8 := domains.Comment{Id: 8, Text:"Why?!!", PostId: 5, User: userDao.GetUserById(1)}
	comment9 := domains.Comment{Id: 9, Text:"me too :(", PostId: 6, User: userDao.GetUserById(5)}
	comment10 := domains.Comment{Id: 10, Text:"without me !!", PostId: 7, User: userDao.GetUserById(1)}
	comment11 := domains.Comment{Id: 11, Text:"have a safe trip", PostId: 7, User: userDao.GetUserById(2)}
	comment12 := domains.Comment{Id: 12, Text:"Enjoy!!", PostId: 8, User: userDao.GetUserById(1)}
	comment13 := domains.Comment{Id: 13, Text:"Totally Agree", PostId: 9, User: userDao.GetUserById(5)}
	comment14 := domains.Comment{Id: 14, Text:"(Y)", PostId: 10, User: userDao.GetUserById(2)}

	comments[0] = comment1
	comments[1] = comment2
	comments[2] = comment3
	comments[3] = comment4
	comments[4] = comment5
	comments[5] = comment6
	comments[6] = comment7
	comments[7] = comment8
	comments[8] = comment9
	comments[9] = comment10
	comments[10] = comment11
	comments[11] = comment12
	comments[12] = comment13
	comments[13] = comment14
}

func GetCommentsById(id int) domains.Comment {
	return comments[id - 1]
}

func GetCommentsByPostId (postId int) [] domains.Comment {
	var postComments [] domains.Comment
	for i := range comments {
		if (comments[i].PostId == postId) {
			postComments = append(postComments, comments[i])
		}
	}
	return postComments;
}
