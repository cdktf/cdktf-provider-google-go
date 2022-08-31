// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ServiceNetworkingPeeredDnsDomainTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/service_networking_peered_dns_domain#create ServiceNetworkingPeeredDnsDomain#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/service_networking_peered_dns_domain#delete ServiceNetworkingPeeredDnsDomain#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/service_networking_peered_dns_domain#read ServiceNetworkingPeeredDnsDomain#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

