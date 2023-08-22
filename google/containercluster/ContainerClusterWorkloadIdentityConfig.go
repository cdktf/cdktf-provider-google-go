package containercluster


type ContainerClusterWorkloadIdentityConfig struct {
	// The workload pool to attach all Kubernetes service accounts to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/container_cluster#workload_pool ContainerCluster#workload_pool}
	WorkloadPool *string `field:"optional" json:"workloadPool" yaml:"workloadPool"`
}

