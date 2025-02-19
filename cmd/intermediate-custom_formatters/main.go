package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("User{Name: %s, Age: %d}", u.Name, u.Age)
}

func main() {
	u := User{Name: "Bob", Age: 30}
	fmt.Println(u)
}
