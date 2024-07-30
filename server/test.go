package server

import (
	"context"
	"io"
	"time"

	"github.com/sembh1998/golang-grpc-case/models"
	"github.com/sembh1998/golang-grpc-case/repository"
	"github.com/sembh1998/golang-grpc-case/studentpb"
	"github.com/sembh1998/golang-grpc-case/testpb"
)

type TestServer struct {
	repo repository.Repository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Repository) *TestServer {
	return &TestServer{
		repo: repo,
	}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &testpb.Test{
		Id:   test.ID,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
	test := &models.Test{
		ID:   req.Id,
		Name: req.Name,
	}
	err := s.repo.SetTest(ctx, test)
	if err != nil {
		return nil, err
	}
	return &testpb.SetTestResponse{
		Id: test.ID,
	}, nil
}

func (s *TestServer) SetQuestions(stream testpb.TestService_SetQuestionsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}
		if err != nil {
			return err
		}
		question := &models.Question{
			ID:       msg.Id,
			Question: msg.Question,
			Answer:   msg.Answer,
			TestID:   msg.TestId,
		}
		err = s.repo.SetQuestion(stream.Context(), question)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

func (s *TestServer) EnrollStudents(stream testpb.TestService_EnrollStudentsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.EnrollmentResponse{
				Ok: true,
			})
		}
		if err != nil {
			return err
		}
		enrollement := &models.Enrollment{
			StudentID: msg.StudentId,
			TestID:    msg.TestId,
		}
		err = s.repo.SetEnrollment(stream.Context(), enrollement)
		if err != nil {
			return stream.SendAndClose(&testpb.EnrollmentResponse{
				Ok: false,
			})
		}
	}
}

func (s *TestServer) GetStudentsPerTest(req *testpb.GetStudentsPerTestRequest, stream testpb.TestService_GetStudentsPerTestServer) error {
	students, err := s.repo.GetStudentsPerTest(stream.Context(), req.TestId)
	if err != nil {
		return err
	}
	for _, student := range students {
		to_send := &studentpb.Student{
			Id:   student.ID,
			Name: student.Name,
			Age:  student.Age,
		}
		err = stream.Send(to_send)
		time.Sleep(2 * time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
