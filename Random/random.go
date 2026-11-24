package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
	// "syscall"
)

func printArgs() {
	fmt.Println("Args:")
	length := 0
	for _, arg := range os.Args {
		if runeCount := utf8.RuneCountInString(arg); runeCount > length {
			length = runeCount
		}
	}
	for i, arg := range os.Args {
		paddedArg := fmt.Sprintf("%-*s", length, arg)
		fmt.Printf("Arg %d:\t%s (%T)\n", i+1, paddedArg, arg)
	}
}

func doStuff() []int {
	doubler := func(a int) int { return a * 2 }
	validArgs := []int{}
	if len(os.Args) == 1 {
		fmt.Println("No funny arguments were passed")
		return []int{}
	}
	for _, arg := range os.Args {
		if num, err := strconv.Atoi(arg); err == nil {
			validArgs = append(validArgs, num)
		}
	}
	if len(validArgs) == 0 {
		fmt.Println("There were arguments, but they all sukked")
		return []int{}
	}
	for _, num := range validArgs {
		fmt.Printf("%d -> %d\n", num, doubler(num))
	}
	return validArgs
}

func mapPrint(m map[string]string) {
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func anyKwargs() {
	kwargs := make(map[string]string)
	for _, arg := range os.Args[1:] {
		key := strings.TrimLeft(arg, "-")
		if key, val, found := strings.Cut(key, "="); found {
			kwargs[key] = val
		}
	}
	if len(kwargs) == 0 {
		fmt.Println("NO KWARGS???")
	} else {
		fmt.Println("Your kwarg(s):")
		mapPrint(kwargs)
	}
}

func why(n int) {
	a := func() int {
		if n > 0 {
			return n
		} else {
			return -n
		}
	}()
	fmt.Printf("a = %d\n", a)
}
