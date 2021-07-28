package downloader

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/MateoM24/eCatholic/model"
)

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
			Date:      row[0],
			Name:      row[1],
			Address:   row[2],
			Address2:  row[3],
			City:      row[4],
			State: 	   row[5],
			Zipcode:   row[6],
			Telephone: row[7],
			Mobile:    row[8],
			Amount:    row[9],
			Processor: row[10],
			ImportDate: getFormattedDate(time.Now()),
		})
	}
	return candidates, nil
}

func getFormattedDate(date time.Time) string {
	return fmt.Sprintf("%02d-%02d-%04d", date.Month(), date.Day(), date.Year())
}
