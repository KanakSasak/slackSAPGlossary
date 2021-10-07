package repository

import (
	"SlackSAPGlossary/domain"
	"SlackSAPGlossary/utils"
	"errors"
	"fmt"
	"google.golang.org/api/sheets/v4"
	"log"
)

type CommandRepository struct {
	GheetService *sheets.Service
}

func (c CommandRepository) Find(keyword string) (*[]domain.DataList, error) {
	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1sxZh6Jqk6VPdS8t9l3KDiNn_W7LL1yOmOuqhlGz_VzA/edit
	spreadsheetId := utils.GetEnvVariables("SPREADSHEET_ID")
	readRange := utils.GetEnvVariables("SPREADSHEET_RANGE")
	valueRenderOption := "FORMATTED_VALUE"
	resp, err := c.GheetService.Spreadsheets.Values.Get(spreadsheetId, readRange).ValueRenderOption(valueRenderOption).Do()
	if err != nil {
		return nil, errors.New("Unable to retrieve data from sheet: " + err.Error())

	}

	data := make([]domain.DataList, 0)
	if len(resp.Values) == 0 {
		return nil, errors.New("No data found.")
	} else {

		for _, row := range resp.Values {
			var x = ""
			var y = ""
			var z = ""
			//Print columns A and E, which correspond to indices 0 and 4.

			if len(row) == 1 {
				x = fmt.Sprintf("%v", row[0])
				y = ""
				z = ""
			}

			if len(row) == 2 {
				x = fmt.Sprintf("%v", row[0])
				y = fmt.Sprintf("%v", row[1])
				z = ""
			}

			if len(row) == 3 {
				x = fmt.Sprintf("%v", row[0])
				y = fmt.Sprintf("%v", row[1])
				z = fmt.Sprintf("%v", row[2])
			}

			data = append(data, domain.DataList{
				Keyword:     x,
				Description: y,
				LinkDetails: z,
			})
		}

		////data2 := &domain.DataList{}
		////byteData,err := json.Marshal(resp.)
		////if err != nil {
		////	return nil,err
		////}
		//byteData,_ := resp.MarshalJSON()
		//var dat map[string]interface{}
		//if err := json.Unmarshal(byteData, &dat); err != nil {
		//	panic(err)
		//}
		//fmt.Println(dat["values"])
		////json.Unmarshal(byteData,&data2)
		////log.Println(data2)

	}

	return &data, err
}

func (c CommandRepository) ValidateNullString(data interface{}) string {
	if data == nil {
		log.Println("masuk2")
		return ""
	}

	return fmt.Sprintf("%v", data)
}

func NewCommandRepository(srv *sheets.Service) domain.CommandRepository {
	return &CommandRepository{
		srv,
	}
}
