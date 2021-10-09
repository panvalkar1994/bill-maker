package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b *bill) format() string {
	billString := ""
	billString += fmt.Sprintf("Name: %v\n", b.name)
	billString += fmt.Sprintln("----------------------------------")
	var total float64 = 0
	for k, v := range b.items {
		billString += fmt.Sprintf("%-25v %.2f\n", k+":", v)
		total += v
	}
	billString += fmt.Sprintln("----------------------------------")
	billString += fmt.Sprintf("%-25v %0.2f\n", "Tip", b.tip)
	total += b.tip
	billString += fmt.Sprintf("%-25v %0.2f", "Total", total)
	return billString
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(item string, price float64) {
	b.items[item] = price
}

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill Saved")
}
