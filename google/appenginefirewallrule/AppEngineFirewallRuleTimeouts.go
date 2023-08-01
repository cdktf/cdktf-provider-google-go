package appenginefirewallrule


type AppEngineFirewallRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/app_engine_firewall_rule#create AppEngineFirewallRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/app_engine_firewall_rule#delete AppEngineFirewallRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/app_engine_firewall_rule#update AppEngineFirewallRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

