package github

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/niven/gosh/util"
)

const GithubApiListRespositorySecrets = "https://api.github.com/repos/%s/%s/actions/secrets"

func GithubApiCall(url, token string, result any) error {

	httpStatusCode, _, err := util.GetJsonWithHeaders(url, http.Header{
		"Accept":               {"application/vnd.github+json"},
		"Authorization-Type":   {fmt.Sprintf("Bearer %s", token)},
		"X-GitHub-Api-Version": {"2022-11-28"},
	}, result)

	if httpStatusCode != http.StatusOK || err != nil {
		return errors.New(fmt.Sprintf("HTTP %d: %v", httpStatusCode, err))
	}

	return nil
}

type Secret struct {
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type SecretResponse struct {
	TotalCount int      `json:"total_count"`
	Secrets    []Secret `json:"secrets"`
}

func GetRepositorySecretList(url, token string) (secretNames []string, err error) {

	secretResponse := SecretResponse{}
	err = GithubApiCall(url, token, &secretResponse)
	if err != nil {
		return nil, err
	}
	secretNames = make([]string, secretResponse.TotalCount)
	for i, s := range secretResponse.Secrets {
		secretNames[i] = s.Name
	}
	return secretNames, nil
}
