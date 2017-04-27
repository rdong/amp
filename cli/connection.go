package cli

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"crypto/tls"
)

// NewClientConn is a helper function that wraps the steps involved in setting up a grpc client connection to the API.
func NewClientConn(addr string, token string, secure bool) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second),
		grpc.WithPerRPCCredentials(&LoginCredentials{Token: token}),
	}
	if secure {
		//sn := ""
		//if *serverHostOverride != "" {
		//	sn = *serverHostOverride
		//}
		var creds credentials.TransportCredentials
		//if *caFile != "" {
		//	var err error
		//	creds, err = credentials.NewClientTLSFromFile(*caFile, sn)
		//	if err != nil {
		//		grpclog.Fatalf("Failed to create TLS credentials %v", err)
		//	}
		//} else {
		tlsConfig := &tls.Config{
		}
		creds = credentials.NewTLS(tlsConfig)
		//}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	return grpc.Dial(addr, opts...)
}
