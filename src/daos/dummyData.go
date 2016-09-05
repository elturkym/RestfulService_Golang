package daos

import ("domains")


func init() {
	user1  := domains.User{Id: 1, Name: "John Swift", Address: "1000 N 4th St", Email: "john@xyz.com"}
	user2  := domains.User{Id: 2, Name: "Mike Dale", Address: "1000 S 4th St", Email: "mike@xyz.com"}
	user3  := domains.User{Id: 3, Name: "Ali Tim", Address: "1000 W 4th St", Email: "ali@xyz.com"}
	user4  := domains.User{Id: 4, Name: "David Beckham", Address: "1000 E 4th St", Email: "david@xyz.com"}
	user5  := domains.User{Id: 5, Name: "Lerman Joseph", Address: "1010 N 4th St", Email: "lerman@xyz.com"}
	users[1]  = user1;
	users[2]  = user2;
	users[3]  = user3;
	users[4]  = user4;
	users[5]  = user5;

	post1 := domains.Post{Id: 1,  Text: "Hello from my new IPad", User: GetUserById(1)}
	post2 := domains.Post{Id: 2,  Text: "I'm Happy!!", User: GetUserById(1)}
	post3 := domains.Post{Id: 3,  Text: "Go USA #Rio2016", User: GetUserById(2)}
	post4 := domains.Post{Id: 4,  Text: "Phillips <3 #Rio2016", User: GetUserById(2)}
	post5 := domains.Post{Id: 5,  Text: "Can't Sleep :(", User: GetUserById(3)}
	post6 := domains.Post{Id: 6,  Text: "2:00am and still awake", User: GetUserById(3)}
	post7 := domains.Post{Id: 7,  Text: "In my way to Spain", User: GetUserById(4)}
	post8 := domains.Post{Id: 8,  Text: "Spain!!, here we go", User: GetUserById(4)}
	post9 := domains.Post{Id: 9,  Text: "if u can dream it you can do it", User: GetUserById(5)}
	post10 := domains.Post{Id: 10,  Text: "Never give up", User: GetUserById(5)}
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

	comment1 := domains.Comment{Id: 1, Text:"Congrats", PostId: 1, User: GetUserById(2)}
	comment2 := domains.Comment{Id: 2, Text:"Good choice", PostId: 1, User: GetUserById(3)}
	comment3 := domains.Comment{Id: 3, Text:"(Y)", PostId: 2, User: GetUserById(4)}
	comment4 := domains.Comment{Id: 4, Text:"me too, lol!", PostId: 2, User: GetUserById(3)}
	comment5 := domains.Comment{Id: 5, Text:"GO GO :)", PostId: 3, User: GetUserById(1)}
	comment6 := domains.Comment{Id: 6, Text:"What does Rio2016 mean? :P", PostId: 3, User: GetUserById(4)}
	comment7 := domains.Comment{Id: 7, Text:"Incredible!", PostId: 4, User: GetUserById(4)}
	comment8 := domains.Comment{Id: 8, Text:"Why?!!", PostId: 5, User: GetUserById(1)}
	comment9 := domains.Comment{Id: 9, Text:"me too :(", PostId: 6, User: GetUserById(5)}
	comment10 := domains.Comment{Id: 10, Text:"without me !!", PostId: 7, User: GetUserById(1)}
	comment11 := domains.Comment{Id: 11, Text:"have a safe trip", PostId: 7, User: GetUserById(2)}
	comment12 := domains.Comment{Id: 12, Text:"Enjoy!!", PostId: 8, User: GetUserById(1)}
	comment13 := domains.Comment{Id: 13, Text:"Totally Agree", PostId: 9, User: GetUserById(5)}
	comment14 := domains.Comment{Id: 14, Text:"(Y)", PostId: 10, User: GetUserById(2)}

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