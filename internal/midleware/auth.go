package midleware

import (
	"context"
	"net/http"
	"rest_api/internal/auth"
	"strings"
)


type contextKey string

const userIDKey contextKey = "userID"

func Auth(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHandler:=r.Header.Get("Authorization")

		if authHandler == ""{
			http.Error(w,"authorization header is required",http.StatusUnauthorized)
			return 
		}
		parts := strings.SplitN(authHandler, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid authorization header", http.StatusUnauthorized)
			return
		}

tokenString := parts[1]

		userID,err:=auth.ValidateToken(tokenString)
		if err!=nil{
			http.Error(w,"invalid token",http.StatusUnauthorized)
			return 
		}

		ctx:=context.WithValue(
			r.Context(),
			userIDKey,
			userID,
		)

		r=r.WithContext(ctx)

		next.ServeHTTP(w,r)
	})
}

func GetUserID(ctx context.Context)(int, bool){
	userID,ok:=ctx.Value(userIDKey).(int)
	return userID,ok
}