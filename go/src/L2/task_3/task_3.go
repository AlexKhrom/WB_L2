package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// go run task_3.go text.txt -r -u -k 2
// sort text.txt -r -u -k 2
func readCommand(command string) (string, int, bool, bool, bool) {
	var fileName = ""
	var sortColumn = -1
	var sortDigit = false
	var sortReverse = false
	var sortUniq = false
	var err error
	args := strings.Split(command, " ")

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-k":
			i++
			if len(args) < i+1 {
				fmt.Println("should be fields here")
				return "", -1, false, false, false
			}
			sortColumn, err = strconv.Atoi(args[i])
			if err != nil {
				fmt.Println("column for sort should be integer")
				return "", -1, false, false, false
			}
		case "-n":
			sortDigit = true
		case "-r":
			sortReverse = true
		case "-u":
			sortUniq = true
		default:
			fileName = args[i]
		}
	}

	return fileName, sortColumn, sortDigit, sortReverse, sortUniq
}

func readFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 512)
	var textByte []byte

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		textByte = append(textByte, data[:n]...)
	}

	textStrings := strings.Split(string(textByte), "\n")
	var text [][]string

	for _, str := range textStrings {
		strSlice := strings.Split(str, " ")
		//strSlice = deleteSpace(strSlice)
		text = append(text, strSlice)
	}
	return text
}

func deleteSpace(sl []string) []string {
	for i := 0; i < len(sl); i++ {
		if sl[i] == "" {
			sl = append(sl[:i], sl[i+1:]...)
			i--
		}
	}
	return sl
}

func sortArr(sl [][]string, column int, digit bool) [][]string {
	if column != -1 {
		if digit {
			sort.Slice(sl, func(i, j int) bool {
				floatI, _ := strconv.ParseFloat(sl[i][column], 32)
				floatJ, _ := strconv.ParseFloat(sl[j][column], 32)
				if floatI < floatJ {
					return true
				}
				return false
			})
		} else {
			sort.Slice(sl, func(i, j int) bool {
				if sl[i][column] < sl[j][column] {
					return true
				}
				return false
			})
		}
	} else {
		if digit {
			sort.Slice(sl, func(i, j int) bool {
				floatI, _ := strconv.ParseFloat(sl[i][0], 32)
				floatJ, _ := strconv.ParseFloat(sl[j][0], 32)
				if floatI < floatJ {
					return true
				}
				return false
			})
		} else {
			sort.Slice(sl, func(i, j int) bool {
				if strings.Join(sl[i], " ") < strings.Join(sl[j], " ") {
					return true
				}
				return false
			})
		}
	}
	return sl
}

func makeUniqStrings(text [][]string, uniq bool) [][]string {
	if uniq {
		var result [][]string
		temp := map[string]struct{}{}
		for _, item := range text {
			if _, ok := temp[strings.Join(item, " ")]; !ok {
				temp[strings.Join(item, " ")] = struct{}{}
				result = append(result, item)
			}
		}
		return result
	}
	return text
}

func reverseSl(text [][]string, revers bool) [][]string {
	if revers {
		for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {
			text[i], text[j] = text[j], text[i]
		}
	}
	return text
}

func mySort() {
	argsWithProg := os.Args
	fmt.Println(argsWithProg[1:])

	//myscanner := bufio.NewScanner(os.Stdin)
	//myscanner.Scan()
	//command := myscanner.Text()

	fileName, sortColumn, sortDigit, sortReverse, sortUniq := readCommand(strings.Join(argsWithProg[1:], " "))
	if sortColumn != -1 {
		sortColumn--
	}
	fmt.Println(sortColumn, sortDigit, sortReverse, sortUniq)

	text := readFile(fileName)
	text = sortArr(text, sortColumn, sortDigit)
	text = makeUniqStrings(text, sortUniq)
	text = reverseSl(text, sortReverse)

	fmt.Println("res text = ", text)

	fileRes, err := os.Create("result.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer fileRes.Close()

	for _, str := range text {
		fileRes.WriteString(strings.Join(str, " ") + "\n")
	}

	fmt.Println("Done.")
}

func main() {
	mySort()
}
