package routes

import (
	"awesomeProject/dataBase"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	emailverifier "github.com/AfterShip/email-verifier"
)

var (
	verifier = emailverifier.
		NewVerifier().
		EnableAutoUpdateDisposable()
)

type user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewMiddleware() fiber.Handler {
	return AuthMiddleware
}

func AuthMiddleware(c *fiber.Ctx) error {

	sess, err := store.Get(c)

	if strings.Split(c.Path(), "/")[1] == "auth" {
		return c.Next()
	}

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	if sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	return c.Next()
}

func Register(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var data user

	err := c.BodyParser(&data)

	if err != nil {
		fmt.Printf("error BodyParser")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})

	}

	index := strings.Index(data.Email, "@")
	//fmt.Printf("index", index)
	domain := data.Email[index+1:]
	if verifier.IsDisposable(domain) {
		fmt.Printf("%s is a disposable domain\n", domain)

		fmt.Printf("Error temp mail")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error email: ",
		})
	}

	fmt.Printf("%s is not a disposable domain\n", domain)

	password, bcErr := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if bcErr != nil {
		fmt.Printf("error Password")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}

	user := dataBase.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(password),
	}
	err = dataBase.CreateUser(&user)
	if err != nil {
		fmt.Printf("error DataBase")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "registered.",
	})

}

func Login(c *fiber.Ctx) error {
	var data user

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}

	var user dataBase.User
	if !dataBase.CheckEmail(data.Email, &user) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized. email not in db",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized. password not good. " + err.Error(),
		})
	}

	sess, sessErr := store.Get(c)
	if sessErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}

	sess.Set(AUTH_KEY, true)
	sess.Set(USER_ID, user.ID)

	sessErr = sess.Save()
	if sessErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logged in",
	})
}

func Logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "logged out (no session)",
		})
	}

	err = sess.Destroy()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logged out",
	})
}

func HealthCheck(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	auth := sess.Get(AUTH_KEY)

	if auth != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "authenticated",
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}
}

func GetUser(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	if sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	userId := sess.Get(USER_ID)
	if userId == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	var user dataBase.User
	user, err = dataBase.GetUser(fmt.Sprint(userId))
	fmt.Printf("user: %v\n", user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}
	user.Password = ""

	return c.Status(fiber.StatusOK).JSON(user)
}
