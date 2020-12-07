package graph

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type userIDKey struct{}

func (r *Resolver) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// get the authorization header
		h := req.Header.Get("Authorization")
		tok := strings.TrimPrefix(h, "Bearer ")
		if len(tok) == 0 {
			next.ServeHTTP(w, req)
			return
		}

		// inspect the header
		j, err := jwt.Parse(tok, func(token *jwt.Token) (interface{}, error) {
			return r.PublicKey, nil
		})
		if err != nil && err.Error() == "Token is expired" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if err != nil {
			fmt.Printf("Error parsing token: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// ensure the token is valid
		if !j.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// set the user id in the context
		claims := j.Claims.(jwt.MapClaims)
		ctx := context.WithValue(req.Context(), userIDKey{}, claims["sub"])
		*req = *req.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
func UserIDFromContext(ctx context.Context) string {
	v := ctx.Value(userIDKey{})
	if v == nil {
		return ""
	}
	return v.(string)
}
