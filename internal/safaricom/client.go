package safaricom

import (
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	consumerKey    = "bIjWx0kLeV2M3p4XGvuEtB37gdgtA7u3LHxuue3RbQraQ8GE"
	consumerSecret = "ya3TXweYB1G66bNTnyNrcmmiKVNmcloGGqQayf5KKmqKGrKAoL3ezlejdur2LZTJ"
	authURL        = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	stkURL         = "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
)

func GetAccessToken() (string, error) {
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", consumerKey, consumerSecret)))

	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", "Basic "+auth).
		SetHeader("Content-Type", "application/json").
		Get(authURL)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := resp.Unmarshal(&result); err != nil {
		return "", err
	}

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access_token not found in response")
	}
	return token, nil
}

func MakeSTKPush(token string) ([]byte, int, error) {
	timestamp := GenerateTimestamp()
	password := GeneratePassword(timestamp)

	body := map[string]interface{}{
		"BusinessShortCode": "174379",
		"Password":          password,
		"Timestamp":         timestamp,
		"TransactionType":   "CustomerPayBillOnline",
		"Amount":            "1",
		"PartyA":            "254792867200",
		"PartyB":            "174379",
		"PhoneNumber":       "254792867200",
		"CallBackURL":       "https://2dc5-105-27-123-2.ngrok-free.app/stk",
		"AccountReference":  "Test",
		"TransactionDesc":   "Test",
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(stkURL)

	if err != nil {
		return nil, 0, err
	}

	return resp.Body(), resp.StatusCode(), nil
}
