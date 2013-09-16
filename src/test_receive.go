package main

//import (
//	"github.com/gmallard/stompngo"
//	"log"
//	"net"
//	"strings"
//)

//type Log struct {
//	ApplicationName string
//	LogType         string
//	EntryDate       string
//	Description     string
//}

////test := `<Log>
//// <ApplicationName> </ApplicationName>
//// <LogType> </LogType>
//// <EntryDate> </EntryDate>
//// <Description> </Description>
////</Log>`

////body := []byte(`
////<Log>
//// <ApplicationName>an</ApplicationName>
//// <LogType>lt</LogType>
//// <EntryDate>ed</EntryDate>
//// <Description>d</Description>
////</Log>`)
////	log.Printf("body %v", string(body))

////	s := Log{}
////	if err := xml.Unmarshal(body, &s); err != nil {
////		return s, err
////	}
////	return s, nil

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

//	u := stompngo.Uuid()
//	s := stompngo.Headers{"destination", "/topic/*", "id", u}
//	r, err := conn.Subscribe(s)
//	for message := range r {
//		if message.Error != nil {
//			log.Fatalln(message.Error)
//		}
//		log.Println(strings.Join(message.Message.Headers, ","))
//		log.Println(string(message.Message.Body))
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
