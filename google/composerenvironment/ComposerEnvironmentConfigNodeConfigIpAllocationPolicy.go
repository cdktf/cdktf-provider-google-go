package composerenvironment


type ComposerEnvironmentConfigNodeConfigIpAllocationPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#cluster_ipv4_cidr_block ComposerEnvironment#cluster_ipv4_cidr_block}.
	ClusterIpv4CidrBlock *string `field:"optional" json:"clusterIpv4CidrBlock" yaml:"clusterIpv4CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#cluster_secondary_range_name ComposerEnvironment#cluster_secondary_range_name}.
	ClusterSecondaryRangeName *string `field:"optional" json:"clusterSecondaryRangeName" yaml:"clusterSecondaryRangeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#services_ipv4_cidr_block ComposerEnvironment#services_ipv4_cidr_block}.
	ServicesIpv4CidrBlock *string `field:"optional" json:"servicesIpv4CidrBlock" yaml:"servicesIpv4CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#services_secondary_range_name ComposerEnvironment#services_secondary_range_name}.
	ServicesSecondaryRangeName *string `field:"optional" json:"servicesSecondaryRangeName" yaml:"servicesSecondaryRangeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#use_ip_aliases ComposerEnvironment#use_ip_aliases}.
	UseIpAliases interface{} `field:"optional" json:"useIpAliases" yaml:"useIpAliases"`
}

