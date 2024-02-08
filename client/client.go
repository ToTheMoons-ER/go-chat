package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Connect to the server เปิดการเชื่อมต่อไปที่ server ที่รันบน localhost:8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		//ถ้ามีข้อผิดพลาดในการเชื่อมต่อ จะจบการทำงาน.
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	// Create a reader to read user input
	reader := bufio.NewReader(os.Stdin)

	//ในลูป for, โปรแกรมรอรับข้อมูลจากผู้ใช้ที่ป้อนผ่านคีย์บอร์ดด้วย reader.ReadString('\n')
	for {
		// Read user input
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')

		// Check if the user wants to quit
		//ถ้าผู้ใช้ป้อน "quit", โปรแกรมจะพิมพ์ "Quitting the program..." และจบลูป for.
		if strings.TrimSpace(message) == ":quit" {
			fmt.Println("Quitting the program...")
			break
		}

		// Send the message to the server
		conn.Write([]byte(message))

		// Print the number of bytes sent
		fmt.Printf("Sent %d bytes\n", len(message))

		// Receive and print the server's response
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		fmt.Printf("Server response: %s", buffer[:n])
	}
}