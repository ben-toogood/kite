package resolvers

import (
	"crypto/rand"
	"crypto/rsa"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	// generate JWT private / public keys
	reader := rand.Reader
	bitSize := 2048
	key, err := rsa.GenerateKey(reader, bitSize)
	assert.NoErrorf(t, err, "Error generating public key")

	// initialize the resolver and a test handler to use
	r := Resolver{PublicKey: &key.PublicKey}
	h := r.AuthMiddleware(new(testHandler))

	t.Run("NoAuthHeader", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "http://localhost:3000/foo", nil)
		h.ServeHTTP(w, r)

		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Empty(t, UserIDFromContext(r.Context()))
	})

	t.Run("ValidAuthHeader", func(t *testing.T) {
		// generate a valid token
		uid := ksuid.New().String()
		at := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.StandardClaims{
			Issuer:    "kite",
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   uid,
		})
		tok, err := at.SignedString(key)
		assert.NoError(t, err)

		// call the handler
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "http://localhost:3000/foo", nil)
		r.Header.Set("Authorization", "Bearer "+tok)
		h.ServeHTTP(w, r)

		// check the result
		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, uid, UserIDFromContext(r.Context()))
	})

	t.Run("InvalidAuthHeader", func(t *testing.T) {
		// generate an invalid token
		uid := ksuid.New().String()
		at := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.StandardClaims{
			Issuer:    "kite",
			ExpiresAt: time.Now().Add(time.Minute * -5).Unix(),
			IssuedAt:  time.Now().Add(time.Minute * -10).Unix(),
			Subject:   uid,
		})
		tok, err := at.SignedString(key)
		assert.NoError(t, err)

		// call the handler
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "http://localhost:3000/foo", nil)
		r.Header.Set("Authorization", "Bearer "+tok)
		h.ServeHTTP(w, r)

		// check the result
		resp := w.Result()
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		assert.Empty(t, UserIDFromContext(r.Context()))
	})
}

type testHandler struct{}

func (h *testHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}
