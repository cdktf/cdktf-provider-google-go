package computesecuritypolicy


type ComputeSecurityPolicyAdaptiveProtectionConfig struct {
	// layer_7_ddos_defense_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_security_policy#layer_7_ddos_defense_config ComputeSecurityPolicy#layer_7_ddos_defense_config}
	Layer7DdosDefenseConfig *ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig `field:"optional" json:"layer7DdosDefenseConfig" yaml:"layer7DdosDefenseConfig"`
}

