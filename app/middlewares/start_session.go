package middlewares

import (
	"goblog/pkg/session"
	"net/http"
)

func StartSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//启动会话StartSession
		session.StartSession(w, r)
		// 2. . 继续处理接下去的请求
		next.ServeHTTP(w, r)

	})
}
