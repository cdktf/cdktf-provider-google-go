package containercluster


type ContainerClusterWorkloadIdentityConfig struct {
	// The workload pool to attach all Kubernetes service accounts to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#workload_pool ContainerCluster#workload_pool}
	WorkloadPool *string `field:"optional" json:"workloadPool" yaml:"workloadPool"`
}

