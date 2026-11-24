package main

import "fmt"

func sumOf(a, b int) int {
	return a + b
}
func diffOf(a, b int) int {
	return a - b
}
func ops(a, b int, sumFunc func(int, int) int, diffFunc func(int, int) int) {
	fmt.Printf("%d + %d = %d\n", a, b, sumFunc(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, diffFunc(a, b))
}
func gibberish() (x int, int, z int) {
	return x, 2, z
}
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
func goodbye(name string) {
	fmt.Printf("Goodbye, %s!\n", name)
}
func switchStuff(value int, name string) string {
	greet(name)
	defer goodbye(name)
	switch value {
	case 0:
		return "None"
	case 1, 2:
		return "Not much"
	case 3, 4, 5:
		return "Pretty much"
	default:
		fmt.Println("Wow, really?")
		return "Very much"
	}
}
func main() {
	var myName string = "Timmy"
	var myWeight float64 = 1000.0
	var myAge int = 420
	// var text string = fmt.Sprintf("My name is %s, I am %d years old and I weigh %.2f kilogramms.", myName, myAge, myWeight)
	fmt.Printf("My name is %s, I am %d years old and I weigh %.2f kilogramms.\n", myName, myAge, myWeight)
	num1 := 4
	num2 := 5
	fmt.Printf("%d + %d = %d\n", num1, num2, sumOf(num1, num2))
	_, y, z := gibberish()
	fmt.Println(y, z)
	ops(num1, num2, sumOf, diffOf)
	ops(num1, num2, func(a, b int) int {return a + b}, func(a, b int) int {return a - b})
	fmt.Println(switchStuff(69, myName))
	for i := 0; i <= 9; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	for i, r := range myName {
		fmt.Printf("[%d]\t%s\t(%d)\n", i, string(r), r)
	}
}
