package publisher

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MateoM24/eCatholic/model"
	"io/ioutil"
	"net/http"
)

var Client HttpClient

const XApiKey = "X-API-KEY"
const ContentType = "Content-Type"
const ContentTypeValue = "application/json"

/*ServerConfig stored information about server we publish payment data to*/
type ServerConfig struct {
	Url    string
	ApiKey string
}

func init() {
	Client = http.DefaultClient
}

/*PublishPaymentData uses Http POST method to send payment data into server. Body is send as json.
Use ServerConfig to specify server url and api key which will be send in X-API-KEY header*/
func PublishPaymentData(config ServerConfig, candidates []model.Candidate) (status int, e error) {

	dto := candidatesToDto(candidates)

	byteBody, err := dtoToJson(dto)
	if err != nil {
		return 0, err
	}

	request, err := http.NewRequest("POST", config.Url, bytes.NewBuffer(byteBody))
	if err != nil {
		return 0, err
	}

	request.Header.Set(ContentType, ContentTypeValue)
	request.Header.Set(XApiKey, config.ApiKey)

	response, err := Client.Do(request)
	if err != nil {
		if response != nil {
			return response.StatusCode, err
		}
		return 0, err
	}

	if response.StatusCode != 200 {
		return handleNot200Status(response)
	}

	return response.StatusCode, nil
}

func handleNot200Status(response *http.Response) (status int, e error){
	body := response.Body
	var bodyString string
	if body != nil {
		bodyBytes, e := ioutil.ReadAll(body)
		if e != nil {
			bodyString = string(bodyBytes)
		}
		defer body.Close()
	}
	return response.StatusCode,
		errors.New(fmt.Sprintf("Publishing payment records has failed. Status: %v, additional info: %v",
			response.Status, bodyString))
}

func candidatesToDto(candidates []model.Candidate)  model.PaymentRecordDto{
	return model.PaymentRecordDto{Candidates: candidates}
}

func dtoToJson(dto model.PaymentRecordDto) ([]byte, error) {
	return json.Marshal(dto)
}
