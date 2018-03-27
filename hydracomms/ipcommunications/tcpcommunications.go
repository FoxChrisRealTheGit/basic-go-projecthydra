package main

// import(
// 	"bufio"
// 	"flag"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// 	"os"
// 	"strings"
// 	"time"
// )

// func main(){
// 	op := flag.String("type", "", "Server (s) or client (c) ?")
// 	address := flag.String("addr", ":8000", "address? host:port ")

// 	switch strings.ToUpper(*op){
// 	case "S":
// 		runClient(*address)
// 	case "C":
// 		runClient(*address)
// 	}
// }

// func runClient(address string) error{
// 	conn, err := net.Dial("tcp", address)
// 	if err != nil{
// 		return err
// 	}
// 	defer conn.Close()

// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Println("what message would you like to send?")
// 	for scanner.Scan(){
// 		fmt.Println("Writing ", scanner.Text())
// 		conn.Write(append(scanner.Bytes(), '\r'))

// 		fmt.Println("What message would you like to send?")
// 		buffer := make([]byte, 1024)
// 		//conn.SetReadDeadLine(time.Now().Add(5 * time.Second))
// 		_, err := conn.Read(buffer)

// 		if err != nil && err != io.EOF{
// 			log.Fatal(err)
// 		} else if err == io.EOF{
// 			log.Println("connection is close")
// 			return nil
// 		}
// 		fmt.Println(string(buffer))
// 	}
// 	return scanner.Err()
// }
// func runServer(addrwess string) error{
// 	l, err := net.Listen("tcp", address)
// 	if err != nil{
// 		return err
// 	}
// 	log.Println("Listening...")
// 	defer l.Close()
// 	for{
// 		c, err := l.Accept()
// 		if err != nil {
// 			return err
// 		}
// 		go handleConnection(c)
// 	}
// }

// func handleConnection(c net.Conn){
// 	defer c.Close()
// 	reader := bufio.NewReader(c)
// 	writer := bufio.NewWriter(c)
// 	for{
// 		//buffer := make([]byte, 1024)
// 		c.SetDeadline(time.Now().Add(5 * time.Second))
// 		line, err := reader.ReadString('\r')
// 		//_, err := c.Read(buffer)
// 		if err != nil && err != io.EOF{
// 			log.Println(err)
// 			return
// 		} else if err == io.EOF{
// 			log.Println("Connection closed")
// 			return
// 		}
// 		fmt.Printf("Recieved %s from address %s \n", line[:len(line)-1], c.RemoteAddr())
// 		writer.WriteString("Message recieved...")
// 		writer.Flush()
// 	}
// }