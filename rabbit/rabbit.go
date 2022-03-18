package rabbit

import (
	"fmt"

	"github.com/Assembless/rabbit-problem-go/randbool"
)

type Rabbit struct {
	Holes    int
	Position int
}

func New(holes int, initialPosition int) Rabbit {
	return Rabbit{Holes: holes, Position: initialPosition}
}

func (r *Rabbit) Move() {
	fmt.Println()

	fmt.Printf("Rabbit will move from hole %d \n", r.Position)

	if randbool.RandBool() {
		r.Position += 1
	} else {
		r.Position -= 1
	}

	if r.Position < 0 {
		r.Position = 0
	} else if r.Position > r.Holes {
		r.Position = r.Holes
	}

	fmt.Printf("Rabbit moved to hole %d \n", r.Position)
}

func (r *Rabbit) Shoot(hole int) {
	fmt.Println()

	fmt.Println("Attempting to shoot at hole", hole)

	if hole == r.Position {
		fmt.Println("Rabbit was shot!")
	} else {
		fmt.Println("Rabbit was not shot!")
	}
}
