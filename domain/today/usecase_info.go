package today

import "time"

type InfoOutput struct {
	Year int
	Month time.Month
	Day int
}

type InfoUseCase struct {
	Now Now
}

func (u *InfoUseCase) Get() *InfoOutput {
	// 現在時刻から年月日オブジェクトへの変換を行う
	now := u.Now.Time()
	out := &InfoOutput{}
	out.Year, out.Month, out.Day = now.Date()
	return out
}