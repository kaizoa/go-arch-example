package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kaizoa/go-arch-example/app"
)

type TodayHandler struct {
	Service *app.TodayService
}

func (h *TodayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	out := h.Service.Info.Get()

	err = json.NewEncoder(w).Encode(
		&struct {
			OK    bool       `json:"ok"`
			Year  int        `json:"year"`
			Month time.Month `json:"month"`
			Day   int        `json:"day"`
		}{
			OK:    true,
			Year:  out.Year,
			Month: out.Month,
			Day:   out.Day,
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"ok":false, "error":%s}`, err)))
	}
}
