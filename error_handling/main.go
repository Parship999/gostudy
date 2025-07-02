package main

import "fmt"

func divide(a,b float64) (float64, error) {
	if b==0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a/b, nil
}

func main(){
	fmt.Println("started error handling")
	// ans, err := divide(10,0)
	// if err != nil {
	// 	fmt.Println("Error Handling")
	// }
	// fmt.Println("Div of two no.s is ", ans)

	ans, _ := divide(10, 5)
	fmt.Println("division of two no.s is", ans)
}

// _ is use to hide, that i dont want to handle it/error