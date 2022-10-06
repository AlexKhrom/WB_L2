package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//cut -d '-' -f 1,2 -s
//Mers-green-1000
//Bmw-black-990
//Audi-grey-990
//Lambo

func cut() {
	//myscanner := bufio.NewScanner(os.Stdin)
	//myscanner.Scan()
	//4_command := myscanner.Text()

	argsWithProg := os.Args

	args := argsWithProg[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "" {
			args = append(args[:i], args[i+1:]...)
			i--
		}
	}

	sep := "\t"
	var fields []int
	var withSepr bool

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-f":
			i++
			if len(args) < i+1 {
				fmt.Println("should be fields here")
				return
			}
			stringFields := strings.Split(args[i], ",")
			for j := range stringFields {
				intField, err := strconv.Atoi(stringFields[j])
				intField--
				if err != nil {
					fmt.Println("fields should be integer")
					return
				}
				fields = append(fields, intField)
			}
		case "-d":
			i++
			if len(args) < i+1 {
				fmt.Println("should be delimiter here")
				return
			}
			sep = strings.ReplaceAll(args[i], "'", "")
		case "-s":
			withSepr = true
		}
	}

	fmt.Println("enter strings with end - \\end : ")

	var allText [][]string

	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	line := myscanner.Text()

	for line != "\\end" {
		allText = append(allText, strings.Split(line, sep))
		myscanner = bufio.NewScanner(os.Stdin)
		myscanner.Scan()
		line = myscanner.Text()
	}
	fmt.Println("========================\n")

	sort.Ints(fields)
	for _, arr := range allText {
		firstElem := true
		if fields != nil {
			for _, iter := range fields {
				if len(arr) > iter {
					if withSepr {
						if len(arr) > 1 {
							addSepToStr(&firstElem, arr[iter], sep)
						}
					} else {
						addSepToStr(&firstElem, arr[iter], sep)
					}
				}
			}
		} else {
			for _, it := range arr {
				if withSepr {
					if len(arr) > 1 {
						addSepToStr(&firstElem, it, sep)
					}
				} else {
					addSepToStr(&firstElem, it, sep)
				}
			}
		}
		if firstElem != true {
			fmt.Print("\n")
		}
	}
	fmt.Println()
}

func addSepToStr(firstElem *bool, str, sep string) {
	if *firstElem {
		*firstElem = false
		fmt.Print(str)
	} else {
		fmt.Print(sep + str)
	}
}

func main() {
	cut()
}
