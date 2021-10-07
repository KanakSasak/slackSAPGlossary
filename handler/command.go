package handler

import (
	"SlackSAPGlossary/domain"
	"github.com/shomali11/slacker"
	"log"
)

type CommandHandler struct {
	CommandService domain.CommandService
}

func (h CommandHandler) PingHandler() (string, error) {
	data, err := h.CommandService.Ping()
	if err != nil {
		return "", err
	}
	return data, nil
}

func (h CommandHandler) HelpHandler() (string, error) {
	data, err := h.CommandService.Help()
	if err != nil {
		return "", err
	}
	return data, nil
}

func (h CommandHandler) FindHandler(word string) (string, error) {
	data, err := h.CommandService.Find(word)
	if err != nil {
		return "", err
	}
	return data, nil
}

func NewCommandHandler(r *slacker.Slacker, us domain.CommandService) {
	handler := &CommandHandler{
		CommandService: us,
	}

	r.Init(func() {
		log.Println("Connected!")
	})

	r.Err(func(err string) {
		log.Println(err)
	})

	r.DefaultCommand(func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {

		response.Reply(":man-bowing: Sorry, what? \n" +
			"\n" +
			"\n" +
			"\n" +
			"\n" +
			"Find the acronym you want to know, for example, find the list of `abap library`, just type: `find abap library` \n" +
			"Type `help` for more usage")
	})

	r.Command("find <word>", &slacker.CommandDefinition{
		Description: "find a word!",
		Example:     "find abap library",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			word := request.Param("word")
			data, err := handler.FindHandler(word)
			if err != nil {
				response.ReportError(err)
			}
			response.Reply(data)

		},
	})

	r.Command("ping", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			data, err := handler.PingHandler()
			if err != nil {
				response.ReportError(err)
			}
			response.Reply(data)
		},
	})

}
