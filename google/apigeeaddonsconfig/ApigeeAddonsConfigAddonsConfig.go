package apigeeaddonsconfig


type ApigeeAddonsConfigAddonsConfig struct {
	// advanced_api_ops_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/apigee_addons_config#advanced_api_ops_config ApigeeAddonsConfig#advanced_api_ops_config}
	AdvancedApiOpsConfig *ApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig `field:"optional" json:"advancedApiOpsConfig" yaml:"advancedApiOpsConfig"`
	// api_security_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/apigee_addons_config#api_security_config ApigeeAddonsConfig#api_security_config}
	ApiSecurityConfig *ApigeeAddonsConfigAddonsConfigApiSecurityConfig `field:"optional" json:"apiSecurityConfig" yaml:"apiSecurityConfig"`
	// connectors_platform_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/apigee_addons_config#connectors_platform_config ApigeeAddonsConfig#connectors_platform_config}
	ConnectorsPlatformConfig *ApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig `field:"optional" json:"connectorsPlatformConfig" yaml:"connectorsPlatformConfig"`
	// integration_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/apigee_addons_config#integration_config ApigeeAddonsConfig#integration_config}
	IntegrationConfig *ApigeeAddonsConfigAddonsConfigIntegrationConfig `field:"optional" json:"integrationConfig" yaml:"integrationConfig"`
	// monetization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/apigee_addons_config#monetization_config ApigeeAddonsConfig#monetization_config}
	MonetizationConfig *ApigeeAddonsConfigAddonsConfigMonetizationConfig `field:"optional" json:"monetizationConfig" yaml:"monetizationConfig"`
}

