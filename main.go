package main

import (
	"bufio"
	"fmt"
	"os"

)

func main()  {

	fmt.Println("Hello World!!")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("How Old are You?")
	scanner.Scan()
	age := 12
	//age := 18"%q scanner.Text()"
	if age >= 18 {
		fmt.Println("You can Ride!!")
	}else if age >= 14 {
		fmt.Println("You can Ride!!, with your Faro!")
	} else {
		fmt.Println("Commot for here Joor, you dey ment??")
		fmt.Printf("come back after %d year(s) okponu!!", 18-age)
	}

}