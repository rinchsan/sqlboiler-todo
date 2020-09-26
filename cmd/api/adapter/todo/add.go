package todo

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rinchsan/sqlboiler-todo/pkg/entity"
	"github.com/rinchsan/sqlboiler-todo/pkg/presenter"
	"github.com/rinchsan/sqlboiler-todo/pkg/transaction"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type AddInput struct {
	Title           string    `json:"title"`
	Detail          string    `json:"detail"`
	AuthorUserID    uint64    `json:"author_user_id"`
	DueDate         time.Time `json:"due_date"`
	AssigneeUserIDs []uint64  `json:"assignee_user_ids"`
}

func (in AddInput) Todo() entity.Todo {
	return entity.Todo{
		Title:        in.Title,
		Detail:       in.Detail,
		DueDate:      in.DueDate,
		AuthorUserID: in.AuthorUserID,
	}
}

func (h handler) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var in AddInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		presenter.BadRequest(w, err.Error())
		return
	}

	f := func(tx *sql.Tx) error {
		todo := in.Todo()
		if err := todo.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}

		users, err := entity.Users(entity.UserWhere.ID.IN(in.AssigneeUserIDs)).All(ctx, tx)
		if err != nil {
			return err
		}

		if err := todo.AddUsers(ctx, tx, false, users...); err != nil {
			return err
		}

		return nil
	}

	if err := transaction.Run(h.db, f); err != nil {
		presenter.Error(w, err)
		return
	}

	presenter.Success(w)
}
