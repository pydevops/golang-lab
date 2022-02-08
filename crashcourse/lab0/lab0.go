package main

import "fmt"

// type User struct {
// 	name  string `json:"firstName"`
// 	id    int
// 	email Email
// }

// type UserMap map[int]User

// type Email string

func main() {
	// slice
	// collection := []int{1, 2, 3, 4, 5}
	// collection[1] = 5
	// collection = append(collection, 10)
	// map
	// collection := make(map[string]int)
	// collection["hello"] = 1
	// collection["world"] = 2
	// delete(collection, "hello")
	// 	collection := map[string]int{
	// 		"hello": 3,
	// 		"world": 2,
	// 	}
	// if
	// 	fmt.Println(collection)
	// 	if collection["hello"] > 3 {
	// 		fmt.Println("greater than 3")
	// 	} else if collection["hello"] > 2 {
	// 		fmt.Println("greater than 2")
	// 	} else {
	// 		fmt.Println("else")
	// 	}
	// loop
	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// range over a slice
	// x := []int{0, 1, 2, 3, 4, 5}
	// for _, value := range x {
	// 	fmt.Println("Value:", value)
	// }
	// range over a map
	// x := map[string]int{"one": 1, "two": 2}
	// for key, value := range x {
	// 	fmt.Println("key:", key, "Value:", value)
	// }
	// infinite for loop
	// i := 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// 	if i > 5 {
	// 		break
	// 	}
	// }
	// i := -5.0
	// x, err := sqrt(i)
	// if err != nil {
	// 	fmt.Printf("There was an error: %s", err)
	// }
	// fmt.Println(x)

	// myUser := User{
	// 	name:  "Bob",
	// 	id:    5,
	// 	email: "bob@bob.com",
	// }
	// myUser.name = "Dole"
	// if myUser.email.validate() {
	// 	fmt.Println("This is valid email")
	// } else {
	// 	fmt.Println("This is invalid  email")
	// }

	// fmt.Println(myUser)

	x := 0
	inc(&x)
	fmt.Println(x)
}

func inc(x *int) {
	*x++
}

// func (e Email) validate() bool {
// 	return true
// }

// func sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, fmt.Errorf("The value of x %F cannot be less than 0", x)
// 	}
// 	return math.Sqrt(x), nil
// }
