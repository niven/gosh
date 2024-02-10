package integrations

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/niven/gosh/net"
)

type VaultError struct {
	Errors []string `json:"errors"`
}

type VaultKeyValue struct {
	Auth          string            `json:"auth"`
	Data          map[string]string `json:"data"`
	LeaseDuration int               `json:"lease_duration"`
	LeaseId       string            `json:"lease_id"`
	Renewable     bool              `json:"renewable"`
}

func VaultGetKeyValue(vaultApiUrl, secretPath, vaultToken string) (VaultKeyValue, error) {

	result := VaultKeyValue{}
	url := fmt.Sprintf("%s%s", vaultApiUrl, secretPath)
	httpStatusCode, content, err := net.GetJsonWithHeaders(url, http.Header{
		"X-Vault-Token": {vaultToken},
		"Content-Type":  {"application/json"},
	}, &result)

	if err != nil {
		return VaultKeyValue{}, err
	}
	if httpStatusCode != http.StatusOK {
		vaultError := VaultError{}
		err = json.Unmarshal(content, &vaultError)
		if err != nil {
			return VaultKeyValue{}, err
		}
		return VaultKeyValue{}, errors.New(fmt.Sprintf("Vault error: %v", vaultError.Errors))
	}
	return result, nil
}
