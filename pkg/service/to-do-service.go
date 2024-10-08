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

func (s *toDoServiceServer) ReadAll(ctx context.Context, req *todo.ReadAllToDoRequest) (*todo.ReadAllToDoResponse, error) {
	query := "SELECT id, title, description, reminder FROM todo"

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve todos: "+err.Error())
	}
	// defer closing rows to ensure it runs after data has been processed
	defer rows.Close()


	var todos []*todo.ToDo
	for rows.Next() {
		var(
			id int64
			title string
			description string
			reminder time.Time
		)

		// scan the values from the current row
		if err := rows.Scan(&id, &title,&description, &reminder); err != nil {
			return nil, status.Error(codes.Internal, "failed to scan todo item: "+err.Error())
		}

		// Convert reminder to protobuf timestamp
		reminderProto := timestamppb.New(reminder)

		todos = append(todos, &todo.ToDo{
			Id: id,
			Title: title,
			Description: description,
			Reminder: reminderProto,
		})
	}
	
	// Check if there was an error during row iteration
	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve todos: "+err.Error())
	}

	return &todo.ReadAllToDoResponse{
		ToDo: todos,
	}, nil
}

func (s *toDoServiceServer) Update(ctx context.Context, req *todo.UpdateToDoRequest) (*todo.UpdateToDoResponse, error) {
	if req.ToDo.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "todo id is required")
	}

	if err := util.ValidateTitle(req.ToDo.GetTitle()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	query := "UPDATE todo SET title = ?, description = ?, reminder = ? WHERE id = ?"

	reminder := req.ToDo.GetReminder().AsTime()
	res, err := s.db.ExecContext(ctx, query, req.ToDo.GetTitle(), req.ToDo.GetDescription(), reminder, req.ToDo.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update todo: " + err.Error())
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve affected rows: " + err.Error())
	}
	if rowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "todo not found")
	}

	return &todo.UpdateToDoResponse{
		Success: true,
	}, nil
}

func (s *toDoServiceServer) Delete(ctx context.Context, req *todo.DeleteRequest) (*todo.DeleteResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "todo id is required")
	}

	query := "DELETE FROM todo WHERE id = ?"

	res, err := s.db.ExecContext(ctx, query, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete todo: " + err.Error())
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve affected rows: " + err.Error())
	}
	if rowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "todo not found")
	}

	return &todo.DeleteResponse{
		Success: true,
	}, nil
}
