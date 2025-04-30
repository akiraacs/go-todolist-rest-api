package main

import (
    "log"

    "github.com/akiraacs/go-todolist-rest-api/internal/api"
    "github.com/akiraacs/go-todolist-rest-api/internal/api/handlers"
    "github.com/akiraacs/go-todolist-rest-api/internal/db"
    "github.com/akiraacs/go-todolist-rest-api/internal/repository"
    "github.com/akiraacs/go-todolist-rest-api/internal/usecase"
)

func main() {
    log.Println("Starting application...")

    // Configura o banco de dados
    database := db.ConnectDB()

    // Inicializa as dependências
    taskRepo := repository.NewTaskRepository(database)
    taskUseCase := usecase.NewTaskUseCase(taskRepo)
    taskHandler := handlers.NewTaskHandler(taskUseCase)

    // Configura as rotas
    router := api.SetupRoutes(taskHandler)

    // Inicia o servidor
    router.Run(":8080")
}
