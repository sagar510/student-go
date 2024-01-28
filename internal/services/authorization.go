package controller

import (
	"encoding/json"
	"log"
	"net/http"
)



func ValidateToken(token string) (bool, string) {
	url := "http://localhost:3000/verifytoken"
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error forwarding token to rails : ", err)
		return false, ""
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Token Verification failed")
		return false, ""
	}

	var responseData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseData)

	if err != nil {
		log.Println("Error decoding role", err)
		return false, ""
	}

	role, ok := responseData["role"].(string)

	if !ok {
		log.Println("Role does not found in response")
		return false, ""
	}

	return true, role
}
