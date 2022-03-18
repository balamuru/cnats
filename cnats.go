package main

import "C"
import (
	"fmt"
	"github.com/nats-io/nats.go"
)


//export Pub
func Pub(url, subj, msg *C.char) {
			//note: had to transform func (u *UserTweet) GetLastTweet(user string) (string, error) 

			// Connect to NATS cluster on un-encrypted port
			natsConnection, _ := nats.Connect(C.GoString(url))
			defer natsConnection.Close()
			// Publish message on subject

			err := natsConnection.Publish(C.GoString(subj), []byte(C.GoString(msg)))
			if err != nil {
				fmt.Println("publish error: %v", err)
			}
			fmt.Printf("Published msg: [%s]\n", C.GoString(msg))
}

func main(){}