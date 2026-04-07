package main

import (
	"fmt"
	"sync"
	"time"
)

// A real delivery order
type Order struct {
	ID         int
	Restaurant string
	Customer   string
	Address    string
}

func main() {
	// A queue of orders waiting for riders
	orders := make(chan Order, 10)

	// We have 3 riders on duty
	var wg sync.WaitGroup
	numRiders := 3

	// Launch 3 riders (workers)
	for i := 1; i <= numRiders; i++ {
		wg.Add(1)

		go func(riderID int) {
			defer wg.Done()

			// Each rider keeps taking orders from the queue
			for order := range orders {
				fmt.Printf("🏍️ Rider %d → Pick up from %s → Deliver to %s at %s\n",
					riderID, order.Restaurant, order.Customer, order.Address)

				// Simulate delivery taking time
				time.Sleep(1 * time.Second)

				fmt.Printf("✅ Rider %d → Order #%d delivered!\n", riderID, order.ID)
			}
		}(i)
	}

	// 10 customer orders come in
	allOrders := []Order{
		{1, "KFC", "Rahim", "Gulshan-1"},
		{2, "Pizza Hut", "Karim", "Banani-2"},
		{3, "BFC", "Sumi", "Dhanmondi-27"},
		{4, "Madchef", "Nadia", "Uttara-10"},
		{5, "Takeout", "Fahim", "Mirpur-12"},
		{6, "Chillox", "Tania", "Bashundhara R/A"},
		{7, "Star Kabab", "Jisan", "Mohakhali"},
		{8, "Domino's", "Riya", "Niketon"},
		{9, "Burger King", "Arif", "Baridhara"},
		{10, "Subway", "Mila", "Nikunja-2"},
	}

	for _, order := range allOrders {
		orders <- order
	}

	// No more orders coming in
	close(orders)

	// Wait for all riders to finish all deliveries
	wg.Wait()

	fmt.Println("\n🎉 All orders delivered! Riders clocked out.")
}
