package main

import (
	"fmt"

	"github.com/borncrusader/gnucash-scripts/accounts"
	"github.com/borncrusader/gnucash-scripts/matchers"
)

func main() {
	a := accounts.Account{
		Name:       "CitiBank",
		HeaderRows: []int{0},
	}

	err := a.ReadRecords("./samples/CitiBank.csv")
	if err != nil {
		fmt.Println(err)
	}

	matchers.FindMatchers("AUTOPAY 200201113930890RAUTOPAY AUTO-PMT")

	fmt.Print(a.PrettyPrint())
}
