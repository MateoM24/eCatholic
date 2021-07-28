package eCatholic

import (
	"github.com/MateoM24/eCatholic/downloader"
	"github.com/MateoM24/eCatholic/publisher"
)

func ProcessPaymentData(sourceFileUrl, targetUrl, apiKey string) (httpStatus int, e error) {
	candidates, err := downloader.FetchCandidates(sourceFileUrl)
	if err != nil {
		return 0, err
	}

	config := publisher.ServerConfig{
		Url:    targetUrl,
		ApiKey: apiKey,
	}

	status, err := publisher.PublishPaymentData(config, candidates)
	if err != nil {
		return 0, err
	}

	return status, nil
}
