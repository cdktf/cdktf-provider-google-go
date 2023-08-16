package dataproccluster


type DataprocClusterVirtualClusterConfig struct {
	// auxiliary_services_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_cluster#auxiliary_services_config DataprocCluster#auxiliary_services_config}
	AuxiliaryServicesConfig *DataprocClusterVirtualClusterConfigAuxiliaryServicesConfig `field:"optional" json:"auxiliaryServicesConfig" yaml:"auxiliaryServicesConfig"`
	// kubernetes_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_cluster#kubernetes_cluster_config DataprocCluster#kubernetes_cluster_config}
	KubernetesClusterConfig *DataprocClusterVirtualClusterConfigKubernetesClusterConfig `field:"optional" json:"kubernetesClusterConfig" yaml:"kubernetesClusterConfig"`
	// A Cloud Storage bucket used to stage job dependencies, config files, and job driver console output.
	//
	// If you do not specify a staging bucket, Cloud Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's staging bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_cluster#staging_bucket DataprocCluster#staging_bucket}
	StagingBucket *string `field:"optional" json:"stagingBucket" yaml:"stagingBucket"`
}

