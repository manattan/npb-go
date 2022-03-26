package main

import (
	"fmt"
	"os"

	"github.com/manattan/npb_api/database"
	"github.com/manattan/npb_api/usecase"
	"github.com/manattan/npb_api/web"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		fmt.Printf("shutting down the server with error: %v", err)
		os.Exit(1)
	}

	teamRepo := database.NewTeamRepository(db)

    userUC := usecase.NewTeamUseCase(teamRepo)

	s := web.NewServer(userUC)

	if err := s.Start(":1323"); err != nil {
		fmt.Printf("shutting down the server with error: %v", err)
		os.Exit(1)
	}
}