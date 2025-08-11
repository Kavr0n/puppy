package puppy

import (
	"github.com/Kavr0n/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof Woof!"
}

func BigBark() string {
	return dog.WhenGrowUp()
}
