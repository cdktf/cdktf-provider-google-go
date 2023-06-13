package computeexternalvpngateway


type ComputeExternalVpnGatewayInterface struct {
	// The numeric ID for this interface.
	//
	// Allowed values are based on the redundancy type
	// of this external VPN gateway
	// '0 - SINGLE_IP_INTERNALLY_REDUNDANT'
	// '0, 1 - TWO_IPS_REDUNDANCY'
	// '0, 1, 2, 3 - FOUR_IPS_REDUNDANCY'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_external_vpn_gateway#id ComputeExternalVpnGateway#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// IP address of the interface in the external VPN gateway.
	//
	// Only IPv4 is supported. This IP address can be either from
	// your on-premise gateway or another Cloud provider's VPN gateway,
	// it cannot be an IP address from Google Compute Engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_external_vpn_gateway#ip_address ComputeExternalVpnGateway#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}

