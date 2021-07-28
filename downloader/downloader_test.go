package downloader

import (
	"github.com/MateoM24/eCatholic/model"
	"os"
	"testing"
	"time"
)

func TestDownloadData(t *testing.T) {
	_, err := downloadData("https://s3.amazonaws.com/ecatholic-hiring/data.csv")
	if err != nil {
		t.Fatalf("Downloading file has failed. Error = %s", err)
	}
}

func TestMapToModel(t *testing.T) {
	testDataFilePath := "../test-data.csv"
	file, err := os.Open(testDataFilePath)
	if err != nil {
		t.Fatalf("Failed reading test data file: %v,\n %v", testDataFilePath, err)
	}
	candidates, err := mapToModel(file)
	if err != nil {
		t.Fatal(err)
	}
	if len(candidates) != 15 {
		t.Fatalf("Expected to read 12 candidates but got %v", len(candidates))
	}
	validate15thCandidate(candidates[14], t)
}

func validate15thCandidate(candidate model.Candidate, t *testing.T) {
	if candidate.Date != "06/15/2019" {
		t.Fatalf("Expected date: 06/15/2019 but got %v", candidate.Date)
	}
	if candidate.Name != "Mike Smith" {
		t.Fatalf("Expected name: Mike Smith but got %v", candidate.Name)
	}
	if candidate.Address != "2483 Farland Avenue" {
		t.Fatalf("Expected address: 2483 Farland Avenue but got %v", candidate.Address)
	}
	if candidate.Address2 != "" {
		t.Fatalf("Expected address2 to be blank but got %v", candidate.Address2)
	}
	if candidate.City != "Warrensburg" {
		t.Fatalf("Expected city: Warrensburg but got %v", candidate.City)
	}
	if candidate.State != "MO" {
		t.Fatalf("Expected state: MO but got %v", candidate.State)
	}
	if candidate.Zipcode != "64093" {
		t.Fatalf("Expected zip code: 64093 but got %v", candidate.Zipcode)
	}
	if candidate.Telephone != "443-323-6215" {
		t.Fatalf("Expected telephone: 443-323-6215 but got %v", candidate.Telephone)
	}
	if candidate.Mobile != "410-726-6477" {
		t.Fatalf("Expected mobile: 410-726-6477 but got %v", candidate.Mobile)
	}
	if candidate.Amount != "$40" {
		t.Fatalf("Expected amount: $40 but got %v", candidate.Amount)
	}
	if candidate.Processor != "Stripe" {
		t.Fatalf("Expected processor: Stripe but got %v", candidate.Processor)
	}
	if candidate.ImportDate == ""  {
		t.Fatalf("Expected import date not to be blank")
	}
}

func TestGetFormattedDate(t *testing.T) {
	location, _ := time.LoadLocation("Local")
	date := time.Date(2021, 7, 28, 10, 10, 10, 10, location)
	expected := "07-28-2021"
	if getFormattedDate(date) != expected {
		t.Fatalf("Failed to format date correctly. Expected: %v, got: %v", expected, getFormattedDate(date))
	}
}


