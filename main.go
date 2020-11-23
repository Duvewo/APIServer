package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"os"
)

func main(){

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(403, "Access denied")
	})
	e.GET("/api/encrypt/:type/:text", encrypt)
	//e.GET("/api/decrypt/{text}", decrypt)

	err := e.Start(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}

}

func encrypt(c echo.Context) error {
	hashType := c.Param("type")
	encryptText := c.Param("text")

	switch hashType {
	case "md5":
		hash := md5.New()
		io.WriteString(hash, encryptText)
		hashedText := fmt.Sprintf("%x", hash.Sum(nil))
		return c.String(200, hashedText)
		break
	case "sha256":
		hash := sha256.New()
		io.WriteString(hash, encryptText)
		hashedText := fmt.Sprintf("%x", hash.Sum(nil))
		return c.String(200, hashedText)
		break
	case "sha512":
		hash := sha512.New()
		io.WriteString(hash, encryptText)
		hashedText := fmt.Sprintf("%x", hash.Sum(nil))
		return c.String(200, hashedText)
		break
	}

	return nil
}