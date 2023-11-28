package sotreus

type DeployRequest struct {
	Name   string `json:"name,omitempty"`
	Domain string `json:"endpoint,omitempty"`
	Region string `json:"password,omitempty"`
}
type DeployerCreateRequest struct {
	Endpoint  string `json:"endpoint,omitempty"`
	SotreusID string `json:"sotreusID,omitempty"`
}
type SotreusRequest struct {
	VpnId string `json:"vpnId,omitempty"`
}
type SotreusResponse struct {
	Todo    string `json:"todo"`
	Result  string `json:"result"`
	Message struct {
		VpnID             string `json:"vpn_id"`
		VpnEndpoint       string `json:"vpn_endpoint"`
		VpnAPIPort        int    `json:"vpn_api_port"`
		VpnExternalPort   int    `json:"vpn_external_port"`
		DashboardPassword string `json:"dashboard_password"`
	} `json:"message"`
}
