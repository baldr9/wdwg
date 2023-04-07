package main

import (
	"html/template"
	"os"
)

type User struct {
	Name          string
	JerseyNumber  int
	GamesPlayed   int
	FavoriteFoods []string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name:          "Jackie Robinson",
		JerseyNumber:  42,
		GamesPlayed:   33,
		FavoriteFoods: []string{"burger", "steak", "salmon"},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
