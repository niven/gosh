package integrations_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/niven/gosh/integrations"
	"github.com/stretchr/testify/assert"
)

func TestVaultStructs(t *testing.T) {

	input := `
	{
		"auth": null,
		"data": {
		  "foo": "bar",
		  "ttl": "1h"
		},
		"lease_duration": 3600,
		"lease_id": "",
		"renewable": false
	  }
	`

	result := integrations.VaultKeyValue{}
	err := json.Unmarshal([]byte(input), &result)
	assert.Nil(t, err)
	assert.Equal(t, "bar", result.Data["foo"])
	assert.Equal(t, 3600, result.LeaseDuration)

	input = `
	{
		"errors": [
		  "message",
		  "another message"
		]
	  }
	`

	errResult := integrations.VaultError{}
	err = json.Unmarshal([]byte(input), &errResult)
	assert.Nil(t, err)
	assert.Equal(t, len(errResult.Errors), 2)
	assert.Equal(t, errResult.Errors[0], "message")

}

func TestVaultGetKeyValue(t *testing.T) {

	input := `
	{
		"auth": null,
		"data": {
		  "username": "root",
		  "password": "hunter2"
		},
		"lease_duration": 3600,
		"lease_id": "",
		"renewable": false
	  }
	`
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, input)
	}))
	defer svr.Close()

	vkv, err := integrations.VaultGetKeyValue(svr.URL, "/secret/ignored", "no-token-for-test")

	assert.Nil(t, err)
	assert.Equal(t, "hunter2", vkv.Data["password"])

}
