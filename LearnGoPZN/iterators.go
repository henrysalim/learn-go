package main

import (
	"fmt"
	"iter"
)

type User struct {
	ID   int
	Name string
}

// Traditional approach - returns full slice
func GetAllUsers() []User {
	var users []User
	// Imagine fetching 1M users from database
	for i := 0; i < 1_000_000; i++ {
		users = append(users, User{ID: i, Name: fmt.Sprintf("User%d", i)})
	}
	return users
}

/*
func iterateAllUsers(yield func(User) bool) {
	for i := 0; i < 1_000_000; i++ {
		if !yield(User{ID: i, Name: fmt.Sprintf("User%d", i)}) {
			return
		}
	}
}
*/

/*
func iterateAllUsers() iter.Seq[User] {
	return func(yield func(User) bool) {
		for i := 0; i < 1_000_000; i++ {
			if !yield(User{ID: i, Name: fmt.Sprintf("User%d", i)}) {
				return
			}
		}
	}
}
*/

func iterateAllUsers() iter.Seq2[int, User] {
	return func(yield func(int, User) bool) {
		for i := 0; i < 1_000_000; i++ {
			if !yield(i, User{ID: i, Name: fmt.Sprintf("User%d", i)}) {
				return
			}
		}
	}
}

func main() {
	// for user := range iterateAllUsers {
	// 	fmt.Println("👤 user", user)
	// }

	// for user := range iterateAllUsers() {
	// 	fmt.Println("👤 user", user)
	// }

	for idx, user := range iterateAllUsers() {
		fmt.Println("👤 user", user, "at index", idx)
	}
}
