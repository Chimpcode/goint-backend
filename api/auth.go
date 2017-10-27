package api

import (
	"github.com/kataras/iris"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"strings"
	"../db"
	"../types"

	"time"
	"log"
)

const SECRET = "goint2017"

func myHandler(c iris.Context) {
	log.Println("into dashboard handler")
	user := c.Values().Get("jwt").(*jwt.Token)
	//user.Claims.(jwt.MapClaims)["logged"] = sess.Start(c).Get("authenticated")

	c.JSON(user.Claims)
}

func getJWTToken(user *types.User) (string, error) {
	signer := jwt.New(jwt.SigningMethodHS256)

	signer.Claims.(jwt.MapClaims)["iss"] = user.Group
	signer.Claims.(jwt.MapClaims)["exp"] = time.Now().Add(TimeToExpires).Unix()
	signer.Claims.(jwt.MapClaims)["user"] = struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Logged bool `json:"logged"`
		Group string `json:"group"`
	}{user.FullName, user.Email, true, user.Group}

	return signer.SignedString([]byte(SECRET))


}

func LinkAuthApi (auth iris.Party) error {

	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {

			return []byte(SECRET), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
	})

	auth.Any("/login", func(c iris.Context) {

		session := sess.Start(c)

		data := struct {
			Username string `json:"username"`
			Email string `json:"email"`
			Password string `json:"password"`
		}{}

		err := c.ReadJSON(&data)


		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{"error": err.Error()})
			return
		}

		if !strings.EqualFold(data.Email, "") {
			// Search by Email
			user, err := db.GetUserByEmail(strings.ToLower(data.Email))

			if err != nil {
				if strings.Contains(err.Error(), "exist") {
					c.StatusCode(iris.StatusForbidden)
					c.JSON(iris.Map{"error": err.Error()})
					return
				}
				c.StatusCode(iris.StatusInternalServerError)
				c.JSON(iris.Map{"error": err.Error()})
				return
			}

			if strings.EqualFold(user.Password, strings.ToLower(data.Password)) {
				// good data
				token, err := getJWTToken(user)
				if err != nil {
					c.StatusCode(iris.StatusInternalServerError)
					c.JSON(iris.Map{"error": err.Error()})
					return
				}
				c.JSON(iris.Map{"token": token})
				session.Set("authenticated", true)
				return

			} else {
				c.StatusCode(iris.StatusForbidden)
				c.JSON(iris.Map{"error": "Invalid credentials"})
				return
			}
		} else if !strings.EqualFold(data.Username, "") {
			// Search by Username
			user, err := db.GetUserByUsername(strings.ToLower(data.Username))

			if err != nil {
				if strings.Contains(err.Error(), "exist") {
					c.StatusCode(iris.StatusForbidden)
					c.JSON(iris.Map{"error": err.Error()})
					return
				}
				c.StatusCode(iris.StatusInternalServerError)
				c.JSON(iris.Map{"error": err.Error()})
				return
			}

			if strings.EqualFold(user.Password, strings.ToLower(data.Password)) {
				// good data
				token, err := getJWTToken(user)
				if err != nil {
					c.StatusCode(iris.StatusInternalServerError)
					c.JSON(iris.Map{"error": err.Error()})
					return
				}
				c.JSON(iris.Map{"token": token})
				session.Set("authenticated", true)
				return


			} else {
				c.StatusCode(iris.StatusForbidden)
				c.JSON(iris.Map{"error": "Invalid credentials"})
				return
			}

		} else {
			c.StatusCode(iris.StatusForbidden)
			c.JSON(iris.Map{"error": "Invalid credentials"})
			return
		}
	})

	auth.Post("/logout", func(c iris.Context) {
		session := sess.Start(c)
		// Revoke users authentication
		session.Set("authenticated", false)
	})



	auth.Use(jwtHandler.Serve)

	auth.Any("/dashboard", myHandler)



	return nil
}
