package crafty

import (
	"fmt"
	"testing"
)

func TestSetURL(t *testing.T) {
	url := "mc.jaydent.uk"
	port := "8443"
	SetURL(url, port)
	fmt.Println(baseURL)
}
