package probez

import (
	"fmt"
	"net/http"
)

type Probe func() (bool, error)

func HandleHTTP(fn ...Probe) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, probe := range fn {
			if ok, err := probe(); !ok || err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, "not ok")
				return
			}
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	}
}
