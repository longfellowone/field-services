package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
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
			return token, errors.New("Invalid audience.")
		}

		iss := "https://dev-vqglrbz9.auth0.com/"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("Invalid issuer.")
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
	CredentialsOptional: true,
	//Debug:               true,
})

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			err := jwtToken.CheckJWT(w, r)
			if err != nil {
				return
			}

			token, err := jwtmiddleware.FromAuthHeader(r)
			if err != nil {
				return
			}

			if token == "" {
				fmt.Println("no token")
				next.ServeHTTP(w, r)
				return
			}

			hasScope := checkScope("read:messages", token)
			if !hasScope {
				responseJSON("Insufficient scope.", w, http.StatusForbidden)
				return
			}

			fmt.Println("got token: ", token)

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

type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

func checkScope(scope string, tokenString string) bool {
	token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, nil)

	claims, _ := token.Claims.(*CustomClaims)

	hasScope := false
	result := strings.Split(claims.Scope, " ")
	for i := range result {
		if result[i] == scope {
			hasScope = true
		}
	}

	return hasScope
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
