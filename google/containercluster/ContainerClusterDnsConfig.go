package containercluster


type ContainerClusterDnsConfig struct {
	// Which in-cluster DNS provider should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#cluster_dns ContainerCluster#cluster_dns}
	ClusterDns *string `field:"optional" json:"clusterDns" yaml:"clusterDns"`
	// The suffix used for all cluster service records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#cluster_dns_domain ContainerCluster#cluster_dns_domain}
	ClusterDnsDomain *string `field:"optional" json:"clusterDnsDomain" yaml:"clusterDnsDomain"`
	// The scope of access to cluster DNS records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#cluster_dns_scope ContainerCluster#cluster_dns_scope}
	ClusterDnsScope *string `field:"optional" json:"clusterDnsScope" yaml:"clusterDnsScope"`
}

