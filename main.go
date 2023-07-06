package main

import (
	"os"

	"github.com/jared-weinberger/ChoreRotationExtension/api"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}
	todoistUrl := os.Getenv("TODOIST_REST_API_URL")
	todoistToken := os.Getenv("TODOIST_REST_API_TOKEN")
	projectId := os.Getenv("TODOIST_PROJECT_ID")
	todoistApi := api.MakeClient(todoistUrl, todoistToken)
	collaborators, err := todoistApi.GetProjectCollaborators(projectId)
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting collaborators")
	}
	log.Info().Interface("collaborators", collaborators).Msg("Collaborators successfully fetched")
}
