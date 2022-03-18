package randbool

import (
	"math/rand"
	"time"
)

/*
RandBool
    This function returns a random boolean value based on the current time
*/
func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}
