package service

import (
	"SlackSAPGlossary/domain"
	"github.com/paddie/gokmp"
	"strings"
)

type CommadService struct {
	commandRepo domain.CommandRepository
}

func (c CommadService) Ping() (string, error) {
	return "PONG!", nil
}

func (c CommadService) Help() (string, error) {

	return "!!", nil
}

func (c CommadService) Find(keyword string) (string, error) {

	data, err := c.commandRepo.Find(keyword)
	if err != nil {
		return "", err
	}

	newData := make([]domain.DataList, 0)
	for _, row := range *data {
		kmp, _ := gokmp.NewKMP(strings.ToUpper(keyword))
		flag := kmp.FindAllStringIndex(strings.ToUpper(row.Keyword))
		if len(flag) >= 1 {
			newData = append(newData, domain.DataList{
				Keyword:     row.Keyword,
				Description: row.Description,
				LinkDetails: row.LinkDetails,
			})
		}
	}

	searchResulttxt := ""
	searchDetailtxt := ""
	if len(newData) != 0 {
		for _, row := range newData {
			searchDetailtxt += "- " + "`" + row.Keyword + "` : " + row.Description + "\n"
		}

		searchResulttxt = ":woman-tipping-hand: Here's what I found for `" + keyword + "`\n" +
			"\n" +
			"\n" +
			"" + searchDetailtxt + ""

	} else {
		searchResulttxt = ":man-bowing: Sorry, not able to find anything with `" + keyword + "` \n" +
			"\n" +
			"\n" +
			"\n" +
			"Please update here in case you find the meaning. Thanks! :pray:"
	}

	return searchResulttxt, nil
}

func NewCommandService(a domain.CommandRepository) domain.CommandService {
	return &CommadService{
		commandRepo: a,
	}
}
