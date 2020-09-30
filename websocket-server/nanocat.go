package main

//go:generate protoc -I=proto --go_out=import_path=main:. proto/common.proto  proto/mon_msg.proto

import (
	"flag"
	"fmt"
	_ "encoding/json"
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/sub"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
	"os"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

func die(format string, v ...interface{}) {
	log.Print("["+time.Now().String()+"] - Log message : ")
	log.Println(os.Stderr, fmt.Sprintf(format, v...))
}

func main() {
	var url = flag.String("url", "tcp://localhost:5555", "URL for socket")
	var data []byte
	sock, err := initSocket(*url)
	flag.Parse()
	log.Print("["+time.Now().String()+"] - Log message : ")
	log.Printf("Nanocat listening to %s\n", *url)
	go func(){
		for {
			data, err = sock.Recv()
			if err != nil {
				die("Cannot recv: %s", err.Error())
			}
			//Reference to MonitoringEvt interface
			msg := &MonitoringEvt{}
			//De-serialize the recieved message into proto interface
			err = proto.Unmarshal(data, msg)
			if err != nil {
				log.Print("["+time.Now().String()+"] - Log message : ")
				log.Printf("Error< %v\n", err)
			} else {
				log.Print("["+time.Now().String()+"] - Log message : ")
				log.Printf("RECEIVED %v\n", string(msg.String()))
				publishMsg(msg)
			}
		}
	}()
	//Initialise wamp router
	//This function is contained in "wamp-server.go"
	wampInit()
}

// initSocket creates a new socket subscriber object.
// It connects using the url passed in the parameter
// It returns an object socket and the error parameter.
func initSocket(url string) (mangos.Socket, error){
	var sock mangos.Socket
	var err error
	// Create a new socket
	if sock, err = sub.NewSocket(); err != nil {
		die("can't get new sub socket: %s", err.Error())
	}
	// Protocols specifications
	sock.AddTransport(ipc.NewTransport())
	sock.AddTransport(tcp.NewTransport())
	if err = sock.Dial(url); err != nil {
		die("can't dial on sub socket: %s", err.Error())
	}
	// Empty byte array effectively subscribes to everything
	err = sock.SetOption(mangos.OptionSubscribe, []byte(""))
	if err != nil {
		die("cannot subscribe: %s", err.Error())
	}
	return sock, err
}
