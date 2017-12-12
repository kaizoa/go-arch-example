package main

import (
	"github.com/kaizoa/go-arch-example/bootstrap"
	"github.com/kaizoa/go-arch-example/middleware"
	"github.com/kaizoa/go-arch-example/middleware/http"
	"github.com/kaizoa/go-arch-example/middleware/os"
)

func main() {
	var err error

	rt := http.NewRouter()

	// ミドルウェアの実装をインジェクトする
	m := &middleware.M{
		Now:    &os.Now{},
		Router: rt,
	}

	// Todayサービスの初期化
	bootstrap.Today(m)

	// サーバーを起動
	err = http.ListenAndServe(":8080", rt)
	if err != nil {
		panic(err)
	}
}
