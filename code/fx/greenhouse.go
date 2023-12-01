package fx

import (
	"encoding/json"
	"fmt"
)

func ParseGreenhouse(url string) GreenHouse {
	jobs := GreenHouse{}
	respBody := FetchResponse(url)

	jsonErr := json.Unmarshal(respBody, &jobs)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	return jobs
}
