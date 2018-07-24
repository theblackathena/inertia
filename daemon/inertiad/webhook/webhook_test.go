package webhook

import (
	"bytes"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// The endpoint does not really matter, we are only interested in
// how the request body gets parsed by the Webhook package
func getMockRequest(endpoint string, rawBody []byte) *http.Request {
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(rawBody))
	req.Header.Add("Content-Type", "application/json")
	return req
}
func TestParse(t *testing.T) {
	testCases := []struct {
		reqBody     []byte
		eventHeader string
		eventValue  string
	}{
		{githubPushRawJSON, "x-github-event", GithubPushHeader},
		{gitlabPushRawJSON, "x-gitlab-event", GitlabPushHeader},
		{bitbucketPushRawJSON, "x-event-key", BitbucketPushHeader},
	}

	for _, tc := range testCases {
		req := getMockRequest("/webhook", tc.reqBody)
		req.Header.Add(tc.eventHeader, tc.eventValue)

		// Special case for Bitbucket because Bitbucket
		if tc.eventHeader == "x-event-key" {
			req.Header.Add("User-Agent", "Bitbucket")
		}

		payload, err := Parse(req, os.Stdout)
		assert.Nil(t, err)

		assert.Equal(t, "push", payload.GetEventType())
		assert.Equal(t, "inertia-deploy-test", payload.GetRepoName())
		assert.Equal(t, "refs/heads/master", payload.GetRef())
	}
}

func TestParseDocker(t *testing.T) {
	req := getMockRequest("/docker-webhook", dockerPushRawJSON)
	payload, err := ParseDocker(req, os.Stdout)
	assert.Nil(t, err)

	assert.Equal(t, "briannguyen", payload.GetPusher())
	assert.Equal(t, "latest", payload.GetTag())
	assert.Equal(t, "ubclaunchpad/inertia", payload.GetRepoName())
	assert.Equal(t, "inertia", payload.GetName())
	assert.Equal(t, "ubclaunchpad", payload.GetOwner())
}