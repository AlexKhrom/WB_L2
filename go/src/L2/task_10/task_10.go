package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type webServer struct {
	host    string
	port    string
	timeout time.Duration
}

func getArguments() *webServer {
	argsWithProg := os.Args

	args := argsWithProg[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "" {
			args = append(args[:i], args[i+1:]...)
			i--
		}
	}
	fmt.Println("args = ", args)

	var server webServer
	if len(args) < 3 {
		server.host = args[0]
		server.port = args[1]
		server.timeout = time.Second * 10
	} else {
		var i = 0
		for j, arg := range args {
			sepArgs := strings.Split(arg, "=")
			if sepArgs[0] == "--timeout" {
				if len(sepArgs) < 1 {
					fmt.Println("err witn --timeout argument")
					return nil
				}
				timeout := strings.Split(sepArgs[1], "s")[0]
				timeoutInt, err1 := strconv.Atoi(timeout)
				if err1 != nil {
					fmt.Println("err witn --timeout argument")
					return nil
				}
				i = j
				server.timeout = time.Second * time.Duration(timeoutInt)
			}
		}

		for j, arg := range args {
			if i != j {
				if server.host == "" {
					server.host = arg
				} else {
					server.port = arg
				}
			}
		}
	}

	return &server

}

func telnet() {
	serverInfo := getArguments()
	fmt.Println(*serverInfo)

	strEcho := "GET /api/events_for_day HTTP/1.1\nHost: 127.0.0.1"
	servAddr := serverInfo.host + ":" + serverInfo.port

	time.Sleep(serverInfo.timeout)
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	done := make(chan bool, 1)

	go func() {
		defer close(done)
		defer conn.Close()

		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			str := sc.Text()
			fmt.Println("str = ", str)

			httpRequest := "GET / HTTP/1.1\nHost: localhost\n\n"

			_, err = conn.Write([]byte(httpRequest))
			if err != nil {
				println("Write to server failed:", err.Error())
				os.Exit(1)
			}

			println("write to server = ", strEcho)

			fmt.Println("response\n//////////////////////////////////////////")

			reply := make([]byte, 1024)

			_, err = conn.Read(reply)
			if err != nil {
				println("Write to server failed:", err.Error())
				os.Exit(1)
			}

			println("reply from server=", string(reply))
			time.Sleep(time.Second)
			fmt.Println("//////////////////////////////////////////")

		}

	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("done")
}

func main() {
	telnet()
}
