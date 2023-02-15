package pkg

import (
	"log"
	"net/http"

	"astuart.co/goq"
)

type example struct {
	Title []string `goquery:"th"`
	Body  []string `goquery:"tbody tr td"`
}

func Parse(url string) ([]string, []string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var ex example

	err = goq.NewDecoder(res.Body).Decode(&ex)
	if err != nil {
		log.Fatal(err)
	}

	return ex.Title, ex.Body
}
