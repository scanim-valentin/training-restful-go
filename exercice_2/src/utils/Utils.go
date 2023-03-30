package utils

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
)

type DummyStruct struct {
	Dummy int
}

var charset = []byte("azertyuiopqsdfghjklmwxcvbnAZERTYUIOPQSDFHJKLMWXCVBN0123456789")

func RandomString(nbMaxChar int) []byte {
	nbChar := rand.Intn(nbMaxChar)
	content := make([]byte, 0)
	for k := 0; k < nbChar; k++ {
		content = append(content, charset[rand.Intn(len(charset)-1)])
	}
	return content
}
func SQLErrorToHTTPStatus(method string, err error) int {
	fmt.Print("\n SQL Error: ", err)
	fmt.Print("\n Converting SQL Error to HTTP Status: ")
	var s int
	switch err {
	case nil:
		switch method {
		case http.MethodPut:
			fallthrough
		case http.MethodPost:
			s = http.StatusCreated
			fmt.Println(http.StatusCreated)
			break
		default:
			s = http.StatusOK
			fmt.Println(http.StatusOK)

			break
		}
	case sql.ErrNoRows:
		s = http.StatusNotFound
		fmt.Println(http.StatusNotFound)

		break
	default:
		s = http.StatusInternalServerError
		fmt.Println(http.StatusInternalServerError)

	}
	return s
}
