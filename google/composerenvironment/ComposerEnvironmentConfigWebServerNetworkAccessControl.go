package composerenvironment


type ComposerEnvironmentConfigWebServerNetworkAccessControl struct {
	// allowed_ip_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#allowed_ip_range ComposerEnvironment#allowed_ip_range}
	AllowedIpRange interface{} `field:"optional" json:"allowedIpRange" yaml:"allowedIpRange"`
}

