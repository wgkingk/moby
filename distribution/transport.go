package distribution

import "net/http"

func NewMixTransport(proxy http.RoundTripper, noProxy http.RoundTripper) http.RoundTripper {
	return &MixTransport{
		proxy,
		noProxy,
	}
}

type MixTransport struct {
	proxyTransport   http.RoundTripper
	noProxyTransport http.RoundTripper
}

func (t *MixTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := t.proxyTransport.RoundTrip(req)
	if err != nil || resp.StatusCode > 399 {
		resp, err = t.noProxyTransport.RoundTrip(req)
	}

	return resp, err
}
