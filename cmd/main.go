package main

import (
	"github.com/sonikq/HFlabs-task/internal/handler"
)

func main() {
	url := "https://confluence.hflabs.ru/pages/viewpage.action?pageId=1181220999"
	title, body := handler.Parse(url)

	handler.WriteToSheet(title, body)
}
