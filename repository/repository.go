package repository

import (
	"context"

	"github.com/sembh1998/golang-grpc-case/models"
)

type Repository interface {
	SetStudent(ctx context.Context, student *models.Student) error
	GetStudent(ctx context.Context, id string) (*models.Student, error)

	SetTest(ctx context.Context, test *models.Test) error
	GetTest(ctx context.Context, id string) (*models.Test, error)

	SetQuestion(ctx context.Context, question *models.Question) error
	GetQuestion(ctx context.Context, id string) (*models.Question, error)

	SetEnrollment(ctx context.Context, enrollement *models.Enrollment) error
	GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error)

	GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error)
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}

func SetTest(ctx context.Context, test *models.Test) error {
	return implementation.SetTest(ctx, test)
}

func GetTest(ctx context.Context, id string) (*models.Test, error) {
	return implementation.GetTest(ctx, id)
}

func SetQuestion(ctx context.Context, question *models.Question) error {
	return implementation.SetQuestion(ctx, question)
}

func GetQuestion(ctx context.Context, id string) (*models.Question, error) {
	return implementation.GetQuestion(ctx, id)
}

func SetEnrollment(ctx context.Context, enrollement *models.Enrollment) error {
	return implementation.SetEnrollment(ctx, enrollement)
}
func GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	return implementation.GetStudentsPerTest(ctx, testId)
}

func GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	return implementation.GetQuestionsPerTest(ctx, testId)
}
