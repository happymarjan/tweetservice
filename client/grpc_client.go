package main

import (
	"fmt"
	"io"
	"log"
	//"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "google.golang.org/pb"
)

const (
	address = "localhost:10000"
)

func getNews(client pb.TweetsServiceClient, emptyParam *pb.EmptyParam) {
	grpclog.Printf("GRPC client getting stream news")
	stream, err := client.GetTweet(context.Background(), emptyParam)
	if err != nil {
		fmt.Println("First error")
		grpclog.Fatalf("%v.GetTweet(_) = _, %v first error", client, err)
	}

	for {
		nws, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Second error")
			grpclog.Fatalf("%v.GetTweet(_) = _, %v second error", client, err)
		}
		grpclog.Println(nws)
		grpclog.Println("\n")
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTweetsServiceClient(conn)

	var emptyParam *pb.EmptyParam
	emptyParam = new(pb.EmptyParam)
	getNews(c, emptyParam)
}
