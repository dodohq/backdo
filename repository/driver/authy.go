package driver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/dodohq/backdo/models"
)

type authyResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// RequestVerificationCode get phone number verification code
func (r *privateDriverRepo) RequestVerificationCode(via string, countryCode int, phoneNumber string) *models.HTTPError {
	if via != "sms" && via != "call" {
		return models.NewErrorUnprocessableEntity("Verification Mode Invalid")
	}
	urlStr := `https://api.authy.com/protected/json/phones/verification/start`
	data := map[string]interface{}{
		"api_key":      os.Getenv("TWILIO_API_KEY"),
		"via":          via,
		"country_code": countryCode,
		"phone_number": phoneNumber,
	}

	return r.authyPostReq(urlStr, data)
}

// VerifyAuthyCode verify the code with phone number
func (r *privateDriverRepo) VerifyAuthyCode(countryCode int, phoneNumber, verificationCode string) *models.HTTPError {
	urlStr := `https://api.authy.com/protected/json/phones/verification/check`
	data := map[string]interface{}{
		"api_key":           os.Getenv("TWILIO_API_KEY"),
		"country_code":      countryCode,
		"phone_number":      phoneNumber,
		"verification_code": verificationCode,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return models.NewErrorInternalServer(err.Error())
	}
	resp, err := http.Post(urlStr, "application/json", bytes.NewReader(body))
	if err != nil {
		return models.NewErrorInternalServer(err.Error())
	}
	defer resp.Body.Close()

	return r.authyPostReq(urlStr, data)
}

func (r *privateDriverRepo) authyPostReq(urlStr string, data map[string]interface{}) *models.HTTPError {
	body, err := json.Marshal(data)
	if err != nil {
		return models.NewErrorInternalServer(err.Error())
	}
	resp, err := http.Post(urlStr, "application/json", bytes.NewReader(body))
	if err != nil {
		return models.NewErrorInternalServer(err.Error())
	}
	defer resp.Body.Close()

	aRes := new(authyResp)
	if err := json.NewDecoder(resp.Body).Decode(aRes); err != nil {
		return models.NewErrorInternalServer(err.Error())
	}
	if !aRes.Success {
		return models.NewErrorInternalServer(aRes.Message)
	}

	return nil
}
