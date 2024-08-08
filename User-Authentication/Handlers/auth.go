package auth

import (
    "net/http"
	"user-auth-api/storage"
	"user-auth-api/models"

    "github.com/labstack/echo/v4"
    "golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func verifyPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func registerHandler(c echo.Context) error {
    var user User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
    }
    
    hashedPassword, err := hashPassword(user.Password)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error hashing password"})
    }

    user.Password = hashedPassword
    _, err = DB.NamedExec(`INSERT INTO users (username, email, password, firstname, lastname) 
                            VALUES (:username, :email, :password, :firstname, :lastname)`, user)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error creating user"})
    }

    return c.JSON(http.StatusCreated, user)
}

func loginHandler(c echo.Context) error {
    var user User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
    }

    var dbUser User
    err := DB.Get(&dbUser, "SELECT * FROM users WHERE username=$1 OR email=$2", user.Username, user.Email)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, echo.Map{"error": "User not found"})
    }

    if err := verifyPassword(dbUser.Password, user.Password); err != nil {
        return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid password"})
    }

    return c.JSON(http.StatusOK, echo.Map{"uuid": dbUser.UUID})
}
