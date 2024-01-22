package ipfs

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/NetSepio/gateway/config/envconfig"
)

func UploadToIpfs(osFile io.Reader, fileName string) (*NFTStorageRes, error) {
	// Create a buffer to store the request body
	var buf bytes.Buffer

	// Create a new multipart writer with the buffer
	w := multipart.NewWriter(&buf)

	// Create a new form field
	fw, err := w.CreateFormFile("file", fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to create file %s: %w", fileName, err)
	}

	// Copy the contents of the file to the form field
	if _, err := io.Copy(fw, osFile); err != nil {
		return nil, fmt.Errorf("failed to copy contents to %s: %w", fileName, err)
	}

	// Close the multipart writer to finalize the request
	w.Close()

	// Send the request
	req, err := http.NewRequest("POST", "https://api.nft.storage/upload", &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request for nft.storage: %w", err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", envconfig.EnvVars.NFT_STORAGE_KEY))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request for nft.storage: %w", err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body of request for nft.storage: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to upload using nft.storage, status code: %d, response body: %s", resp.StatusCode, bodyBytes)
	}
	nftRes, err := UnmarshalNFTStorageRes(bodyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body of request for nft.storage: %w", err)
	}
	defer resp.Body.Close()
	return &nftRes, nil
}
