package nftstorage

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"

	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/gin-gonic/gin"
)

type response struct {
	IpfsHash string `json:"ipfsHash"`
}

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
	var res []response
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
		var apiKey = envconfig.EnvVars.NFT_STORAGE_KEY
		metadataURI, err := uploadToNFTStorage(apiKey, fileData)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error uploading image to NFT.Storage: %s", f.Name))
			return
		}
		res = append(res, response{
			IpfsHash: metadataURI,
		})
	}
	c.JSON(http.StatusOK, res)
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

	type resBody struct {
		Ok    bool `json:"ok"`
		Value struct {
			Cid     string    `json:"cid"`
			Size    int       `json:"size"`
			Created time.Time `json:"created"`
			Type    string    `json:"type"`
			Scope   string    `json:"scope"`
			Pin     struct {
				Cid  string `json:"cid"`
				Name string `json:"name"`
				Meta struct {
				} `json:"meta"`
				Status  string    `json:"status"`
				Created time.Time `json:"created"`
				Size    int       `json:"size"`
			} `json:"pin"`
			Files []struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"files"`
			Deals []struct {
				BatchRootCid   string    `json:"batchRootCid"`
				LastChange     time.Time `json:"lastChange"`
				Miner          string    `json:"miner"`
				Network        string    `json:"network"`
				PieceCid       string    `json:"pieceCid"`
				Status         string    `json:"status"`
				StatusText     string    `json:"statusText"`
				ChainDealID    int       `json:"chainDealID"`
				DealActivation time.Time `json:"dealActivation"`
				DealExpiration time.Time `json:"dealExpiration"`
			} `json:"deals"`
		} `json:"value"`
	}

	var response resBody
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}
	return response.Value.Cid, nil
}
