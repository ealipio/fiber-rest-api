package main

type userType struct {
	Id      string
	Name    string
	IsAdmin bool
}

var users = []userType{
	{Id: "1", Name: "Jon", IsAdmin: true},
	{Id: "2", Name: "Genesis", IsAdmin: false},
	{Id: "3", Name: "Olga", IsAdmin: true},
}
