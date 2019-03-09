package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "rpc-say/api"

	"google.golang.org/grpc"
)

func main() {
	backend := flag.String("b", "localhost:8023", "addr of the flite backend")
	output := flag.String("o", "output.wav", "wav file where to write")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Printf("usage: \n\t%s \"text to speak\"\n", os.Args[0])
		os.Exit(1)
	}
	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %s : %v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)

	text := &pb.Text{Text: flag.Arg(0)}
	res, err := client.Say(context.Background(), text)
	if err != nil {
		log.Fatalf("could not say %s: %v", text.Text, err)
	}
	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("could not write to %s: %v", *output, err)
	}
}
