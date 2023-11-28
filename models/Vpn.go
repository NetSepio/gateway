package models

type Sotreus struct {
	Name          string `json:"name,omitempty"`
	WalletAddress string `json:"walletAddress,omitempty"`
	Region        string `json:"region,omitempty"`
}
type AdguardContainerInfo struct {
	ContainerID   string `json:"ContainerID,omitempty"`
	Image         string `json:"Image,omitempty"`
	ContainerName string `json:"ContainerName,omitempty"`
	UIPort        string `json:"UIPort,omitempty"`
	DNSPOrt       string `json:"DNSPOrt,omitempty"`
	SetupPort     string `json:"SetupPort,omitempty"`
}
type SotreusContainerInfo struct {
	ContainerID   string `json:"ContainerID,omitempty"`
	Image         string `json:"Image,omitempty"`
	ContainerName string `json:"ContainerName,omitempty"`
	ApiPort       string `json:"ApiPort,omitempty"`
	VPNPort       string `sjson:"VPNPort,omitempty"`
}
