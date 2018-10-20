package main

import "fmt"

type Weekday int

const (
	Sunday    Weekday = 0
	Monday    Weekday = 1
	Tuesday   Weekday = 2
	Wednesday Weekday = 3
	Thursday  Weekday = 4
	Friday    Weekday = 5
	Saturday  Weekday = 6
)

func (day Weekday) String() string {
	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}

	if day < Sunday || day > Saturday {
		return "Unknown"
	}

	return names[day]
}

func (day Weekday) Weekend() bool {
	switch day {
	case Friday, Saturday:
		return true

	default:
		return false
	}
}

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)

	fmt.Println(Sunday.Weekend())
	fmt.Println(Monday.Weekend())
	fmt.Println(Tuesday.Weekend())
	fmt.Println(Wednesday.Weekend())
	fmt.Println(Thursday.Weekend())
	fmt.Println(Friday.Weekend())
	fmt.Println(Saturday.Weekend())
}
