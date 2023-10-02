package main

import (
	"context"
	"flag"
	"io"
	"log"
	pb "stream"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50055", "the address to connect to")
)

var testGetWeigth = true
var testGeState = false
var testStreaming = false

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewApiCallerScaleClient(conn)

	if testGetWeigth {
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()
		
		r, err := c.GetInstantWeight(ctx, &pb.Empty{})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("GetInstantWeight: %s  %s", r.GetMessage(), r.GetError())

	}
	

	if testGeState {
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		r2, err2 := c.GetState(ctx, &pb.Empty{})
		if err2 != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("GetState: Error %s, Message: %s, %s, %s", r2.GetError(), r2.GetMessage(), r2.GetSubtype(), r2.GetType())

	}

	
	if testStreaming {
		stream, err := c.ScalesMessageOutChannel(context.Background())
		if err != nil {
			log.Fatalf("openn stream error %v", err)
		}

		//var max int32
		ctx := stream.Context()
		done := make(chan bool)


		// first goroutine sends msg to stream and closes it after 10 iterations
		go func() {
			for i := 1; i <= 10; i++ {
				
				req := pb.RequestScale{ Message: "weigth", Type: "", Subtype: ""}
				if err := stream.Send(&req); err != nil {
					log.Fatalf("can not send %v", err)
				}
				log.Printf("%d sent", req.Message)
				time.Sleep(time.Millisecond * 200)
			}
			if err := stream.CloseSend(); err != nil {
				log.Println(err)
			}
		}()

		// second goroutine receives data from stream	
		// if stream is finished it closes done channel
		go func() {
			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					close(done)
					return
				}
				if err != nil {
					log.Fatalf("can not receive %v", err)
				}
				log.Printf("%d received", resp.Message)
			}
		}()


		// third goroutine closes done channel
		// if context is done
		go func() {
			<-ctx.Done()
			if err := ctx.Err(); err != nil {
				log.Println(err)
			}
			close(done)
		}()

		<-done
		log.Printf("finished")
	}
}
