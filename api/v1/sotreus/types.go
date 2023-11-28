package sotreus

type DeployRequest struct {
	Name   string `json:"name,omitempty"`
	Region string `json:"region,omitempty"`
}
type DeployerCreateRequest struct {
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
		FirewallEndpoint  string `json:"firewall_endpoint"`
		DashboardPassword string `json:"dashboard_password"`
	} `json:"message"`
}

type DeployResponse struct {
	VpnID             string `json:"vpn_id"`
	VpnEndpoint       string `json:"vpn_endpoint"`
	FirewallEndpoint  string `json:"firewall_endpoint"`
	DashboardPassword string `json:"dashboard_password"`
}
