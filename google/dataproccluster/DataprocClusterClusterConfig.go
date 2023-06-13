package dataproccluster


type DataprocClusterClusterConfig struct {
	// autoscaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#autoscaling_config DataprocCluster#autoscaling_config}
	AutoscalingConfig *DataprocClusterClusterConfigAutoscalingConfig `field:"optional" json:"autoscalingConfig" yaml:"autoscalingConfig"`
	// dataproc_metric_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#dataproc_metric_config DataprocCluster#dataproc_metric_config}
	DataprocMetricConfig *DataprocClusterClusterConfigDataprocMetricConfig `field:"optional" json:"dataprocMetricConfig" yaml:"dataprocMetricConfig"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#encryption_config DataprocCluster#encryption_config}
	EncryptionConfig *DataprocClusterClusterConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#endpoint_config DataprocCluster#endpoint_config}
	EndpointConfig *DataprocClusterClusterConfigEndpointConfig `field:"optional" json:"endpointConfig" yaml:"endpointConfig"`
	// gce_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#gce_cluster_config DataprocCluster#gce_cluster_config}
	GceClusterConfig *DataprocClusterClusterConfigGceClusterConfig `field:"optional" json:"gceClusterConfig" yaml:"gceClusterConfig"`
	// initialization_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#initialization_action DataprocCluster#initialization_action}
	InitializationAction interface{} `field:"optional" json:"initializationAction" yaml:"initializationAction"`
	// lifecycle_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#lifecycle_config DataprocCluster#lifecycle_config}
	LifecycleConfig *DataprocClusterClusterConfigLifecycleConfig `field:"optional" json:"lifecycleConfig" yaml:"lifecycleConfig"`
	// master_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#master_config DataprocCluster#master_config}
	MasterConfig *DataprocClusterClusterConfigMasterConfig `field:"optional" json:"masterConfig" yaml:"masterConfig"`
	// metastore_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#metastore_config DataprocCluster#metastore_config}
	MetastoreConfig *DataprocClusterClusterConfigMetastoreConfig `field:"optional" json:"metastoreConfig" yaml:"metastoreConfig"`
	// preemptible_worker_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#preemptible_worker_config DataprocCluster#preemptible_worker_config}
	PreemptibleWorkerConfig *DataprocClusterClusterConfigPreemptibleWorkerConfig `field:"optional" json:"preemptibleWorkerConfig" yaml:"preemptibleWorkerConfig"`
	// security_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#security_config DataprocCluster#security_config}
	SecurityConfig *DataprocClusterClusterConfigSecurityConfig `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// software_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#software_config DataprocCluster#software_config}
	SoftwareConfig *DataprocClusterClusterConfigSoftwareConfig `field:"optional" json:"softwareConfig" yaml:"softwareConfig"`
	// The Cloud Storage staging bucket used to stage files, such as Hadoop jars, between client machines and the cluster.
	//
	// Note: If you don't explicitly specify a staging_bucket then GCP will auto create / assign one for you. However, you are not guaranteed an auto generated bucket which is solely dedicated to your cluster; it may be shared with other clusters in the same region/zone also choosing to use the auto generation option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#staging_bucket DataprocCluster#staging_bucket}
	StagingBucket *string `field:"optional" json:"stagingBucket" yaml:"stagingBucket"`
	// The Cloud Storage temp bucket used to store ephemeral cluster and jobs data, such as Spark and MapReduce history files.
	//
	// Note: If you don't explicitly specify a temp_bucket then GCP will auto create / assign one for you.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#temp_bucket DataprocCluster#temp_bucket}
	TempBucket *string `field:"optional" json:"tempBucket" yaml:"tempBucket"`
	// worker_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataproc_cluster#worker_config DataprocCluster#worker_config}
	WorkerConfig *DataprocClusterClusterConfigWorkerConfig `field:"optional" json:"workerConfig" yaml:"workerConfig"`
}

