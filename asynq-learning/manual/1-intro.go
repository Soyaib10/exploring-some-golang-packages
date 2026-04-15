package main

import (
	"fmt"
	"time"
)

func sendWelcomeEmail(userID int) {
	// ধরো এটা email পাঠানোর কাজ, ২ সেকেন্ড লাগে
	time.Sleep(2 * time.Second)
	fmt.Printf("Email sent to user %d\n", userID)
}

func handleRegister(userID int) {
	fmt.Printf("User %d registered, sending email in background...\n", userID)
	go sendWelcomeEmail(userID)
	fmt.Printf("Response sent to user %d immediately\n", userID)
}

func main() {
	// হঠাৎ ১০০ জন user একসাথে register করল
	for i := 1; i <= 100; i++ {
		handleRegister(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Program done")
}
