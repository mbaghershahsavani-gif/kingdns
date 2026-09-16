package cluster

import "net/http"

func Protected(
	handler http.HandlerFunc,
) http.HandlerFunc {

	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		if !Authenticate(r) {

			http.Error(
				w,
				"unauthorized",
				http.StatusUnauthorized,
			)

			return
		}

		handler(
			w,
			r,
		)
	}
}
