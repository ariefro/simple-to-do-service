package service_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	todo "github.com/ariefro/simple-to-do-service/api/protogen/todos"
	"github.com/ariefro/simple-to-do-service/pkg/service"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const testTitle = "Dummy title"
const testDescription = "This is a test description"
const readQuery = "SELECT id, title, description, reminder FROM todo Where id = ?"

func TestCreateToDoSuccess(t *testing.T) {
	// mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// create a service instance
	srv := service.NewTodoServiceServer(db)

	// define input
	req := &todo.CreateToDoRequest{
		ToDo: &todo.ToDo{
			Title:       testTitle,
			Description: testDescription,
			Reminder:    timestamppb.New(time.Now()),
		},
	}

	// mock database behavior for success case
	mock.ExpectExec("INSERT INTO todo").WithArgs(req.ToDo.GetTitle(), req.ToDo.GetDescription(), req.ToDo.GetReminder().AsTime()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := srv.Create(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, int64(1), res.Id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateToDoEmptyTitle(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	srv := service.NewTodoServiceServer(db)

	req := &todo.CreateToDoRequest{
		ToDo: &todo.ToDo{
			Title:       "",
			Description: testDescription,
			Reminder:    timestamppb.New(time.Now()),
		},
	}

	res, err := srv.Create(context.Background(), req)

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
	assert.Contains(t, err.Error(), "title cannot be empty")
}

func TestCreateToDoDatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	srv := service.NewTodoServiceServer(db)

	req := &todo.CreateToDoRequest{
		ToDo: &todo.ToDo{
			Title:       testTitle,
			Description: testDescription,
			Reminder:    timestamppb.New(time.Now()),
		},
	}

	mock.ExpectExec("INSERT INTO todo").
		WithArgs(req.ToDo.GetTitle(), req.ToDo.GetDescription(), req.ToDo.GetReminder().AsTime()).
		WillReturnError(sql.ErrConnDone)

	res, err := srv.Create(context.Background(), req)

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Equal(t, codes.Internal, status.Code(err))
	assert.Contains(t, err.Error(), "failed to insert into todo")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestReadToDoSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	srv := service.NewTodoServiceServer(db)

	mock.ExpectQuery(readQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "reminder"}).
			AddRow(1, testTitle, testDescription, time.Now()))

	req := &todo.ReadToDoRequest{
		Id: 1,
	}

	res, err := srv.Read(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, int64(1), res.ToDo.Id)
	assert.Equal(t, testTitle, res.ToDo.Title)
}

func TestReadToDoNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	srv := service.NewTodoServiceServer(db)

	mock.ExpectQuery(readQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{}))

	req := &todo.ReadToDoRequest{
		Id: 1,
	}

	res, err := srv.Read(context.Background(), req)

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))
	assert.Contains(t, err.Error(), "todo not found")
}

func TestReadToDoQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	srv := service.NewTodoServiceServer(db)

	mock.ExpectQuery(readQuery).
		WithArgs(1).
		WillReturnError(sql.ErrConnDone)

	req := &todo.ReadToDoRequest{
		Id: 1,
	}

	res, err := srv.Read(context.Background(), req)

	assert.Nil(t, res)
	assert.Error(t, err)
	assert.Equal(t, codes.Internal, status.Code(err))
	assert.Contains(t, err.Error(), "failed to retrive todo")
}
