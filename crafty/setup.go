package crafty

import "fmt"

var (
	baseURL string
)

func SetURL(url string, port string) error {
	if baseURL != "" {
		return fmt.Errorf("URL already set to %s", baseURL)
	}
	baseURL = fmt.Sprintf("https://%v:%v/api/v2/", url, port)
	return nil
}

func CheckURL() error {
	if baseURL == "" {
		return fmt.Errorf("URL not set")
	}
	return nil
}
