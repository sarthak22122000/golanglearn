package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var order string
	var quantity int
	var sales = []int{}
	counter := 0
	for true {
		fmt.Println("**Welcome To XYZ Cafe**")
		fmt.Println("*******MENU*******")
		fmt.Print("C: coffee, 40rs\n D: dosa, 80 rs\n T: tomato soup, 20rs\n I : idli , 20rs\n V : vada, 25rs.\n B: bhature &chhole, 30rs\n P: paneer pakoda, 30rs\n M: manchurian, 90rs\n H: hakka noodle, 70rs.\n F: French fries, 30rs\n J: jalebi, 30rs\n L: lemonade, 15rs\n S: spring roll, 20rs\n")
		fmt.Println("*_*_*_*_*_*_*_*_*_*_")
		fmt.Println("Press END For closing the day.")
		fmt.Println("Please place the order: ")
		fmt.Println("What You want(Enter Food Code)")
		fmt.Scan(&order)
		order = strings.ToUpper(order)

		if order == "END" {
			total_income(sales)
		} else {
			fmt.Println("Enter Quantity")
			fmt.Scan(&quantity)
		}
		bill := bill(quantity, order)
		fmt.Println("####################")
		fmt.Println("Total bill: ", bill)
		fmt.Println("####################")
		sales = append(sales, bill)
		counter++

	}
}

func total_income(sale []int) {
	var sum int = 0

	for i := 0; i < len(sale); i++ {
		sum = sale[i] + sum
	}
	fmt.Println("Total Income for the day is : ", sum)
	os.Exit(0)
}

func bill(quantity int, order string) int {
	m := map[string]int{
		"C": 40,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	billy := quantity * m[order]
	return billy
}
