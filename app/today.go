package app

import "github.com/kaizoa/go-arch-example/domain/today"

type TodayService struct {
	Info *today.InfoUseCase
}