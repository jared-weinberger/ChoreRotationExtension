package api

import (
	"errors"
	"net/http"
	"testing"
)

type MockHttpClient struct {
	GetFunc func(client *MockHttpClient, url string) (*http.Response, error)
}

func (client *MockHttpClient) Get(url string) (*http.Response, error) {
	return client.GetFunc(client, url)
}
func PropogateApiError(t *testing.T) {
	mockClient := MockHttpClient{GetFunc: func(_ *MockHttpClient, url string) (*http.Response, error) {
		return nil, errors.New("Ur bad")
	}}
	projectId := "1234"
	testApi := TodoistApiImpl{HttpClient: &mockClient}
	collaborators, err := testApi.GetProjectCollaborators(projectId)
	if err == nil {
		t.Fail()
	}
}
