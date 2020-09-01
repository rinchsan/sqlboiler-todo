package user

import (
	"encoding/json"
	"net/http"

	"github.com/rinchsan/sqlboiler-todo/pkg/entity"
	"github.com/rinchsan/sqlboiler-todo/pkg/presenter"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type AddInput struct {
	Username string `json:"username"`
}

func (in AddInput) User() entity.User {
	return entity.User{
		Username: in.Username,
	}
}

func (h handler) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var in AddInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		presenter.BadRequest(w, err.Error())
		return
	}

	user := in.User()
	if err := user.Insert(ctx, h.db, boil.Infer()); err != nil {
		presenter.Error(w, err)
		return
	}

	presenter.Success(w)
}
