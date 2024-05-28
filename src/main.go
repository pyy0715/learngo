package main

import (
	"fmt"
	"strings"

	"github.com/pyy0715/learngo/accounts"
	mydicts "github.com/pyy0715/learngo/mydict"
)

func main() {
	// Demonstrates basic arithmetic function
	result := multiplyFunc(2, 3)
	fmt.Println("--- Result of multiplication:", result, "---")

	// Repeats and prints words
	fmt.Println("--- Repeating words: ---")
	repeatWordsFunc("Hello", "World")

	// Calculates length and converts to uppercase
	length, uppercase := lenAndUpperFunc("hello world!")
	fmt.Println("--- Length and uppercase version:", length, uppercase, "---")

	// Computes sum of a range of numbers
	totalSum, totalSum2 := rangeFunc(1, 2, 3, 4, 5)
	fmt.Println("--- Total sums of numbers:", totalSum, totalSum2, "---")

	// Conditional checks using if-else and switch
	fmt.Println("--- If-else check result:", ifelseFunc(18), "---")
	fmt.Println("--- Switch check result:", switchFunc(18), "---")

	// Demonstrates various data structures and pointer usage
	fmt.Println("--- Demonstrating pointer usage: ---")
	pointerFunc()
	fmt.Println("--- Demonstrating array usage: ---")
	arrayFunc()
	fmt.Println("--- Demonstrating slice usage: ---")
	sliceFunc()
	fmt.Println("--- Demonstrating map usage: ---")
	mapFunc()
	fmt.Println("--- Demonstrating struct usage: ---")
	structFunc()

	// Operations on a custom account type from accounts package
	account := accounts.NewAccount("nico")
	account.Deposit(100)
	fmt.Println("--- Account balance after deposit:", account.Balance(), "---")
	err := account.Withdraw(2000)
	if err != nil {
		fmt.Println("Error during withdrawal:", err)
	}
	account.Withdraw(20)
	account.ChangeOwner("yyeon2")
	fmt.Println("Account balance and owner after changes:", account.Balance(), account.Owner())
	fmt.Println("Account details:", account)

	// Demonstrates usage of custom dictionary type from mydicts package
	dictionary := mydicts.Dictionary{"Hello": "First"}
	definition, err := dictionary.Search("Hello")
	if err != nil {
		fmt.Println("--- Error during dictionary search:", err, "---")
	} else {
		fmt.Println("Definition found:", definition)
	}

	// Attempt to add a new word that already exists to check error handling
	err2 := dictionary.Add("Hello", "Second")
	if err2 != nil {
		fmt.Println("--- Error during dictionary addition:", err2, "---")
	} else {
		fmt.Println("Definition added:", definition)
	}

	// Update an existing word in the dictionary
	err3 := dictionary.Update("Hello", "Third")
	if err3 != nil {
		fmt.Println("--- Error during dictionary update:", err3, "---")
	} else {
		definition, _ = dictionary.Search("Hello") // Reuse the variable 'definition' to fetch updated value
		fmt.Println("Definition updated:", definition)
	}

	// Attempt to update a non-existing word to check error handling
	err4 := dictionary.Update("World", "Fourth")
	if err4 != nil {
		fmt.Println("Error during dictionary update:", err4)
	}

	// Attempt to delete a non-existing word to check error handling
	err5 := dictionary.Delete("World")
	if err5 != nil {
		fmt.Println("Error during dictionary deletion:", err5)
	}
}

func multiplyFunc(a, b int) int {
	return a * b
}

func repeatWordsFunc(words ...string) {
	for i, word := range words {
		fmt.Println(i, word)
	}
}

func lenAndUpperFunc(s string) (int, string) {
	defer fmt.Println("deferred")
	return len(s), strings.ToUpper(s)
}

func rangeFunc(numbers ...int) (int, int) {
	totalSum, totalSum2 := 0, 0
	for _, number := range numbers {
		totalSum += number
	}
	for _, number := range numbers {
		totalSum2 += number
	}
	return totalSum, totalSum2
}

func ifelseFunc(age int) bool {
	if age < 18 {
		return true
	} else {
		return false
	}
}

func switchFunc(age int) bool {
	switch age {
	case 18:
		return true
	default:
		return false
	}
}

func pointerFunc() {
	a := 2
	b := a // b is a copy of a
	a = 10 // a is now 10
	c := &a // c is a pointer to a
	d := &a // d is a pointer to a
	*d = 100 // dereference d and set it to 100
	fmt.Println(a, &a, &b, c, *c, d, *d)
}

func arrayFunc() {
	names := [5]string{"nico", "lynn", "dal", "maru"}
	names[3] = "japan"
	fmt.Println(names)
}

func sliceFunc() {
	names := []string{"nico", "lynn", "dal", "maru"}
	names = append(names, "japan")
	fmt.Println(names)
}

func mapFunc() {
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)
	for key, value := range nico {
		fmt.Println(key, value)
	}
}

type person struct {
	name string
	age  int
	favFood []string
}

func structFunc() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{"nico", 18, favFood}
	fmt.Println(nico)
}

