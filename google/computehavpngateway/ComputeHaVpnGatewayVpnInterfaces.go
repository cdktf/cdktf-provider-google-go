package computehavpngateway


type ComputeHaVpnGatewayVpnInterfaces struct {
	// The numeric ID of this VPN gateway interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_ha_vpn_gateway#id ComputeHaVpnGateway#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// URL of the interconnect attachment resource.
	//
	// When the value
	// of this field is present, the VPN Gateway will be used for
	// IPsec-encrypted Cloud Interconnect; all Egress or Ingress
	// traffic for this VPN Gateway interface will go through the
	// specified interconnect attachment resource.
	//
	// Not currently available publicly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_ha_vpn_gateway#interconnect_attachment ComputeHaVpnGateway#interconnect_attachment}
	InterconnectAttachment *string `field:"optional" json:"interconnectAttachment" yaml:"interconnectAttachment"`
}

