package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type TodoistClient struct {
	baseUrl string
	token   string
}

func MakeClient(baseUrl string, token string) *TodoistClient {
	return &TodoistClient{baseUrl: baseUrl, token: token}
}

type TodoistCollaborator struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (client *TodoistClient) sendRequest(method string, path string, body io.Reader) (*http.Response, error) {
	req_url, err := url.JoinPath(client.baseUrl, path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", req_url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.token))
	return http.DefaultClient.Do(req)
}

func MakeEntityFromResponse[T any](res *http.Response, entity *T) error {
	var buffer bytes.Buffer
	_, bodyerr := io.Copy(&buffer, res.Body)
	if bodyerr != nil {
		return bodyerr
	}
	return json.Unmarshal(buffer.Bytes(), entity)
}

func (client *TodoistClient) GetProjectCollaborators(projectId string) ([]TodoistCollaborator, error) {
	path := fmt.Sprintf("/projects/%s/collaborators", projectId)
	res, err := client.sendRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		errMsg, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("%s to %s failed with %d: %s", res.Request.Method, res.Request.URL, res.StatusCode, string(errMsg))
	}
	var collaborators []TodoistCollaborator
	err = MakeEntityFromResponse(res, &collaborators)
	if err != nil {
		return nil, err
	}
	return collaborators, nil
}
