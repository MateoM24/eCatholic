package downloader

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/MateoM24/eCatholic/model"
)

const MISSING = "missing"

func FetchCandidates(url string) ([]model.Candidate, error) {
	data, err := downloadData(url)
	if err != nil {
		return []model.Candidate{}, err
	}
	candidates, err := mapToModel(data)
	if err != nil {
		return []model.Candidate{}, err
	}
	return candidates, nil
}

func downloadData(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, err
	}
	return resp.Body, nil
}

func mapToModel(reader io.Reader) ([]model.Candidate, error) {
	allRows, err := csv.NewReader(reader).ReadAll()
	if err != nil {
		return []model.Candidate{}, err
	}
	var candidates []model.Candidate
	for i, row := range allRows {
		if i == 0 {
			continue
		}
		candidates = append(candidates,
			model.Candidate{
				Date:       replaceIfBlank(row[0]),
				Name:       replaceIfBlank(row[1]),
				Address:    replaceIfBlank(row[2]),
				Address2:   replaceIfBlank(row[3]),
				City:       replaceIfBlank(row[4]),
				State:      replaceIfBlank(row[5]),
				Zipcode:    replaceIfBlank(row[6]),
				Telephone:  replaceIfBlank(row[7]),
				Mobile:     replaceIfBlank(row[8]),
				Amount:     replaceIfBlank(row[9]),
				Processor:  replaceIfBlank(row[10]),
				ImportDate: getFormattedDate(time.Now()),
			})
	}
	return candidates, nil
}

func removeDuplicates(candidates []model.Candidate) []model.Candidate {
	var uniqueCandidates []model.Candidate
candidatesLoop:
	for _, candidate := range candidates {
		for _, uniqueCandidate := range uniqueCandidates {
			if uniqueCandidate.Equals(candidate) {
				continue candidatesLoop
			}
		}
		uniqueCandidates = append(uniqueCandidates, candidate)
	}
	return uniqueCandidates
}

func replaceIfBlank(value string) string {
	if value == "" {
		return MISSING
	}
	return value
}

func getFormattedDate(date time.Time) string {
	return fmt.Sprintf("%02d-%02d-%04d", date.Month(), date.Day(), date.Year())
}
