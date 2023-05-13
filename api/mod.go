package api

import (
	"fmt"

	"github.com/jared-weinberger/ChoreRotationExtension/model"
)

const ApiBaseUrl string = "https://api.todoist.com/rest/v2"

func getProjectBaseUrl() string {
	return ApiBaseUrl + "/projects"
}

func GetProjectCollaborators(project_id string) model.Collaborator {
	url := fmt.Sprintf("%s/%s", getProjectBaseUrl(), project_id)

}
