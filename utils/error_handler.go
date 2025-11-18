package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func ConverResponse(err error) int {
	log.Printf("Received error: %v", err)

	switch err {
	case ErrGetData:
		return http.StatusBadRequest
	case ErrEmailTaken, ErrUsernameTaken, ErrUsernameOrEmailTaken, ErrPlaceIDUniqueTaken:
		return http.StatusConflict // 409
	case ErrUsernameEmpty, ErrEmailEmpty, ErrPasswordEmpty, ErrConfirmPassword, ErrFormatEmail, ErrFormatPassword:
		return http.StatusBadRequest // 400
	case ErrOtpExpire, ErrOtpNotMatch:
		return http.StatusUnauthorized // 401
	case ErrIDNotFound:
		return http.StatusNotFound // 404
	default:
		return http.StatusInternalServerError
	}
}

func HandleEchoError(err error) (int, string) {
	if _, ok := err.(*gin.Error); ok {
		return http.StatusBadRequest, BadInput
	}
	return http.StatusBadRequest, BadInput
}

func UnauthorizedError(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"message": Unauthorized})
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": InternalServer})
}

func JWTErrorHandler(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{"message": InternalServer})
}

func ParsePQError(err error) error {
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code == "23505" { // unique_violation
			switch pqErr.Constraint {
			case "users_email_key":
				return ErrEmailTaken
			case "users_username_key":
				return ErrUsernameTaken
			case "tempat_pariwisata_place_id_key":
				return ErrPlaceIDUniqueTaken
			}
		}
		// return fmt.Errorf("terjadi kesalahan saat menyimpan data")
	}
	return err
}
