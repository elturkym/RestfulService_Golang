package domains;

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

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