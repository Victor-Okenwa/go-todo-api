// package main

// import (
// 	"fmt"
// )

// // ==================== STRUCT EXAMPLE ====================

// type Todo struct {
// 	ID    int
// 	Title string
// 	Done  bool
// }

// // This function takes a pointer to struct (*Todo)
// func markDoneStruct(t *Todo) {
// 	t.Done = true
// 	fmt.Println("Inside markDoneStruct - Todo is now done:", t.Done)
// }

// // ==================== INTERFACE EXAMPLE ====================

// // This is an interface (similar to http.ResponseWriter)
// type Notifier interface {
// 	Notify(message string)
// }

// // Concrete type that implements the interface
// type EmailNotifier struct {
// 	Email string
// }

// // This method has a pointer receiver (*EmailNotifier)
// func (n *EmailNotifier) Notify(message string) {
// 	fmt.Printf("Email sent to %s: %s\n", n.Email, message)
// }

// // This function takes the INTERFACE (no * needed)
// func sendNotification(notifier Notifier, message string) {
// 	notifier.Notify(message) // We call method on the interface
// }

// func main() {
// 	fmt.Println("=== Struct with Pointer ===")
// 	todo := Todo{ID: 1, Title: "Learn Interfaces", Done: false}
// 	fmt.Println("Before:", todo.Done)

// 	markDoneStruct(&todo) // Pass pointer to struct
// 	fmt.Println("After:", todo.Done)

// 	fmt.Println("\n=== Interface Example ===")
// 	email := &EmailNotifier{Email: "victor@example.com"}

// 	// Notice: We pass the pointer, but the function parameter is just "Notifier" (no *)
// 	sendNotification(email, "Your todo has been completed!")

// 	// Bonus: This also works if the method had value receiver, but pointer receivers are more common}
// }
