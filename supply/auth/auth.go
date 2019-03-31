package auth

import (
	"context"
	"encoding/json"
	"errors"
	"field/supply"
	"fmt"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strings"

	"net/http"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

var jwtToken = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		aud := "http://192.168.0.104:8080/graphql"
		checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
		if !checkAud {
			return token, errors.New("Invalid audience")
		}

		iss := "https://dev-vqglrbz9.auth0.com/"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("Invalid issuer")
		}

		cert, err := getPemCert(token)
		if err != nil {
			log.Fatal(err.Error())
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	},
	SigningMethod: jwt.SigningMethodRS256,
	ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
		responseJSON(err, w, http.StatusUnauthorized)
	},
})

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := jwtToken.CheckJWT(w, r)
			if err != nil {
				return
			}

			authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
			token := authHeaderParts[1]

			user := populateUserDetails(token)

			if !user.IsForeman && !user.IsPurchaser {
				responseJSON("Insufficient scope", w, http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *supply.User {
	raw, _ := ctx.Value(userCtxKey).(*supply.User)
	return raw
}

type CustomClaims struct {
	UserID string `json:"sub"`
	Email  string `json:"https://localhost/email"`
	Scope  string `json:"scope"`
	jwt.StandardClaims
}

func populateUserDetails(tokenString string) *supply.User {
	token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, nil)
	claims, _ := token.Claims.(*CustomClaims)
	scopes := strings.Split(claims.Scope, " ")

	user := &supply.User{
		ID:          claims.UserID,
		Email:       claims.Email,
		IsForeman:   checkScope("create:orders", scopes),
		IsPurchaser: checkScope("manage:orders", scopes),
	}

	fmt.Println(user)
	return user
}

func checkScope(scope string, scopes []string) bool {
	ok := false
	for i := range scopes {
		if scopes[i] == scope {
			ok = true
		}
	}
	return ok
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://dev-vqglrbz9.auth0.com/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}

type Response struct {
	Message string `json:"message"`
}

func responseJSON(message string, w http.ResponseWriter, statusCode int) {
	response := Response{message}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
