package nftstorage

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var apiKey = "YOUR_NFT_STORAGE_API_KEY"

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/nft")
	{
		// g.Use(paseto.PASETO(false))
		g.POST("", handleUpload)
	}
}

func handleUpload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Error retrieving file")
		return
	}
	defer file.Close()

	// Open the zip file
	r, err := zip.NewReader(file, header.Size)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error opening zip file")
		return
	}

	// Process each image file in the zip
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error opening image file: %s", f.Name))
			return
		}
		defer rc.Close()

		// Read the image file contents
		fileData, err := io.ReadAll(rc)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading image file: %s", f.Name))
			return
		}

		// Upload the image to NFT.Storage
		metadataURI, err := uploadToNFTStorage(apiKey, fileData)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error uploading image to NFT.Storage: %s", f.Name))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("Image %s uploaded with metadata URI: %s\n", f.Name, metadataURI))
	}
}
func uploadToNFTStorage(apiKey string, fileData []byte) (string, error) {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	part, err := writer.CreateFormFile("file", filepath.Base("image.jpeg")) // Use a generic filename
	if err != nil {
		return "", err
	}
	part.Write(fileData)
	writer.Close()

	// Send a POST request to NFT.Storage API
	req, err := http.NewRequest("POST", "https://api.nft.storage/upload", buf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Handle the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	// Extract the metadata URI
	metadataURI, ok := response["metadata"].(string)
	if !ok {
		return "", errors.New("Metadata URI not found in response")
	}

	return metadataURI, nil
}
