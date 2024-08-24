package controllers

import (
	"context"
	"fmt"
	"os"
	"sample-auth-backend/app/models"
	"sample-auth-backend/pkg/utils"
	"sample-auth-backend/platform/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SignUp(c *fiber.Ctx) error {
	db, err := database.OpenDbConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	user := &models.User{}
	// Check, if received JSON data is valid.
	if err := c.BodyParser(user); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.Active = true
	user.LoginType = "simple"
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), 14)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
			"token": nil,
		})
	}
	user.PasswordHash = string(bytes)
	if user.Username == "" {
		user.Username = utils.GenerateRandomUsername()
	}
	validate := utils.NewValidator()
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ValidatorError(err),
			"token": nil,
		})
	}
	// Delete book by given ID.
	if err := db.CreateUser(user); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"token": nil,
		})
	}

	token, err := utils.GenerateNewAccessToken(map[string]interface{}{
		"id":    user.ID,
		"name":  user.Username,
		"email": user.Email,
	})

	if err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"token": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": nil,
		"token": token,
	})
}

func SignIn(c *fiber.Ctx) error {
	db, err := database.OpenDbConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"token": nil,
		})
	}
	signDetails := &models.SignIn{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(signDetails); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
			"token": nil,
		})
	}

	validate := utils.NewValidator()
	if err := validate.Struct(signDetails); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.ValidatorError(err),
			"token": nil,
		})
	}

	if err := validate.Var(signDetails, "username_or_email"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
			"token": nil,
		})
	}

	identifer := ""
	if signDetails.Email != "" {
		identifer = signDetails.Email
	} else {
		identifer = signDetails.Username
	}
	user, err := db.GetUserWithEmailOrUserName(
		identifer,
		identifer,
	)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
			"token": nil,
		})
	}

	// Check if the password matches
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(signDetails.Password),
	); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Password Mismatch. Invalid username or password",
			"token": nil,
		})
	}

	token, err := utils.GenerateNewAccessToken(map[string]interface{}{
		"id":    user.ID,
		"name":  user.Username,
		"email": user.Email,
	})

	if err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"token": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": nil,
		"token": token,
	})
}

func LoggedIn(c *fiber.Ctx) error {
	tokenMetaData, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
			"user":  nil,
		})
	}
	db, err := database.OpenDbConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"user":  nil,
		})
	}

	identifer := tokenMetaData.Email
	user, err := db.GetUserWithEmailOrUserName(
		identifer,
		identifer,
	)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
			"user":  nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": nil,
		"user":  user.ToUserResponse(),
	})
}

func GoogleCallbackHandler(c *fiber.Ctx) error {
	code := c.Query("code")
	googlOauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"profile", "email"}, // Adjust scopes as needed
		Endpoint:     google.Endpoint,
	}

	googleToken, err := googlOauthConfig.Exchange(context.Background(), code)

	if err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"token": nil,
		})
	}
	userInfo, err := utils.GetGoogleAutherUserInfo(googleToken.AccessToken)

	if err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"token": nil,
		})
	}

	db, err := database.OpenDbConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	fmt.Println(userInfo, userInfo["name"].(string), userInfo["email"].(string))
	dbUser, err := db.GetUserWithEmailOrUserName(
		userInfo["email"].(string),
		userInfo["email"].(string),
	)

	user := &models.User{}
	if dbUser == nil {
		user.ID = uuid.New()
		user.CreatedAt = time.Now()
		user.Active = true
		validate := utils.NewValidator()
		user.LoginType = "google"
		user.Name = userInfo["name"].(string)
		user.Email = userInfo["email"].(string)

		if err := validate.Struct(user); err != nil {
			// Return, if some fields are not valid.
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   utils.ValidatorError(err),
				"token": nil,
			})
		}
		if user.Username == "" {
			user.Username = utils.GenerateRandomUsername()
		}
		// Delete book by given ID.
		if err := db.CreateUser(user); err != nil {
			// Return status 500 and error message.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
				"token": nil,
			})
		}
	} else {
		user = dbUser
	}

	accessToken, err := utils.GenerateNewAccessToken(map[string]interface{}{
		"id":    user.ID,
		"name":  user.Username,
		"email": user.Email,
	})

	if err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"token": nil,
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"token": accessToken,
	})
}
