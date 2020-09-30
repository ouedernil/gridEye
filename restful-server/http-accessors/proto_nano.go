package http_accessors

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/pair"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
	"fmt"
	"log"
	"time"
	"github.com/golang/protobuf/proto"
	"os"
	urllib "net/url"
	"strings"
)

func protoMaker(tMRC uint32, tV uint32) {
	var tagMonRemoteConfig uint32 = tMRC
	var tagValue uint32 = tV
	var updateConfig = MonMsgEnumMsgId(MonMsg_update_config)
	var remoteConfigMsg = &MonRemoteConfigOneParam{
		TagMonRemoteConfig: &tagMonRemoteConfig,
		TagValue:           &tagValue,
	}

	monMessage := &MonMsg{
		&updateConfig,
		remoteConfigMsg,
		nil,
	}
	data, err := proto.Marshal(monMessage)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	initNanoSocket(data)
}

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}


func newPairSocket(url string) (mangos.Socket, error) {
	socket, err := pair.NewSocket()
	if err != nil {
		return nil, err
	}
	socket.AddTransport(ipc.NewTransport())
	socket.AddTransport(tcp.NewTransport())

	// check URL
	u, err := urllib.Parse(url)
	if err != nil {
		die("Invalid URL: %s", err.Error())
	}

	// Check if cleanup needed
	if strings.ToLower(u.Scheme) == "ipc" {
		_, err = os.Stat(u.Path)
		if !os.IsNotExist(err) {
			log.Printf("Warning, socket still exists!")
			err := os.Remove(u.Path)
			if err != nil {
				die ("Unable to unlink %v", u.Path)
			}
		}
	}

	if err = socket.Listen(url); err != nil {
		die("can't listen on socket: %s", err.Error())
	}
	return socket, nil
}

func send(socket mangos.Socket, message []byte) error {
	err := socket.Send(message)
	if err != nil{
		fmt.Println(err)
	}
	return err
}

func runServer(url string, msg []byte) {
	socket, err := newPairSocket(url)
	if err != nil {
		log.Fatalf("Cannot listen on %s: %s\n", url, err.Error())
	}

	defer socket.Close()

	time.Sleep(2 * time.Second)
	err = send(socket, msg)
	if err != nil {
		log.Fatalf("Cannot send message %s: %s\n", err.Error())
	}
}

func initNanoSocket(msg []byte) {
	var url string= "ipc:///tmp/monitoring.ipc"
	runServer(url, msg)
}