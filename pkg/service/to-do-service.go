package service

import (
	"context"
	"database/sql"
	"time"

	todo "github.com/ariefro/simple-to-do-service/api/protogen/todos"
	"github.com/ariefro/simple-to-do-service/pkg/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// toDoServiceServer is implementation of ToDoServiceServer proto interface
type toDoServiceServer struct {
	todo.UnimplementedToDoServiceServer
	db *(sql.DB)
}

func NewTodoServiceServer(db *sql.DB) todo.ToDoServiceServer {
	return &toDoServiceServer{db: db}
}

func (s *toDoServiceServer) Create(ctx context.Context, req *todo.CreateToDoRequest) (*todo.CreateToDoResponse, error) {
	err := util.ValidateTitle(req.ToDo.GetTitle())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	reminder := req.ToDo.GetReminder().AsTime()

	res, err := s.db.ExecContext(ctx, "INSERT INTO todo(`title`, `description`,`reminder`) VALUES (?, ?, ?)", req.ToDo.GetTitle(), req.ToDo.GetDescription(), reminder)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to insert into todo: "+err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve id for created todo: "+err.Error())
	}

	return &todo.CreateToDoResponse{
		Id: id,
	}, nil
}

func (s *toDoServiceServer) Read(ctx context.Context, req *todo.ReadToDoRequest) (*todo.ReadToDoResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "todo id is required")
	}

	query := "SELECT id, title, description, reminder FROM todo Where id = ?"

	row := s.db.QueryRowContext(ctx, query, req.Id)

	var id int64
	var title, description string
	var reminder time.Time

	if err := row.Scan(&id, &title, &description, &reminder); err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "todo not found")
		}

		return nil, status.Error(codes.Internal, "failed to retrive todo: "+err.Error())
	}

	reminderProto := timestamppb.New(reminder)

	return &todo.ReadToDoResponse{
		ToDo: &todo.ToDo{
			Id:          id,
			Title:       title,
			Description: description,
			Reminder:    reminderProto,
		},
	}, nil
}
