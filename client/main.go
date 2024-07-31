package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sembh1998/golang-grpc-case/testpb"
	"github.com/sembh1998/golang-grpc-case/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial("localhost:5061", grpc.WithTransportCredentials((insecure.NewCredentials())))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := testpb.NewTestServiceClient(cc)

	DoBidirectionalStreaming(c)

}

func DoUnary(c testpb.TestServiceClient) {
	req := &testpb.GetTestRequest{
		Id: "t1",
	}

	res, err := c.GetTest(context.Background(), req)
	if err != nil {
		log.Println("error:", err)
		return
	}

	log.Printf("response: %v", res)
}

func DoClientStreaming(c testpb.TestServiceClient) {
	commonTestId := "t1"
	questionList := []*testpb.Question{
		{
			Id:       utils.CutStringTo32(utils.NewUUIDV7()),
			Answer:   uuid.New().String(),
			Question: uuid.New().String(),
			TestId:   commonTestId,
		},
		{
			Id:       utils.CutStringTo32(utils.NewUUIDV7()),
			Answer:   uuid.New().String(),
			Question: uuid.New().String(),
			TestId:   commonTestId,
		},
		{
			Id:       utils.CutStringTo32(utils.NewUUIDV7()),
			Answer:   uuid.New().String(),
			Question: uuid.New().String(),
			TestId:   commonTestId,
		},
		{
			Id:       utils.CutStringTo32(utils.NewUUIDV7()),
			Answer:   uuid.New().String(),
			Question: uuid.New().String(),
			TestId:   commonTestId,
		},
		{
			Id:       utils.CutStringTo32(utils.NewUUIDV7()),
			Answer:   uuid.New().String(),
			Question: uuid.New().String(),
			TestId:   commonTestId,
		},
	}
	stream, err := c.SetQuestions(context.Background())
	if err != nil {
		log.Println("error:", err)
		return
	}
	for _, item := range questionList {
		log.Println("sending question:", item.Id)
		err = stream.Send(item)
		if err != nil {
			log.Println("error:", err)
		}
		time.Sleep(2 * time.Second)
	}

	msg, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("error:", err)
		return
	}
	log.Printf("server response: %v", msg)

}

func DoServerStreaming(c testpb.TestServiceClient) {
	req := &testpb.GetStudentsPerTestRequest{
		TestId: "t1",
	}

	stream, err := c.GetStudentsPerTest(context.Background(), req)
	if err != nil {
		log.Println("error:", err)
		return
	}

	for {
		student, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("error:", err)
		}
		log.Println("response from server:", student)
	}
	err = stream.CloseSend()
	if err != nil {
		log.Println("error:", err)
		return
	}
}

func DoBidirectionalStreaming(c testpb.TestServiceClient) {

	answer := testpb.TakeTestRequest{
		Answer: uuid.NewString(),
	}

	numberOfQuestions := 4

	waitChannel := make(chan struct{})

	stream, err := c.TakeTest(context.Background())
	if err != nil {
		log.Println("error:", err)
		return
	}

	go func() {
		for i := 0; i < numberOfQuestions; i++ {
			stream.Send(&answer)
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("error:", err)
				continue
			}
			log.Printf("response from server: %v·\n", res)
		}
		close(waitChannel)
	}()
	<-waitChannel
}
