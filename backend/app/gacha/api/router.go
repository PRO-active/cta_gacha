package api

import (
	"github.com/go-chi/chi/v5"

	"github.com/pro-active/cta_gacha/app/gacha/api/user"
)

func (a *api) Router(r chi.Router) {
	r.Route("/api/v1", func(r chi.Router) {
		user.RegisterRouter(r, a.adminAccountUseCase)
	})
}
