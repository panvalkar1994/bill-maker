package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Enter Customer Name: ", reader)
	b := newBill(name)
	fmt.Println("created bill ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a->add Item, s->save bill, t->give tip) :", reader)
	switch opt {
	case "a":
		itemName, _ := getInput("Enter Item Name: ", reader)
		price, _ := getInput("Enter Item Price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
			return
		}
		b.addItem(itemName, p)
		promptOptions(b)
	case "s":
		fmt.Println("Saving Bill...")
		b.save()
	case "t":
		tip, _ := getInput("Enter tip (in $): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Tip must be a number")
			promptOptions(b)
			return
		}
		b.updateTip(t)
		promptOptions(b)
	default:
		fmt.Println("Invalid")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
	billString := mybill.format()
	fmt.Printf(billString)
}
