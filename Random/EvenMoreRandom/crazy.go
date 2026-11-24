package evenmorerandom

import (
	"errors"
	"fmt"
)

func WowReally() {
	fmt.Println("👉 👌 ❔")
}

func LOLWUT(notzero int) {
	if notzero == 0 {
		err := errors.New("You tucked up")
		fmt.Println(err)
	} else {
		fmt.Println(69 / notzero)
	}
}
