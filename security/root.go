package security

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"math/rand"
	// "net/http"
	"os"
	"strconv"
	"strings"
	"time"

	// "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/argon2"
)

//TODO MAKE UTILITY OF THIS
//TODO configurable token timeout
//TODO configurable hash secret

var (
	memory                 uint32 = 64 * 1024
	iterations             uint32 = 3
	parallelism            uint8  = 2
	saltLength             uint32 = 16
	keyLength              uint32 = 32
	errInvalidHash                = errors.New("the encoded hash is not in the correct format")
	errIncompatibleVersion        = errors.New("incompatible version of argon2")
	defaulttokentimeout           = time.Now().Add(1 * time.Minute)
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// //TokenCheck Middleware to secure endpoints
// func TokenCheck(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("Reached Handler For Path ", r.URL.Path)
// 		// r.Context()
// 		// ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
// 		// defer cancel()
// 		// r = r.WithContext(ctx)

// 		//SET ANY CONTEXT SPECIFIC VALUES HERE

// 		//GET PATH from request and SECURE EVERYTHING but /auth //SIMPLE USECASE
// 		/*
// 			path := strings.Split(html.EscapeString(r.URL.Path), "/")
// 			lastref := path[len(path)-1]
// 			if lastref == "auth" { //FORWARD WITHOUT TOKEN CHECK
// 				next.ServeHTTP(w, r)
// 				return
// 			}
// 		*/
// 		//IDEALLY DO ONLY COOKIE //EASIER TO MANAGE TOKEN RENEWAL
// 		// Get token from the Authorization header
// 		// format: Authorization: Bearer
// 		var token string
// 		tokens, ok := r.Header["Authorization"]
// 		if ok && len(tokens) >= 1 {
// 			token = tokens[0]
// 			token = strings.TrimPrefix(token, "Bearer ")
// 		} else {
// 			//CHECK IN COOKIE
// 			cookie, err := r.Cookie("token")
// 			if err != nil {
// 				if err == http.ErrNoCookie {
// 					// If the cookie is not set, return an unauthorized status
// 					w.WriteHeader(http.StatusUnauthorized)
// 					return
// 				}
// 				// For any other type of error, return a bad request status
// 				w.WriteHeader(http.StatusBadRequest)
// 				return
// 			}

// 			token = cookie.Value
// 		}

// 		valid, err := ValidateToken(&token) //VALIDATES AND RENEWS
// 		if err != nil {
// 			http.Error(w, "Authentication Failure", http.StatusUnauthorized)
// 			return
// 		}

// 		if !valid {
// 			http.Error(w, "Authentication Failure", http.StatusUnauthorized)
// 			return
// 		}
// 		http.SetCookie(w, &http.Cookie{
// 			Name:    "token",
// 			Value:   token,
// 			Expires: tokenexpirationTime(),
// 		})
// 		r.Header.Set("token", token)
// 		next.ServeHTTP(w, r)
// 	})
// }

//Authentication
/*
func doauth(w http.ResponseWriter, r *http.Request) {
	// authheaders := r.Header.Get("Authorization")
	ctx := r.Context()
	user, pass, ok := r.BasicAuth()
	if ok {
		tx := BeginTx()
		userdata, err := getuserbyusername(&ctx, user, tx)
		tx.CommitTX()

		unAuthorizedResponse(w, err)

		match, err := comparePasswordAndHash(pass, userdata.Password)
		unAuthorizedResponse(w, err)
		if !match {
			http.Error(w, "Authentication Failure", http.StatusUnauthorized)
		}

		//GENERATE TOKEN
		token, err := generateToken(user)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
		}

		//SEND JWT TOKEN
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   token,
			Expires: tokenexpirationTime(),
		})
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "token: %v", token)
	}
}

func unAuthorizedResponse(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "Authentication Failure", http.StatusUnauthorized)
		return
	}
}
*/

//PASSWORD GENERATION
func generateRandomPassword() string {
	rand.Seed(time.Now().UnixNano())
	randpwd := RandomString(10)
	log.Println("Generated Password:", randpwd)
	hash, err := GenerateFromPassword(randpwd)
	if err != nil {
		return randpwd //DUMMY
	}
	return hash
}

func decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errIncompatibleVersion
	}
	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

//RandomString Generates a random string of A-Z chars with len = l
func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

// Claims jwt claims container
// type Claims struct {
// 	Username string `json:"username"`
// 	jwt.StandardClaims
// }

// GenerateToken generates a JWT token with a configured token secret
// func GenerateToken(username string) (string, error) {
// 	claims := Claims{
// 		Username: username,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: tokenexpirationTime().Unix(),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte(tokensecret()))
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }

// ValidateToken validates an input token and tries to renew a token if close to expiring
// func ValidateToken(token *string) (bool, error) {
// 	claims := &Claims{}
// 	tokenparsed, err := jwt.ParseWithClaims(*token, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(tokensecret()), nil
// 	})
// 	if err != nil {
// 		return false, err
// 	}
// 	if !tokenparsed.Valid {
// 		return false, errors.New("Unauthorized")
// 	}

// 	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 30*time.Second { //IF TOKEN EXPIRES IN LESS THAN 30 SECONDS //RENEW IT for ANOTHER N Minutes
// 		claims.ExpiresAt = tokenexpirationTime().Unix()
// 		tokenobj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 		newtoken, err := tokenobj.SignedString([]byte(tokensecret()))
// 		if err != nil {
// 			//FAILED TO RENEW TOKEN
// 			log.Println("Failed to renew token", err)
// 		} else {
// 			*token = newtoken
// 		}
// 	}

// 	return true, nil
// }

func tokenexpirationTime() time.Time {
	tokenexpiry := os.Getenv("ENC_TOKEN_EXPIRY") //IN MILLISECONDS
	if len(strings.Trim(tokenexpiry, "")) > 0 {
		timeout, err := strconv.Atoi(tokenexpiry)
		if err != nil {
			log.Println(err)
			return defaulttokentimeout //DEFAULTING
		}
		return time.Now().Add(time.Duration(timeout) * time.Millisecond)
	}
	return defaulttokentimeout //DEFAULTING
}

func tokensecret() string {
	tokensecret := os.Getenv("ENC_SECRET_KEY")
	if len(strings.Trim(tokensecret, "")) == 0 {
		return "dummy" //PREFERABLY YOU DON'T WANT THIS UNLESS YOU DON'T NEED THIS SECURITY PACKAGE
	}
	return tokensecret
}
