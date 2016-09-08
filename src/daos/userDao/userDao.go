package userDao;

import ("domains")

var users  = make(map[int]domains.User)

func init() {
	// create dummy data of users
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
}

type IUserDao interface {
	GetUserById(id int) domains.User
}

type MyDao struct {}

func (d MyDao) GetUserById(id int) domains.User {
	return users[id]
}
