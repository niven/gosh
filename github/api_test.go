package github

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRepositorySecretList(t *testing.T) {
	input := `
	{
		"total_count": 2,
		"secrets": [
		  {
			"name": "GH_TOKEN",
			"created_at": "2019-08-10T14:59:22Z",
			"updated_at": "2020-01-10T14:59:22Z"
		  },
		  {
			"name": "GIST_ID",
			"created_at": "2020-01-10T10:59:22Z",
			"updated_at": "2020-01-11T11:59:22Z"
		  }
		]
	  }
	`
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, input)
	}))
	defer svr.Close()
	secrets, err := GetRepositorySecretList(svr.URL, "token")
	assert.Nil(t, err)
	assert.Equal(t, []string{"GH_TOKEN", "GIST_ID"}, secrets)
}
