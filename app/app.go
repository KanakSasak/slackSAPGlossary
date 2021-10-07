package app

import (
	"SlackSAPGlossary/database"
	"SlackSAPGlossary/handler"
	"SlackSAPGlossary/repository"
	"SlackSAPGlossary/service"
	"SlackSAPGlossary/utils"
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func Start() {
	bot := slacker.NewClient(utils.GetEnvVariables("SLACK_BOT_TOKEN"), utils.GetEnvVariables("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	database.GsheetConnect()

	srv := database.GetService()
	CmdRepo := repository.NewCommandRepository(srv)
	CmdService := service.NewCommandService(CmdRepo)
	handler.NewCommandHandler(bot, CmdService)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
