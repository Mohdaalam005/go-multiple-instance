package service

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber"
	"log"
	"net"
	"os"
	"strconv"
)

func Method1() {
	port := flag.Int("port", 8080, "")
	flag.Parse()

	log.Printf("Listening on port %v", *port)
	conn, _ := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	conn.Accept()
}

func Method2() {

	// Define routes and handlers
	route := fiber.New()

	// Get the port number from the command line
	port, _ := strconv.Atoi(os.Args[1])

	route.Listen(":" + strconv.Itoa(port))
	fmt.Println("Server started on port: ", port)

}
