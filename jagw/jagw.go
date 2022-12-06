package jagw

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type JagwEndpoint struct {
	EndpointAddress string
	EndpointPort    uint16
}

func NewJagwConnection(f interface{}) (connection *grpc.ClientConn, err error) {
	endpoint := getServiceEndpoint(f.(JagwEndpoint))
	connection, err = grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return connection, err
}

func getServiceEndpoint(jagwEndpoint JagwEndpoint) string {
	return fmt.Sprintf("%s:%d", jagwEndpoint.EndpointAddress, jagwEndpoint.EndpointPort)
}
