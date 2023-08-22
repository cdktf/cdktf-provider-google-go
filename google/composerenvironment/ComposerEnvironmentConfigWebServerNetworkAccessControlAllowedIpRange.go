package composerenvironment


type ComposerEnvironmentConfigWebServerNetworkAccessControlAllowedIpRange struct {
	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	//
	// Examples: 192.168.1.1 or 192.168.0.0/16 or 2001:db8::/32 or 2001:0db8:0000:0042:0000:8a2e:0370:7334. IP range prefixes should be properly truncated. For example, 1.2.3.4/24 should be truncated to 1.2.3.0/24. Similarly, for IPv6, 2001:db8::1/32 should be truncated to 2001:db8::/32.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/composer_environment#value ComposerEnvironment#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// A description of this ip range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/composer_environment#description ComposerEnvironment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

