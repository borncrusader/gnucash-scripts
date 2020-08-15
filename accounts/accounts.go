package accounts

import (
	"encoding/csv"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

type Account struct {
	Name            string
	HeaderRows      []int
	IgnoreAlternate bool
	IgnoreRows      []int
	Filename        string
	Records         [][]string
}

func CreateAccount() *Account {
	return &Account{}
}

func LoadAccounts() []*Account {
	return []*Account{}
}

func (a *Account) ReadRecords(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	c := csv.NewReader(file)

	a.Records, err = c.ReadAll()
	if err != nil {
		return err
	}

	return nil
}

func (a *Account) PrettyPrint() string {
	t := table.NewWriter()

	for idx, records := range a.Records {
		row := make([]interface{}, len(records))
		for i, entry := range records {
			row[i] = entry
		}

		if contains(a.HeaderRows, idx) != -1 {
			t.AppendHeader(row)
		} else {
			t.AppendRow(row)
		}
	}

	return t.Render()
}

func contains(slice []int, value int) int {
	for idx, element := range slice {
		if element == value {
			return idx
		}
	}

	return -1
}
