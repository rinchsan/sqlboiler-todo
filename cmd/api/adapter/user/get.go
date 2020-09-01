package user

import (
	"net/http"

	"github.com/rinchsan/sqlboiler-todo/pkg/entity"
	"github.com/rinchsan/sqlboiler-todo/pkg/presenter"
)

func (h handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := entity.Users().All(ctx, h.db)
	if err != nil {
		presenter.Error(w, err)
		return
	}

	presenter.Encode(w, users)
}
