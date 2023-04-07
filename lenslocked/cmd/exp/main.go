package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// user := struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	Name: "Susan Smith",
	// 	Age:  111,
	// }

	user := User{
		Name: "Jon Calhoun",
		Age:  111,
		Meta: UserMeta{
			Visits: 4,
		},
	}

	// fmt.Println(user.Meta.Visits)

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
