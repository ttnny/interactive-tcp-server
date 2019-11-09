package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":888")
	if err != nil {
		log.Panic(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Display main menu right after connection established
	displayMenu(conn)

	// Set connection timeout
	err := conn.SetDeadline(time.Now().Add(30 * time.Second))
	if err != nil {
		log.Println("Connection timed out!")
	}

	// Get input from user
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		in := scanner.Text()

		switch in {
		case "1":
			conn.Write([]byte("Hi. This is Tony, alias ttnny."))
			displayMenu(conn)
		case "2":
			conn.Write([]byte("This email feature is being implemented. Please try back soon."))
			displayMenu(conn)
		case "3":
			conn.Write([]byte("Goodbye!"))
			displayMenu(conn)

			err := conn.Close()
			if err != nil {
				log.Panic(err)
			}
		}
	}

	defer conn.Close()
}

func displayMenu(conn net.Conn) {
	io.WriteString(conn, fmt.Sprint("MAIN MENU\n", "---------\n", "1. About\n", "2. Contact\n", "3. Exit\n"))
}