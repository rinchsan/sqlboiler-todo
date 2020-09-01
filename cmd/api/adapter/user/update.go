package user

import (
	"encoding/json"
	"net/http"

	"github.com/rinchsan/sqlboiler-todo/pkg/entity"
	"github.com/rinchsan/sqlboiler-todo/pkg/presenter"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UpdateInput struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func (h handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var in UpdateInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		presenter.BadRequest(w, err.Error())
		return
	}

	user, err := entity.FindUser(ctx, h.db, in.ID)
	if err != nil {
		presenter.Error(w, err)
		return
	}

	user.Username = in.Username
	if _, err := user.Update(ctx, h.db, boil.Infer()); err != nil {
		presenter.Error(w, err)
		return
	}

	presenter.Success(w)
}
