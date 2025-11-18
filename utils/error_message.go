package utils

import "errors"

const (
	Unauthorized   = "Unauthorized"
	InternalServer = "Internal Server Error"
	BadInput       = "Format data not valid"
)

var (
	ErrGetData            = errors.New("gagal saat mengambil data")
	ErrEmailTaken         = errors.New("email sudah digunakan")
	ErrUsernameTaken      = errors.New("username sudah digunakan")
	ErrPlaceIDUniqueTaken = errors.New("id tempat sudah digunakan, id harus unique")

	//
	ErrUsernameEmpty        = errors.New("username tidak boleh kosong")
	ErrEmailEmpty           = errors.New("email tidak boleh kosong")
	ErrPasswordEmpty        = errors.New("password tidak boleh kosong")
	ErrConfirmPassword      = errors.New("konfirmasi password tidak sama dengan password")
	ErrFormatEmail          = errors.New("format email tidak benar")
	ErrUsernameOrEmailTaken = errors.New("email atau username sudah digunakan")
	ErrFormatPassword       = errors.New("password is invalid")
	ErrOtpExpire            = errors.New("otp telah expired")
	ErrOtpNotMatch          = errors.New("otp salah")

	//
	ErrIDNotFound = errors.New("Id tidak ditemukan atau kosong")
)
