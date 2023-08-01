package dataproccluster


type DataprocClusterClusterConfigWorkerConfigAccelerators struct {
	// The number of the accelerator cards of this type exposed to this instance.
	//
	// Often restricted to one of 1, 2, 4, or 8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_cluster#accelerator_count DataprocCluster#accelerator_count}
	AcceleratorCount *float64 `field:"required" json:"acceleratorCount" yaml:"acceleratorCount"`
	// The short name of the accelerator type to expose to this instance. For example, nvidia-tesla-k80.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_cluster#accelerator_type DataprocCluster#accelerator_type}
	AcceleratorType *string `field:"required" json:"acceleratorType" yaml:"acceleratorType"`
}

