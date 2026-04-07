package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Channel carries FUNCTIONS — each function IS a delivery task
	deliveries := make(chan func(), 10)

	var wg sync.WaitGroup
	numRiders := 3

	// Launch 3 riders (workers)
	for i := 1; i <= numRiders; i++ {
		wg.Add(1)

		go func(riderID int) {
			defer wg.Done()

			// Each rider just runs whatever delivery function they receive
			for delivery := range deliveries {
				fmt.Printf("🏍️ Rider %d → Starting a delivery...\n", riderID)
				delivery() // Execute the delivery function
				fmt.Printf("✅ Rider %d → Delivery completed!\n\n", riderID)
			}
		}(i)
	}

	// Order #1: KFC to Rahim
	deliveries <- func() {
		fmt.Println("   📦 Order #1: Pick up from KFC → Deliver to Rahim at Gulshan-1")
		fmt.Println("   → Rider arrives at KFC, picks up order")
		fmt.Println("   → Rider drives to Gulshan-1")
		fmt.Println("   → Hands food to Rahim")
		time.Sleep(1 * time.Second)
	}

	// Order #2: Pizza Hut to Karim
	deliveries <- func() {
		fmt.Println("   📦 Order #2: Pick up from Pizza Hut → Deliver to Karim at Banani-2")
		fmt.Println("   → Rider arrives at Pizza Hut, picks up order")
		fmt.Println("   → Rider drives to Banani-2")
		fmt.Println("   → Hands food to Karim")
		time.Sleep(1 * time.Second)
	}

	// Order #3: BFC to Sumi
	deliveries <- func() {
		fmt.Println("   📦 Order #3: Pick up from BFC → Deliver to Sumi at Dhanmondi-27")
		fmt.Println("   → Rider arrives at BFC, picks up order")
		fmt.Println("   → Rider drives to Dhanmondi-27")
		fmt.Println("   → Hands food to Sumi")
		time.Sleep(1 * time.Second)
	}

	// Order #4: Madchef to Nadia
	deliveries <- func() {
		fmt.Println("   📦 Order #4: Pick up from Madchef → Deliver to Nadia at Uttara-10")
		fmt.Println("   → Rider arrives at Madchef, picks up order")
		fmt.Println("   → Rider drives to Uttara-10")
		fmt.Println("   → Hands food to Nadia")
		time.Sleep(1 * time.Second)
	}

	// Order #5: Takeout to Fahim
	deliveries <- func() {
		fmt.Println("   📦 Order #5: Pick up from Takeout → Deliver to Fahim at Mirpur-12")
		fmt.Println("   → Rider arrives at Takeout, picks up order")
		fmt.Println("   → Rider drives to Mirpur-12")
		fmt.Println("   → Hands food to Fahim")
		time.Sleep(1 * time.Second)
	}

	// Order #6: Chillox to Tania
	deliveries <- func() {
		fmt.Println("   📦 Order #6: Pick up from Chillox → Deliver to Tania at Bashundhara R/A")
		fmt.Println("   → Rider arrives at Chillox, picks up order")
		fmt.Println("   → Rider drives to Bashundhara R/A")
		fmt.Println("   → Hands food to Tania")
		time.Sleep(1 * time.Second)
	}

	// Order #7: Star Kabab to Jisan
	deliveries <- func() {
		fmt.Println("   📦 Order #7: Pick up from Star Kabab → Deliver to Jisan at Mohakhali")
		fmt.Println("   → Rider arrives at Star Kabab, picks up order")
		fmt.Println("   → Rider drives to Mohakhali")
		fmt.Println("   → Hands food to Jisan")
		time.Sleep(1 * time.Second)
	}

	// Order #8: Domino's to Riya
	deliveries <- func() {
		fmt.Println("   📦 Order #8: Pick up from Domino's → Deliver to Riya at Niketon")
		fmt.Println("   → Rider arrives at Domino's, picks up order")
		fmt.Println("   → Rider drives to Niketon")
		fmt.Println("   → Hands food to Riya")
		time.Sleep(1 * time.Second)
	}

	// Order #9: Burger King to Arif
	deliveries <- func() {
		fmt.Println("   📦 Order #9: Pick up from Burger King → Deliver to Arif at Baridhara")
		fmt.Println("   → Rider arrives at Burger King, picks up order")
		fmt.Println("   → Rider drives to Baridhara")
		fmt.Println("   → Hands food to Arif")
		time.Sleep(1 * time.Second)
	}

	// Order #10: Subway to Mila
	deliveries <- func() {
		fmt.Println("   📦 Order #10: Pick up from Subway → Deliver to Mila at Nikunja-2")
		fmt.Println("   → Rider arrives at Subway, picks up order")
		fmt.Println("   → Rider drives to Nikunja-2")
		fmt.Println("   → Hands food to Mila")
		time.Sleep(1 * time.Second)
	}

	// No more orders
	close(deliveries)

	// Wait for all riders to finish
	wg.Wait()

	fmt.Println("🎉 All orders delivered! Riders clocked out.")
}
