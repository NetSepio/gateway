package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type GoogleUserResult struct {
	Email string `json:"email"`
}

func GetGoogleUserEmail(access_token string) (string, error) {
	rootUrl := "https://www.googleapis.com/oauth2/v1/userinfo?alt=json"

	client := http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest("GET", rootUrl, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		// log body
		body, _ := io.ReadAll(res.Body)
		fmt.Printf("Response body: %s\n", body)
		return "", errors.New("could not retrieve user")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}

	var user GoogleUserResult
	err = json.Unmarshal(body, &user)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}
