package webscrape

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

// CheckDomain if the domain exists, write its content to a file else return error
func CheckDomain(basePath string, domain string) (err error) {
	// Make HTTP request
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, domain, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
	req.Header.Add("Accept", "*/*")
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	parsedURL, err := url.Parse(domain)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create output file
	outFile, err := os.Create(path.Join(basePath, parsedURL.Host))
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Copy data from the response to standard output
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
