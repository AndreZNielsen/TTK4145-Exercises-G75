package main

import (
	"fmt"
	"net"
	"time"
	"os"
	"bufio"
	"strings"
)

func receiverUDP(port int, done chan string) {

	old_addr := fmt.Sprintf(":%d", port)
	addr, err := net.ResolveUDPAddr("udp", old_addr)
	if err != nil {
		fmt.Println("Error resolving UDP adress", err)
	}
	socket, err := net.ListenUDP("udp", addr)

	if err != nil {
		fmt.Println("Error listing on UDP port:", err)
	}
	defer socket.Close()

	buffer := make([]byte, 1024)

	for {
		n, arr, _ := socket.ReadFromUDP(buffer)
		fmt.Printf("%s\n", string(buffer[0:n-1]))
		done<-arr.String()
	}

}

func transmitterUDP(ip string, port int, message string) {
	address := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(ip),

	}
	socket, _ := net.DialUDP("udp", nil, &address)

	defer socket.Close()
	for {
		socket.Write([]byte(message))
		time.Sleep(2*time.Second)
	}
}

func conektTCP(ip string, port string) {

	input := bufio.NewReader(os.Stdin) // lar bruker skrive meling
	address := net.JoinHostPort(ip, port)
	address2 := net.TCPAddr{ // lager addresse
		Port: 20010,
		IP:   net.ParseIP(ip),
	}
	net.ListenTCP("tcp",&address2)
	tcp, _ := net.Dial("tcp", address)

	defer tcp.Close()
	//tcp.Write([]byte("Connect to:  0.100.23.204:20010\000"))
	buffer := make([]byte, 1024)
	n, _ := tcp.Read(buffer)
	fmt.Printf("serveren sa: %s\n", string(buffer[:n]))
	for {
		fmt.Print("hva vil du sende(skriv stopp for å avslute)")
		meling, _ := input.ReadString('\n')
		meling = strings.TrimSpace(meling)
		if meling == "stopp" {//stopper programet 
			fmt.Println("stopper")
			break
		}

		meling = meling + "\000"
		tcp.Write([]byte(meling)) //sender melingen

		n, _ := tcp.Read(buffer) //lesser fra serverne 
		fmt.Printf("serveren sa: %s\n", string(buffer[0:n]))
		
	}

}
var ip_server string = "10.100.23.204"
var port_udp int = 20010

// var ip string = "10.100.23.20"
var ip string = "0.0.0.0"

var port int = 30000

func main() {
	//done := make(chan string)
	//go receiverUDP(port, done)
	// go transmitterUDP(ip_server, port_udp, "God dag!")
	// go receiverUDP(port_udp, done)

	conektTCP(ip_server,"34933")
	// select {
	// case msg := <-done:
	// 	fmt.Println(msg)
	// }
}