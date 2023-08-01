package dnsresponsepolicyrule


type DnsResponsePolicyRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_response_policy_rule#create DnsResponsePolicyRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_response_policy_rule#delete DnsResponsePolicyRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_response_policy_rule#update DnsResponsePolicyRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

