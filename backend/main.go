package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"

	"golang.org/x/net/context"
	flite "rpc-cgo"
	pb "rpc-say/api"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	port := flag.Int("p", 8023, "port to listen to")
	flag.Parse()

	logrus.Infof("listening to port %d", *port)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("could not listen to the tcp on port %d: %v", *port, err)
	}

	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, server{})
	err = s.Serve(listener)
	if err != nil {
		logrus.Fatalf("could not serve: %v", err)
	}
}

type server struct{}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("could not create tmp file %v", err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("could not close %s: %v", f.Name(), err)
	}

	if err := flite.TextToSpeech(f.Name(), text.Text); err != nil {
		return nil, fmt.Errorf("flite failed: %v", err)
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read tmp file: %v", err)
	}
	return &pb.Speech{Audio: data}, nil
}
