package main

/* Imports */
import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Item struct {
	name  string
	price float64
}

type ReturnInformation struct {
	formattedTime string
	path          string
	tax           float64
	afterTax      float64
	beforeTax     float64
}

var billsPath string = "Bills"

func main() {
	// Add a bill with 3 items, one counting 500 items
	// under the name of "John Doe", with a 6% sales tax

	AddBill(
		"John Doe",
		[]Item{
			Item{
				"Chips & Queso",
				50,
			},
			Item{
				"Hungry Meal w/ Coca Cola",
				50.2,
			},
			Item{
				"Sushi Soup x500",
				225.30,
			},
		},
		1.06,
	)
}

func AddBill(name string, items []Item, tax float64) ReturnInformation {
	time := time.Now()

	var formattedTime string = fmt.Sprint(
		strconv.Itoa(time.Year()) + "-" +
			time.Month().String() + "-" +
			strconv.Itoa(time.Day()) + " " +
			strconv.Itoa(time.Hour()) + "." +
			strconv.Itoa(time.Minute()) + "." +
			strconv.Itoa(time.Second()),
	)

	var total float64 = 0

	var formattedBills = fmt.Sprint(
		name+" at "+formattedTime,
		"\n\n",
	)

	for i := 0; i < len(items); i++ {
		formattedBills +=
			items[i].name +
				"\n" + strconv.FormatFloat(items[i].price, 'f', 2, 64) +
				"\n\n"

		total += items[i].price
	}

	formattedBills += "Total before tax: " + strconv.FormatFloat(total, 'f', 2, 64) + "\n"
	formattedBills += "Total after tax: " + strconv.FormatFloat(total*tax, 'f', 2, 64)

	var path = billsPath +
		"/" +
		name +
		" " +
		formattedTime +
		".txt"

	os.WriteFile(
		path,
		[]byte(formattedBills),
		0644,
	)

	return ReturnInformation{
		formattedTime,
		path,
		tax,
		total * tax,
		total,
	}
}
