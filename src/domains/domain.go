package domains;

type User struct {
	Id int
	Name string
	Address string
	Email string
}

type Post struct {
	Id int
	Text string
	User User
}

type Comment struct {
	Id int
	Text string
	PostId int
	User User
}

type Object interface{}

type Hateoas struct {
	Object
	Links map[string]string
}