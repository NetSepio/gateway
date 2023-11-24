package sotreus

type SotreusDeployBody struct {
	OrgID    string `json:"OrgID,omitempty"`
	OrgName  string `json:"OrgName,omitempty"`
	Password string `pjson:"Password,omitempty"`
}

type SotreusRequest struct {
	Uuid string `json:"uuid,omitempty"`
}
type SotreusResponse struct {
	Status  int64               `json:"status,omitempty"`
	Sucess  bool                `json:"sucess,omitempty"`
	Message string              `json:"message,omitempty"`
	Error   string              `json:"error,omitempty"`
	Sotreus *ServiceInfoSotreus `json:"Sotreus,omitempty"`
}
type ServiceInfoSotreus struct {
	Name      string                `json:"Name,omitempty"`
	Type      string                `json:"Type,omitempty"`
	Uuid      string                `json:"Uuid,omitempty"`
	Category  string                `json:"Category,omitempty"`
	Status    string                `json:"Status,omitempty"`
	CreatedAt int64                 `json:"createdAt,omitempty"`
	UpdatedAt int64                 `json:"updatedAt,omitempty"`
	DeletedAt int64                 `json:"deletedAt,omitempty"`
	Sotreus   *SotreusContainerInfo `json:"Sotreus,omitempty"`
	Adguard   *AdguardContainerInfo `json:"Adguard,omitempty"`
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
