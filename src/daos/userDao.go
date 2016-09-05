package daos;

import ("../domains")

var users  = make(map[int]domains.User)

func GetUserById(id int) domains.User {
	return users[id]
}
