package main

import (
	"fmt"
	"math/rand"

	"github.com/Assembless/rabbit-problem-go/rabbit"

	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Println("Rabbit Problem - Made by Assembless")

	c := cron.New()

	// Add your algorithm here:
	randomAlgorithm(c)

	// Start the lifecycle.
	c.Start()

	// Prevents the window from closing / waiting for input.
	fmt.Scanf("h")
	fmt.Println("Bye")
	c.Stop()
}

func randomAlgorithm(c *cron.Cron) {
	rabbitInstance := rabbit.New(10, 5)

	c.AddFunc("@every 5s", func() {
		rabbitInstance.Move()

		nextHole := rand.Intn(10-5) + 5

		rabbitInstance.Shoot(nextHole)
	})
}

/**

// fmt.Printf("Hello, world.\n %s", "Go")

// cards := []string{"Ace of Spades", "Ace of Spades2", "Ace of Spades3", "Ace of Spades4"}

// for i, card := range cards {
// 	fmt.Println(i, card)
// }

// fmt.Println(cards[:1]) // Till the first element of the array.
// fmt.Println(cards[1:]) // From the second element of the array to the end.

// newCards := append(cards, "Ace of Diamonds") // Appends an element to the end of the array.
// fmt.Println(newCards)                        //

// // For loop
// for i := 0; i < 10; i++ {
// 	fmt.Println("Hi !")
// }

// m := map[string]int{
// 	"cow":  1,
// 	"bird": 5,
// 	"dick": 8,
// }

// rabbit := New(10, 0)
// rabbit.move()

// fmt.Println()

// ballsCandidate := "bird"

// if m[ballsCandidate] == 1 {
// 	fmt.Printf("My %s has %d ball", ballsCandidate, m[ballsCandidate])
// } else {
// 	fmt.Printf("My %s has %d balls", ballsCandidate, m[ballsCandidate])
// }
*/
