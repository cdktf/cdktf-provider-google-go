package dnspolicy


type DnsPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_policy#create DnsPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_policy#delete DnsPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_policy#update DnsPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

