package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close() //ปิดเมื่อออกจากฟังก์ชั่น

	//สร้างตัวเเปล เพื่อที่จะเตรียมเอาไว้รองรับข้อมูลที่วิ่งมาหาเรา
	buffer := make([]byte, 1024)
	for {
		// Read data from the client
		n, err := conn.Read(buffer) // Read() blocks until it reads some data from the network and n is the number of bytes read
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		// Print the number of bytes read
		fmt.Printf("Received %d bytes\n", n)

		// Print received data
		fmt.Printf("Received message: %s", buffer[:n]) // :n is a slice operator that returns a slice of the first n bytes of the buffer

		fmt.Printf("Received message as bytes: %v\n", buffer[:n])
		// Send a response back to the client
		response := "Message received successfully\n"
		conn.Write([]byte(response))
	}
}
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	//ปิด listener
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		//ยอมรับการเชื่อมต่อ
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue //ให้วนลูปไปเรื่อยๆ เมื่อมี error เกิดขึ้น หรือ ให้ดำเนินการต่อไปเรื่อยๆ
		}

		fmt.Println("New connection established") //ติดต่อเรียบร้อย

		//ส่งต่อไปที่ฟังก์ชั่นข้างบน โดยส่งตัวเเปล conn ขึ้นไป
		go handleConnection(conn)
	}
}