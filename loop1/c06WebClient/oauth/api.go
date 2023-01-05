package oauth

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetUser(client *http.Client) error {
	url := "https://api.github.com/user"
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Status code from ", url, ":", resp.StatusCode)
	io.Copy(os.Stdout, resp.Body)
	return nil
}
