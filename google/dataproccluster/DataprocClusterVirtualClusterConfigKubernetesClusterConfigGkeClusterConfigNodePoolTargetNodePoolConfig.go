package dataproccluster


type DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig struct {
	// The list of Compute Engine zones where node pool nodes associated with a Dataproc on GKE virtual cluster will be located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#locations DataprocCluster#locations}
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
	// autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#autoscaling DataprocCluster#autoscaling}
	Autoscaling *DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling `field:"optional" json:"autoscaling" yaml:"autoscaling"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#config DataprocCluster#config}
	Config *DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig `field:"optional" json:"config" yaml:"config"`
}

