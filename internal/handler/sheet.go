package pkg

import (
	"io/ioutil"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

func WriteToSheet(title, body []string) {
	data, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(err.Error())
	}

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		panic(err.Error())
	}

	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet("1DelUlr7XkqbOmaeKdP1Tu5PjMZKAep_igFqn8iv_oqo")
	if err != nil {
		panic(err.Error())
	}

	sheet, err := spreadsheet.SheetByIndex(0)
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			sheet.Update(i, j, " ")
		}
	}

	sheet.Update(0, 0, title[0])
	sheet.Update(0, 1, title[1])

	for i := 0; i < len(body)/2; i += 1 {
		sheet.Update(i+1, 0, string(body[(2*i)]))
		sheet.Update(i+1, 1, string(body[(2*i)+1]))
	}

	err = sheet.Synchronize()
	if err != nil {
		panic(err.Error())
	}
}
