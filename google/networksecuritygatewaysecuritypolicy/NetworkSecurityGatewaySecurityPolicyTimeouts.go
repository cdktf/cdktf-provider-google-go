package networksecuritygatewaysecuritypolicy


type NetworkSecurityGatewaySecurityPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_security_gateway_security_policy#create NetworkSecurityGatewaySecurityPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_security_gateway_security_policy#delete NetworkSecurityGatewaySecurityPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_security_gateway_security_policy#update NetworkSecurityGatewaySecurityPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

