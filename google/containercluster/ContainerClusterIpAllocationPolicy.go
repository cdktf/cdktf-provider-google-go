package containercluster


type ContainerClusterIpAllocationPolicy struct {
	// The IP address range for the cluster pod IPs.
	//
	// Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#cluster_ipv4_cidr_block ContainerCluster#cluster_ipv4_cidr_block}
	ClusterIpv4CidrBlock *string `field:"optional" json:"clusterIpv4CidrBlock" yaml:"clusterIpv4CidrBlock"`
	// The name of the existing secondary range in the cluster's subnetwork to use for pod IP addresses.
	//
	// Alternatively, cluster_ipv4_cidr_block can be used to automatically create a GKE-managed one.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#cluster_secondary_range_name ContainerCluster#cluster_secondary_range_name}
	ClusterSecondaryRangeName *string `field:"optional" json:"clusterSecondaryRangeName" yaml:"clusterSecondaryRangeName"`
	// The IP address range of the services IPs in this cluster.
	//
	// Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#services_ipv4_cidr_block ContainerCluster#services_ipv4_cidr_block}
	ServicesIpv4CidrBlock *string `field:"optional" json:"servicesIpv4CidrBlock" yaml:"servicesIpv4CidrBlock"`
	// The name of the existing secondary range in the cluster's subnetwork to use for service ClusterIPs.
	//
	// Alternatively, services_ipv4_cidr_block can be used to automatically create a GKE-managed one.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#services_secondary_range_name ContainerCluster#services_secondary_range_name}
	ServicesSecondaryRangeName *string `field:"optional" json:"servicesSecondaryRangeName" yaml:"servicesSecondaryRangeName"`
}

