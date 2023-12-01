package fx

import (
	"fmt"
	"io"
	"net/http"
)

func FetchResponse(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	respBody, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println("error")
	}

	return respBody
}
