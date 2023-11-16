package puppy

import (
	"github.com/alessandrovita/dog"
)

func Bark() string {
	return "Woaf"
}

func Barks() string {
	return "Woaf Woaf!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
