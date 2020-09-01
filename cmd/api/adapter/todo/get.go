package todo

import (
	"net/http"

	"github.com/rinchsan/sqlboiler-todo/pkg/entity"
	"github.com/rinchsan/sqlboiler-todo/pkg/presenter"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type GetAllResult struct {
	Todo          *entity.Todo     `json:"todo"`
	AuthorUser    *entity.User     `json:"author_user"`
	AssigneeUsers entity.UserSlice `json:"assignee_users"`
}

func (h handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	todos, err := entity.Todos(
		qm.Load(entity.TodoRels.AuthorUser),
		qm.Load(entity.TodoRels.Users),
	).All(ctx, h.db)
	if err != nil {
		presenter.Error(w, err)
		return
	}

	results := make([]GetAllResult, len(todos))
	for i, todo := range todos {
		results[i] = GetAllResult{
			Todo:          todo,
			AuthorUser:    todo.R.AuthorUser,
			AssigneeUsers: todo.R.Users,
		}
	}

	presenter.Encode(w, results)
}
