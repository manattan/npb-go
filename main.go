package main

import (
	"fmt"
	"os"

	"github.com/manattan/npb_go/database"
	"github.com/manattan/npb_go/usecase"
	"github.com/manattan/npb_go/web"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		fmt.Printf("shutting down the server with error: %v", err)
		os.Exit(1)
	}

	teamRepo := database.NewTeamRepository(db)
	playerRepo := database.NewPlayerRepository(db)

	userUC := usecase.NewTeamUseCase(teamRepo)
	playerUC := usecase.NewPlayerUseCase(playerRepo)

	s := web.NewServer(userUC, playerUC)

	if err := s.Start(":1323"); err != nil {
		fmt.Printf("shutting down the server with error: %v", err)
		os.Exit(1)
	}
}
