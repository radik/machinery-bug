package main

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	mt "github.com/RichardKnop/machinery/v1/tasks"
	"log"
)

func PingTask() error {
	fmt.Println("Pong")
	return nil
}

func GetMachineryServer() (*machinery.Server, error) {
	log.Println("init task server")

	taskServer, err := machinery.NewServer(&config.Config{
		Broker:        "redis://localhost:6379",
		ResultBackend: "redis://localhost:6379",
	})
	if err != nil {
		return nil, err
	}

	err = taskServer.RegisterTasks(map[string]interface{}{
		"PingTask": PingTask,
	})

	if err != nil {
		return nil, err
	}

	// Trigger tasks periodically
	err = taskServer.RegisterPeriodicTask(
		"*/1 * * * *",
		"periodic-notifier",
		&mt.Signature{Name: "PingTask"},
	)

	if err != nil {
		return nil, err
	}

	return taskServer, nil
}


func main() {
	taskServer, err := GetMachineryServer()

	if err != nil {
		log.Fatal(err)
	}
	worker := taskServer.NewWorker("machinery_worker", 10)
	if err := worker.Launch(); err != nil {
		log.Fatal(err)
	}
}
