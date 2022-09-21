package main

import (
	"bufio"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"log"
	"os"
	_ "os/exec"
	"strconv"
	"strings"
	"syscall"
)

func cd(args []string) {
	if len(args) < 2 {
		fmt.Println("directory shouldn't be empty")
		return
	}
	if err := os.Chdir(args[1]); err != nil {
		fmt.Printf("Error:\tCould not move into the directory (%s)\n", args[1])
	}
	return
}

func pwd(args []string) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println(path)
}

func echo(args []string) {
	if len(args) < 2 {
		fmt.Println("")
	}
	fmt.Println(strings.Join(args[1:], " "))

}

func kill(args []string) {
	if len(args) < 2 {
		fmt.Println("pid shouldn't be empty")
		return
	}
	pid, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("pid should be integer")
		return
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		log.Println(err)
		return
	}
	// Kill the process
	err = proc.Kill()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("process with id=%d successfully killed\n", pid)
}

func psList(args []string) {
	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		return
	}

	// map ages
	for x := range processList {
		var process ps.Process
		process = processList[x]
		log.Printf("%d\t%s\n", process.Pid(), process.Executable())

		// do os.* stuff on the pid
	}
}

func execFunc(args []string) {

	if len(args) < 2 {
		fmt.Println("exec file shouldn't be empty")
		return
	}
	pid, err := syscall.ForkExec(args[1], args[2:], nil)

	if err != nil {
		fmt.Println("err exec = ", err)
		return
	}

	fmt.Println("new pid = ", pid)
	//cmd := exec.Command(args[1])
	//
	//err = cmd.Run()
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("exec is ok")
}

func shell() {

	var command string
	pid := os.Getpid()

	fmt.Println("pid = ", pid)

	fmt.Println("shell is ready for work, enter your commands :")
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	command = myscanner.Text()

	for command != "\\quit" {
		fmt.Println("command = ", command)
		args := strings.Split(command, " ")

		for i := 0; i < len(args); i++ {
			if args[i] == "" {
				args = append(args[:i], args[i+1:]...)
				i--
			}
		}
		fmt.Println("args = ", args)

		switch args[0] {
		case "cd":
			cd(args)
		case "pwd":
			pwd(args)
		case "echo":
			echo(args)
		case "kill":
			kill(args)
		case "ps":
			psList(args)
		case "exec":
			execFunc(args)
		default:
			fmt.Printf("bad command : '%s'\n", args[0])
		}
		myscanner = bufio.NewScanner(os.Stdin)
		myscanner.Scan()
		command = myscanner.Text()
	}

	fmt.Println("shell completes the work")
}

func main() {
	shell()
}
