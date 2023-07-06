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

type todoistCollaborator struct {
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
	jsonerr := json.Unmarshal(buffer.Bytes(), entity)
	if jsonerr != nil {
		return jsonerr
	}
	return nil
}

func (client *TodoistClient) GetProjectCollaborators(project_id string) ([]todoistCollaborator, error) {
	path := fmt.Sprintf("/projects/%s/collaborators", project_id)
	res, err := client.sendRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Request failed with status code %d", res.StatusCode)
	}
	var collaborators []todoistCollaborator
	err = MakeEntityFromResponse(res, &collaborators)
	if err != nil {
		return nil, err
	}
	return collaborators, nil
}
