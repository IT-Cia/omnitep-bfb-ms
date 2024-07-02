package task

import (
	"context"
	"fmt"
	"gsm/protobuf/taskpb"
	"strings"
)

var idCounter int64 = 1

func getNewID() int64 {
	defer func() {
		idCounter++
	}()

	return idCounter
}

type Task struct {
	ID    int64
	Title string
}

var tasks = make([]*Task, 0)

func Create(ctx context.Context, req *taskpb.CreateRequest) (*taskpb.CreateResponse, error) {
	task := &Task{
		ID:    getNewID(),
		Title: req.GetTitle(),
	}

	tasks = append(tasks, task)

	res := taskpb.CreateResponse{
		Id:    task.ID,
		Title: task.Title,
	}

	return &res, nil
}

func Get(ctx context.Context, req *taskpb.GetRequest) (*taskpb.GetResponse, error) {
	var task Task
	for _, t := range tasks {
		if t.ID == req.GetId() {
			task = *t
		}
	}

	if task.ID == 0 {
		return nil, fmt.Errorf("task not found")
	}

	res := taskpb.GetResponse{
		Id:    task.ID,
		Title: task.Title,
	}

	return &res, nil
}

func Update(ctx context.Context, req *taskpb.UpdateRequest) error {
	found := false
	for _, t := range tasks {
		if t.ID == req.GetId() {
			t.Title = req.GetTitle()
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task not found")
	}

	return nil
}

func Delete(ctx context.Context, req *taskpb.DeleteRequest) error {
	found := false
	filtered := make([]*Task, 0, len(tasks))
	for _, t := range tasks {
		if t.ID != req.GetId() {
			filtered = append(filtered, t)
		} else {
			found = true
		}
	}

	tasks = filtered

	if !found {
		return fmt.Errorf("task not found")
	}

	return nil
}

func GetAll(ctx context.Context, req *taskpb.GetAllRequest) (*taskpb.GetAllResponse, error) {
	size := req.GetSize()
	if size < 10 {
		size = 10
	}

	page := req.GetPage()
	if page < 0 {
		size = 0
	}

	items := make([]*taskpb.Task, 0, req.GetSize())
	for index, t := range tasks {
		if int(size)*int(page) <= index &&
			index < int(size)*int(page+1) &&
			strings.Contains(t.Title, req.GetTerm()) {
			items = append(items, &taskpb.Task{
				Id:    t.ID,
				Title: t.Title,
			})
		}
	}

	res := taskpb.GetAllResponse{
		Items: items,
	}

	return &res, nil
}
