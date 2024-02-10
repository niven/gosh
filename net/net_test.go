package net_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/niven/gosh/net"
)

type TMNT struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Tags []string `json:"tags"`
}

func TestLoadJson(t *testing.T) {

	input := `{
        "name": "Rafael",
        "age": 15,
        "tags": ["ninjutsu", "pizza", "red"]
    }`
	expected := TMNT{}
	json.Unmarshal([]byte(input), &expected)

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, input)
	}))
	defer svr.Close()

	turtle := TMNT{}
	httpStatusCode, content, err := net.GetJson(svr.URL, &turtle)

	assert.Nil(t, err)
	assert.Equal(t, httpStatusCode, http.StatusOK)
	assert.Equal(t, input, string(content))
	assert.Equal(t, expected, turtle)

}
