package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jalapeno-api-gateway/jagw-go/jagw"
)

func main() {
	requestEndpoint := jagw.JagwEndpoint{
		EndpointAddress: os.Args[1],
		EndpointPort:    9903,
	}
	requestConnection, err := jagw.NewJagwConnection(requestEndpoint)
	if err != nil {
		panic(err)
	}
	defer requestConnection.Close()

	client := jagw.NewRequestServiceClient(requestConnection)

	// Request all Nodes
	request := &jagw.TopologyRequest{
		Keys:       []string{},
		Properties: []string{},
	}

	response, err := client.GetLsNodes(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling GetLsNodes on request service: %s", err)
	}

	prettyPrint(response)
}

func prettyPrint(any interface{}) {
	s, _ := json.MarshalIndent(any, "", "  ")
	fmt.Printf("%s\n\n", string(s))
}
