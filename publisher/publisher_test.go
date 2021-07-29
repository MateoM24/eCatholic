package publisher

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/MateoM24/eCatholic/model"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const responseBody = "Internal server error with additional info"


func TestPublishPaymentData200Response(t *testing.T) {
	config := getTestConfig()

	Client = &MockHttpClient{PostDoFunc: func(req *http.Request) (*http.Response, error) {
		return verifyHeaders(req, config)
	}}

	status, err := PublishPaymentData(config, getTestData())

	if err != nil {
		t.Fatalf("Failed to publish payment information: %v", err)
	}
	if status != 200 {
		t.Fatalf("Expected status 200 but got: %v", status)
	}
}

func TestPublishPaymentDataResponseErrorResponseNoBody(t *testing.T) {
	config := getTestConfig()

	Client = &MockHttpClient{PostDoFunc: func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			Status:        "500 Internal Server Error",
			StatusCode:    500,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Request:       req,
			Header:        make(http.Header, 0),
		}, nil
	}}

	status, err := PublishPaymentData(config, getTestData())

	if err == nil {
		t.Fatal("Expected to fail")
	}

	if status != 500 {
		t.Fatalf("Expected to get status: %v but got: %v", 500, status)
	}

	if err.Error() == "" {
		t.Fatalf("Expected to get error message but didn't")
	}

	if strings.Index(err.Error(), responseBody) != -1 {
		t.Fatalf("Expected not to include text from response body")
	}
}

func TestPublishPaymentDataResponseErrorResponseWithBody(t *testing.T) {
	config := getTestConfig()

	Client = &MockHttpClient{PostDoFunc: func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			Status:        "500 Internal Server Error",
			StatusCode:    500,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Body: ioutil.NopCloser(bytes.NewBufferString(responseBody)),
			ContentLength: int64(len(responseBody)),
			Request:       req,
			Header:        make(http.Header, 0),
		}, nil
	}}

	status, err := PublishPaymentData(config, getTestData())

	if err == nil {
		t.Fatal("Expected to fail")
	}

	if status != 500 {
		t.Fatalf("Expected to get status: %v but got: %v", 500, status)
	}

	if err.Error() == "" {
		t.Fatalf("Expected to get error message but didn't")
	}

	if strings.Index(err.Error(), responseBody) == -1 {
		t.Fatalf("Expected to include text from response body")
	}
}

func verifyHeaders(req *http.Request, config ServerConfig) (*http.Response, error){

	contentType := req.Header.Get(ContentType)
	if contentType != ContentTypeValue {
		return &http.Response{
			StatusCode:       400,
		}, errors.New(fmt.Sprintf("Expected header: %v but got: %v", ContentTypeValue, contentType))
	}

	xApiKey := req.Header.Get(XApiKey)
	if xApiKey != config.ApiKey {
		return &http.Response{
			StatusCode:       400,
		}, errors.New(fmt.Sprintf("Expected header: %v but got: %v", XApiKey, xApiKey))
	}

	return &http.Response{
		StatusCode:       200,
	}, nil
}

func TestMarshaling(t *testing.T)  {
	bytes, err := dtoToJson(model.PaymentRecordDto{Candidates: getTestData()})
	if err != nil {
		t.Fatal(err)
	}
	require.JSONEq(t, getExpectedJson(), string(bytes))
}

func getTestConfig() ServerConfig {
	return ServerConfig{
		Url:    "http://test.server.com",
		ApiKey: "abc",
	}
}

func getTestData() []model.Candidate {
	return []model.Candidate{
		model.Candidate{
			Date:       "01/01/2021",
			Name:       "Joe",
			Address:    "1st Street",
			Address2:   "Flat #7",
			City:       "Oklahoma City",
			State:      "Oklahoma",
			ZipCode:    "777",
			Telephone:  "12334566",
			Mobile:     "865747536",
			Amount:     "$99",
			Processor:  "stripe",
			ImportDate: "27/07/2021",
		},
		model.Candidate{
			Date:       "02/02/2020",
			Name:       "Jenny",
			Address:    "2st Street",
			Address2:   "Flat #8",
			City:       "Los Angeles",
			State:      "California",
			ZipCode:    "3333",
			Telephone:  "12334566",
			Mobile:     "865747536",
			Amount:     "150",
			Processor:  "paypal",
			ImportDate: "14/05/2021",
		},
	}
}

func getExpectedJson() string {
	return `{
  "PaymentRecord": [
    {
      "date": "01/01/2021",
      "name": "Joe",
      "address": "1st Street",
      "address2": "Flat #7",
      "city": "Oklahoma City",
      "state": "Oklahoma",
      "zipCode": "777",
      "telephone": "12334566",
      "mobile": "865747536",
      "amount": "$99",
      "processor": "stripe",
      "importDate": "27/07/2021"
    },
    {
      "date": "02/02/2020",
      "name": "Jenny",
      "address": "2st Street",
      "address2": "Flat #8",
      "city": "Los Angeles",
      "state": "California",
      "zipCode": "3333",
      "telephone": "12334566",
      "mobile": "865747536",
      "amount": "150",
      "processor": "paypal",
      "importDate": "14/05/2021"
    }
  ]
}`
}
