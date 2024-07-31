package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/sembh1998/golang-grpc-case/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (repo *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.ID, student.Name, student.Age)
	return err
}

func (repo *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var student models.Student

	for rows.Next() {
		err := rows.Scan(&student.ID, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		} else {
			return &student, nil
		}
	}
	return &student, nil
}

func (repo *PostgresRepository) SetTest(ctx context.Context, test *models.Test) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO tests (id, name) VALUES ($1, $2)", test.ID, test.Name)
	return err
}

func (repo *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name FROM tests WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var test models.Test

	for rows.Next() {
		err := rows.Scan(&test.ID, &test.Name)
		if err != nil {
			return nil, err
		} else {
			return &test, nil
		}
	}
	return &test, nil
}

func (repo *PostgresRepository) SetQuestion(ctx context.Context, question *models.Question) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO questions (id, question, answer, test_id) VALUES ($1, $2, $3, $4)", question.ID, question.Question, question.Answer, question.TestID)
	return err
}

func (repo *PostgresRepository) GetQuestion(ctx context.Context, id string) (*models.Question, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, question, answer, test_id name FROM questions WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var question models.Question

	for rows.Next() {
		err := rows.Scan(&question.ID, &question.Question, &question.Answer, &question.TestID)
		if err != nil {
			return nil, err
		} else {
			return &question, nil
		}
	}
	return &question, nil
}

func (repo *PostgresRepository) SetEnrollment(ctx context.Context, enrollement *models.Enrollment) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO enrollments (test_id, student_id) VALUES ($1, $2)", enrollement.TestID, enrollement.StudentID)
	return err
}
func (repo *PostgresRepository) GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id IN (SELECT student_id FROM enrollments WHERE test_id = $1)", testId)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	students := []*models.Student{}

	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Age)
		if err == nil {
			students = append(students, &student)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return students, nil
}

func (repo *PostgresRepository) GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, question FROM questions WHERE test_id = $1", testId)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	questions := []*models.Question{}

	for rows.Next() {
		var question models.Question
		err := rows.Scan(&question.ID, &question.Question)
		if err == nil {
			questions = append(questions, &question)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return questions, nil
}
