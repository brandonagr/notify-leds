package main

//import (
//	"fmt"
//	"github.com/gmallard/stompngo"
//	"log"
//	"net"
//)

//// Connect to a STOMP 1.1 broker, send some messages and disconnect.
//func main() {
//	log.Println("starts ...")

//	// Open a net connection
//	n, err := net.Dial("tcp", net.JoinHostPort("localhost", "61613"))
//	if err != nil {
//		log.Fatalln(err)
//	}
//	log.Println("dial complete ...")

//	ch := stompngo.Headers{"accept-version", "1.1", "host", "localhost"}
//	conn, err := stompngo.Connect(n, ch)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	log.Println("stomp connect complete ...", conn.Protocol())

//	s := stompngo.Headers{"destination", "/topic/test"} // send headers
//	m := " message: "
//	for i := 1; i <= 10; i++ {
//		t := m + fmt.Sprintf("%d", i)
//		err := conn.Send(s, t)
//		if err != nil {
//			log.Fatalln(err)
//		}
//		log.Println("send complete:", t)
//	}

//	// Disconnect from the Stomp server
//	err = conn.Disconnect(stompngo.Headers{})
//	if err != nil {
//		log.Fatalln(err) // Handle this ......
//	}
//	log.Println("stomp disconnect complete ...")

//	// Close the network connection
//	err = n.Close()
//	if err != nil {
//		log.Fatalln(err) // Handle this ......
//	}
//	log.Println("network close complete ...")

//	log.Println("ends ...")
//}
