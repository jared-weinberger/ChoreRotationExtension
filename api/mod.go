package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/jared-weinberger/ChoreRotationExtension/model"
)

type HttpClient interface {
	Get(url string) (*http.Response, error)
}

type TodoistApi interface {
	GetProjectCollaborators(project_id string) ([]model.Collaborator, error)
}

const TodoistApiBaseUrl string = "https://api.todoist.com/rest/v2"

type todoistCollaborator struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (self *todoistCollaborator) GetName() string {
	return self.Name
}
func (self *todoistCollaborator) GetEmail() string {
	return self.Name
}

type TodoistApiImpl struct {
	HttpClient HttpClient
}

func (self *TodoistApiImpl) get(urlfrag string) (*http.Response, error) {
	return self.HttpClient.Get(TodoistApiBaseUrl + urlfrag)
}

func (self *TodoistApiImpl) GetProjectCollaborators(project_id string) ([]model.Collaborator, error) {
	res, httperr := self.get("/projects/" + project_id + "/collaborators")
	if httperr != nil {
		return nil, httperr
	}

	var buffer bytes.Buffer
	_, bodyerr := io.Copy(&buffer, res.Body)
	if bodyerr != nil {
		return nil, bodyerr
	}
	var collaborators []todoistCollaborator
	jsonerr := json.Unmarshal(buffer.Bytes(), &collaborators)
	if jsonerr != nil {
		return nil, jsonerr
	}
	result := make([]model.Collaborator, len(collaborators))
	for idx, c := range collaborators {
		result[idx] = &c
	}
	return result, nil
}
