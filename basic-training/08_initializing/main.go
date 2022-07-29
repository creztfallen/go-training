package main

import "fmt"

type User struct {
	Id       int
	Name     string
	Location string
}

type Player struct {
	// Id       int
	// Name     string
	// Location string
	User
	GameId int
}

func main() {
	// p := Player{}
	// p.Id = 42
	// p.Name = "Matt"
	// p.Location = "LA"
	// p.GameId = 90404
	// fmt.Printf("%+v", p) // returns {User:{Id:42 Name:Matt Location:LA} GameId:90404}%

	p := Player{User{Id: 42, Name: "Matt", Location: "LA"}, 90404}
	fmt.Printf("Id: %d, Name: %s, Location:%s, Game id:%d\n", p.Id, p.Name, p.Location, p.GameId) //returns Id: 42, Name: Matt, Location:LA, Game id:90404

	p.Id = 11
	fmt.Printf("%+v", p)
}
