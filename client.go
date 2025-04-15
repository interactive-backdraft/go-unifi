package unifi

type Client struct {
	baseURL  string
	username string
	password string
	site     string
	version  string
}

func NewClient(baseURL, username, password, site, version string) *Client {
	return &Client{baseURL: baseURL,
		username: username,
		password: password,
		site:     site,
		version:  version}
}
