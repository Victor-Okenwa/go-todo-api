// package main

// import "fmt"

// type Todo struct {
// 	ID          int
// 	Title       string
// 	Description string
// 	Completed   bool
// }

// // ❌ This function does NOT work as expected
// func updateTitleBad(todo Todo, newTitle string) {
// 	todo.Title = newTitle
// }

// // ✅ We want to make this one work
// func updateTitleGood(todo *Todo, newTitle string) {
// 	todo.Title = newTitle
// }

// func main() {
// 	myTodo := Todo{
// 		ID:          1,
// 		Title:       "Learn Go Pointers",
// 		Description: "Understand & and *",
// 		Completed:   false,
// 	}

// 	fmt.Println("Before update:")
// 	fmt.Printf("Title: %s\n", myTodo.Title)

// 	// Try bad version
// 	updateTitleBad(myTodo, "Learn Pointers - Fixed!")
// 	fmt.Println("\nAfter bad update:")
// 	fmt.Printf("Title: %s\n", myTodo.Title)

// 	// Try good version
// 	updateTitleGood(&myTodo, "Learn Pointers - SUCCESS!")
// 	fmt.Println("\nAfter good update:")
// 	fmt.Printf("Title: %s\n", myTodo.Title)
// }
