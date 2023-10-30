package main

import (
	"calculator/calculator"
	"fmt"
)

func main() {
	// how to decalre variable with default value
	// if didnt set default value it will set default by auto ex.int = 0, string = "", bool = false
	var greet_1 string
	greet_1 = "hello from greet 1"

	var greet_2 string = "hello from greet 2"

	greet_3 := "hello from greet 3"

	fmt.Println(greet_1)
	fmt.Printf("try to print using %v and %s\n", greet_2, greet_3)

	year := 2000
	month := 3
	date := 30

	date_string := fmt.Sprintf("%v / %v / %v", year, month, date)
	fmt.Printf("display date_string %s\n", date_string)

	// flow control
	if 2 > 1 {
		fmt.Print("2 > 1, its true!")
	} else {
		fmt.Print("impossible ja")
	}

	for i := 0; i < 5; i++ { // if didnt set condition its mean while(true)
		fmt.Printf("i in for loop %v\n", i)
	}

	switch year {
	case 1999:
		fmt.Println("im not born yet")
	case 2000:
		fmt.Println("birth!!")
	}

	// vector variable
	// array => fixed lenge
	var cats [2]string
	cats[0] = "platoo"
	cats[1] = "pop"

	fruits := [3]string{"mamuang", "mango", "mangie"}

	fmt.Printf("cats length is %v, fruits length is %v\n", len(cats), len(fruits))

	// slice => like array but dynamic size
	var snacks []string
	snacks = append(snacks, "cookies")
	snacks = append(snacks, "jelly")

	drinks := []string{"coke", "coffee", "tea", "ovaltin"}

	// using loop with vector
	for i := 0; i < len(snacks); i++ {
		fmt.Printf("snack => %v\n", snacks[i])
	}

	for index, drink := range drinks {
		fmt.Printf("drink => %v index: %v\n", drink, index)
	}

	// using function
	r, err := calculator.Divide(10, 5)
	if (err) != nil {
		// handle error
		panic(err) // can use panic() or os.Exit(1)
	}
	fmt.Printf("result from divide function => %v\n", r)

}

func sayHi(name string) {
	fmt.Printf("hihi ja, khun %v\n", name)
}

func multiTypeReturn() (int, string, bool) {
	return 123, "hi in func", true
}
