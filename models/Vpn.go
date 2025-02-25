package models

import "time"

type Sotreus struct {
	Name          string `gorm:"primary_key" json:"name"`
	WalletAddress string `json:"walletAddress"`
	Region        string `json:"region"`
}

type Erebrus struct {
	UUID          string    `gorm:"primary_key" json:"UUID"`
	Name          string    `json:"name"`
	WalletAddress string    `json:"walletAddress"`
	UserId        string    `json:"userId,omitempty"`
	Region        string    `json:"region"`
	NodeId        string    `json:"nodeId"`
	Domain        string    `json:"domain"`
	CollectionId  string    `json:"collectionId"`
	CreatedAt     time.Time `json:"created_at"`
	Chain         string    `json:"chainName"`
	BlobId        string    `json:"blobId"`
}

//	type Erebrus struct {
//		UUID          string `gorm:"primary_key" json:"UUID"`
//		Name          string `json:"name"`
//		WalletAddress string `json:"walletAddress"`
//		Region        string `json:"region"`
//		CollectionId  string `json:"collectionId"`
//	}
