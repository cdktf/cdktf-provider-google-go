package dataproccluster


type DataprocClusterClusterConfigEndpointConfig struct {
	// The flag to enable http access to specific ports on the cluster from external sources (aka Component Gateway).
	//
	// Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#enable_http_port_access DataprocCluster#enable_http_port_access}
	EnableHttpPortAccess interface{} `field:"required" json:"enableHttpPortAccess" yaml:"enableHttpPortAccess"`
}

