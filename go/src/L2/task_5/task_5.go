package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func grep() {
	argsWithProg := os.Args

	args := argsWithProg[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "" {
			args = append(args[:i], args[i+1:]...)
			i--
		}
	}

	args = args[1:]
	//fmt.Println(args)

	var aN, bN, cN int
	var isCount, ignoreCase, invert, fixed, lineNum bool
	var err error
	var template, fileName string

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-A":
			i++
			if len(args) < i+1 {
				fmt.Println("should be fields here")
				return
			}
			aN, err = strconv.Atoi(args[i])
			if err != nil {
				fmt.Println("err with -A convInt")
				return
			}
		case "-B":
			i++
			if len(args) < i+1 {
				fmt.Println("should be fields here")
				return
			}
			bN, err = strconv.Atoi(args[i])
			if err != nil {
				fmt.Println("err with -B convInt")
				return
			}
		case "-C":
			i++
			if len(args) < i+1 {
				fmt.Println("should be fields here")
				return
			}
			cN, err = strconv.Atoi(args[i])
			if err != nil {
				fmt.Println("err with -C convInt")
				return
			}
		case "-c":
			isCount = true
		case "-i":
			ignoreCase = true
		case "-v":
			invert = true
		case "-F":
			fixed = true
		case "-n":
			lineNum = true

		default:
			if template == "" {
				template = args[i]
			} else {
				fileName = args[i]
			}
		}
	}

	var text []string
	var lines []int

	if fileName != "" {
		lines, text = findWordFile(fileName, template, ignoreCase, fixed, lineNum, invert)
	} else {
		lines, text = findWordTerminal(template, ignoreCase, fixed, lineNum, invert)
	}

	//fmt.Println(aN, bN, cN)
	//fmt.Println(isCount, ignoreCase, invert, fixed, lineNum)
	//fmt.Println(template, fileName)
	//
	//fmt.Println("text = ", text)
	//fmt.Println("lines = ", lines)

	//for _, str := range text {
	//	fmt.Println("line = ", []byte(str))
	//}

	printResult(aN, bN, cN, isCount, lines, text)

}

func findWordTerminal(tmp string, ignoreCase, fixed, lineNum, invert bool) ([]int, []string) {

	if ignoreCase {
		tmp = strings.ToLower(tmp)
	}

	var resultLines []int
	var textRes []string

	var j = 1

	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var str string

		str = scanner.Text()

		if str == "/end" {
			break
		}

		textRes = append(textRes, str)

		if ignoreCase {
			str = strings.ToLower(str)
		}

		if fixed {
			if invert {
				if str != tmp {
					if lineNum {
						fmt.Println("line num = ", j)
					}
					resultLines = append(resultLines, j-1)
				}
			} else {
				if str == tmp {
					if lineNum {
						fmt.Println("line num = ", j)
					}
					resultLines = append(resultLines, j-1)
				}
			}
		} else {
			strs := strings.Split(str, " ")
			for _, s := range strs {
				if invert {
					if s != tmp {
						if lineNum {
							fmt.Println("line num = ", j)
						}
						resultLines = append(resultLines, j-1)
					}
				} else {
					if s == tmp {
						if lineNum {
							fmt.Println("line num = ", j)
						}
						resultLines = append(resultLines, j-1)
					}
				}
			}
		}
		j++
	}

	return resultLines, textRes
}

func findWordFile(fileName, tmp string, ignoreCase, fixed, lineNum, invert bool) ([]int, []string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var textRes []string

	if ignoreCase {
		tmp = strings.ToLower(tmp)
	}

	var resultLines []int

	var j = 1

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		var str string
		//n, err := file.Read(data)
		//if err == io.EOF { // если конец файла
		//	break // выходим из цикла
		//}
		textRes = append(textRes, fileScanner.Text())

		str = fileScanner.Text()
		if ignoreCase {
			str = strings.ToLower(str)
		}

		if fixed {
			if invert {
				if str != tmp {
					if lineNum {
						fmt.Println("line num = ", j)
					}
					resultLines = append(resultLines, j-1)
				}
			} else {
				if str == tmp {
					if lineNum {
						fmt.Println("line num = ", j)
					}
					resultLines = append(resultLines, j-1)
				}
			}
		} else {
			strs := strings.Split(str, " ")
			for _, s := range strs {
				if invert {
					if s != tmp {
						if lineNum {
							fmt.Println("line num = ", j)
						}
						resultLines = append(resultLines, j-1)
					}
				} else {
					if s == tmp {
						if lineNum {
							fmt.Println("line num = ", j)
						}
						resultLines = append(resultLines, j-1)
					}
				}
			}
		}
		j++
	}

	return resultLines, textRes
}

func printResult(aN, bN, cN int, count bool, lines []int, text []string) {
	fmt.Println("/////////////////////////")
	if count {
		fmt.Println(len(lines))
		return
	}

	for i, str := range text {
		if isLine(aN, bN, cN, i, lines) {
			fmt.Println(str)
		}
	}
}

func isLine(aN, bN, cN, iter int, lines []int) bool {
	for _, line := range lines {
		if line == iter {
			return true
		}
		if line-aN <= iter && iter <= line && aN != 0 {
			return true
		}
		if line <= iter && iter <= line+bN && bN != 0 {
			return true
		}
		if line-cN <= iter && iter <= line+cN && cN != 0 {
			return true
		}
	}
	return false
}

func main() {
	grep()
}
