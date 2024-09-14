package service

import (
	"context"
	"database/sql"

	todo "github.com/ariefro/simple-to-do-service/api/protogen/todos"
	"github.com/ariefro/simple-to-do-service/pkg/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
