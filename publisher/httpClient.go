package publisher

import "net/http"

/*HttpClient is an abstraction for HTTP communication. Let's us use different clients for production and tests.*/
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
