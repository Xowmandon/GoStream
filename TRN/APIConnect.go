package TRN

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)



type trnAPI struct {

	client http.Client
	ID string

}

func NewTRNconnection(url string) *trnAPI {

	connection := new(trnAPI)

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error with .env file")
	}

	envID := os.Getenv("TRN-Api-Key")

	connection.ID = envID

	return connection

}

func (TRN *trnAPI) Get(url string) (resp *http.Response, err error) {

	res, err := http.NewRequest("GET",url,nil)
	if err != nil {
		return nil, err
	}

	return TRN.Do(res)

}

func (TRN *trnAPI) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err

	}

	req.Header.Set("Content-Type", contentType)

	return TRN.Do(req)

}

func (TRN *trnAPI) Do(req *http.Request) (resp *http.Response, err error) {

	req.Header.Add("TRN-Api-Key", TRN.ID)
	return TRN.client.Do(req)

}