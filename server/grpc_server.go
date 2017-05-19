package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "google.golang.org/pb" //put the two files in the parent directory to this address!
	"log"
	"net"
	//"google.golang.org/grpc/reflection"
	//"reflect"
)

const (
	grpc_port = ":10000"
)

type tweetServer struct {
	loadedStatusArr []*pb.Status
	loadedTweets    *pb.Statuslist
}

var loadedNews []string

func (s *tweetServer) GetTweet(emptyParam *pb.EmptyParam, stream pb.TweetsService_GetTweetServer) error {
	if err := stream.Send(s.loadedTweets); err != nil {
		fmt.Println("Error hapenned in server send: %v", err)
		return err
	}
	return nil
}

func (s *tweetServer) loadTweets() {
	loadedNews = dbObject.queryTable()
	fmt.Println("The length of array is", len(loadedNews))
	//fmt.Println("The type of array is", reflect.TypeOf(loadedNews))
	s.loadedStatusArr = make([]*pb.Status, len(loadedNews))

	for i, nws := range loadedNews {
		s.loadedStatusArr[i] = new(pb.Status)
		s.loadedStatusArr[i].Text = nws
	}

	//fmt.Println("GRPC Server")
	fmt.Println(s.loadedStatusArr)

	s.loadedTweets = new(pb.Statuslist)
	s.loadedTweets.StatusArray = make([]*pb.Status, len(loadedNews))
	//loadedObj := new(pb.Statuslist)

	for i, nws := range s.loadedStatusArr {
		s.loadedTweets.StatusArray[i] = new(pb.Status)
		s.loadedTweets.StatusArray[i].Text = nws.Text
	}
}

func newServer() *tweetServer {
	tweetServerObj := new(tweetServer)
	fmt.Println("GRPC Server initiated")
	tweetServerObj.loadTweets()
	return tweetServerObj
}

func grpcMain(dbObj *DB) {
	dbObject = dbObj
	lis, err := net.Listen("tcp", grpc_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTweetsServiceServer(s, newServer())
	// Register reflection service on gRPC server.
	//reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Println("Error %v", err)
		log.Fatalf("failed to serve: %v", err)
	}

}
