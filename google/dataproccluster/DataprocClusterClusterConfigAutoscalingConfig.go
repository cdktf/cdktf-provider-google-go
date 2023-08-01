package dataproccluster


type DataprocClusterClusterConfigAutoscalingConfig struct {
	// The autoscaling policy used by the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_cluster#policy_uri DataprocCluster#policy_uri}
	PolicyUri *string `field:"required" json:"policyUri" yaml:"policyUri"`
}

