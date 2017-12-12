package bootstrap

import (
	"github.com/kaizoa/go-arch-example/app"
	"github.com/kaizoa/go-arch-example/domain/today"
	"github.com/kaizoa/go-arch-example/middleware"
	"github.com/kaizoa/go-arch-example/web"
)

func Today(m *middleware.M) {

	// Todayサービスに実装をインジェクトする
	s := &app.TodayService{
		Info: &today.InfoUseCase{
			Now: m.Now,
		},
	}

	// today.info
	m.Router.Handle("/today.info", &web.TodayHandler{
		Service: s,
	})
}
