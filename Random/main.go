package main

import "random/evenmorerandom"

func main() {
	printArgs()
	evenmorerandom.WowReally()
	args := doStuff()
	anyKwargs()
	evenmorerandom.LOLWUT(0)
	if len(args) > 0 {
		why(args[0])
	} else {
		why(69)
	}	
}
