package tripay_payment

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func CreateTransaction(transactionRequest TransactionRequest) (map[string]interface{}, error) {
	// create HTTP req
	apiKey := os.Getenv("TRIPAY_API_KEY")
	appMode := os.Getenv("APP_MODE")
	urlTripay := "https://tripay.co.id/api/transaction/create"
	if appMode == "dev" {
		urlTripay = "https://tripay.co.id/api-sandbox/transaction/create"
	}
	// authToken := base64.StdEncoding.EncodeToString([]byte(apiKey + ":"))
	headers := map[string]string{
		"Authorization": "Bearer " + apiKey,
		"Content-Type":  "application/json",
	}

	body, err := json.Marshal(transactionRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", urlTripay, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// get response
	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("tripay error: status %d, response %s", resp.StatusCode, string(bodyResp))
	}

	// parse response to json
	var bodyRespJson map[string]interface{}
	if err := json.Unmarshal(bodyResp, &bodyRespJson); err != nil {
		return nil, err
	}

	return bodyRespJson, nil
}

func CreateSignature(merchantCode, merchantRef, amount, privateKey string) string {
	message := merchantCode + merchantRef + amount
	h := hmac.New(sha256.New, []byte(privateKey))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func VerifyCallback(r *http.Request, req TransactionCallback) error {
	privateKey := os.Getenv("TRIPAY_PRIVATE_KEY")
	callbackSignature := r.Header.Get("X-Callback-Signature")

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("cannot marshal request: %w", err)
	}

	h := hmac.New(sha256.New, []byte(privateKey))
	h.Write(jsonBytes)
	signature := hex.EncodeToString(h.Sum(nil))

	if !hmac.Equal([]byte(signature), []byte(callbackSignature)) {
		return fmt.Errorf("invalid callback signature")
	}

	return nil
}
