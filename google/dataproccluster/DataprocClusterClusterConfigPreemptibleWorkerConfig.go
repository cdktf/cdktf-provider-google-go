package dataproccluster


type DataprocClusterClusterConfigPreemptibleWorkerConfig struct {
	// disk_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster#disk_config DataprocCluster#disk_config}
	DiskConfig *DataprocClusterClusterConfigPreemptibleWorkerConfigDiskConfig `field:"optional" json:"diskConfig" yaml:"diskConfig"`
	// Specifies the number of preemptible nodes to create. Defaults to 0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster#num_instances DataprocCluster#num_instances}
	NumInstances *float64 `field:"optional" json:"numInstances" yaml:"numInstances"`
	// Specifies the preemptibility of the secondary nodes. Defaults to PREEMPTIBLE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster#preemptibility DataprocCluster#preemptibility}
	Preemptibility *string `field:"optional" json:"preemptibility" yaml:"preemptibility"`
}

