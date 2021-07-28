package publisher

import "net/http"

/*MockHttpClient exists for testing purposes. You can provide PostDoFunc to fake real server logic*/
type MockHttpClient struct {
	PostDoFunc func(req *http.Request) (*http.Response, error)
}

func (c *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return c.PostDoFunc(req)
}
