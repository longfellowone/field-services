package auth

import (
	"fmt"
	"net/http"
)

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			fmt.Println(r.Header.Get("authorization"))

			// Allow unauthenticated users in
			//if err != nil || c == nil {
			//	next.ServeHTTP(w, r)
			//	return
			//}

			//userId, err := validateAndGetUserID(c)
			//if err != nil {
			//	http.Error(w, "Invalid cookie", http.StatusForbidden)
			//	return
			//}
			//
			//// get the user from the database
			//user := getUserByID(db, userId)
			//
			//// put it in context
			//ctx := context.WithValue(r.Context(), userCtxKey, user)
			//
			//// and call the next with our new context
			//r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
