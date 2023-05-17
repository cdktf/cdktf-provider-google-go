package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedClusterConfig struct {
	// autoscaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#autoscaling_config DataprocWorkflowTemplate#autoscaling_config}
	AutoscalingConfig *DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig `field:"optional" json:"autoscalingConfig" yaml:"autoscalingConfig"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#encryption_config DataprocWorkflowTemplate#encryption_config}
	EncryptionConfig *DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#endpoint_config DataprocWorkflowTemplate#endpoint_config}
	EndpointConfig *DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig `field:"optional" json:"endpointConfig" yaml:"endpointConfig"`
	// gce_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#gce_cluster_config DataprocWorkflowTemplate#gce_cluster_config}
	GceClusterConfig *DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig `field:"optional" json:"gceClusterConfig" yaml:"gceClusterConfig"`
	// initialization_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#initialization_actions DataprocWorkflowTemplate#initialization_actions}
	InitializationActions interface{} `field:"optional" json:"initializationActions" yaml:"initializationActions"`
	// lifecycle_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#lifecycle_config DataprocWorkflowTemplate#lifecycle_config}
	LifecycleConfig *DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig `field:"optional" json:"lifecycleConfig" yaml:"lifecycleConfig"`
	// master_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#master_config DataprocWorkflowTemplate#master_config}
	MasterConfig *DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig `field:"optional" json:"masterConfig" yaml:"masterConfig"`
	// secondary_worker_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#secondary_worker_config DataprocWorkflowTemplate#secondary_worker_config}
	SecondaryWorkerConfig *DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig `field:"optional" json:"secondaryWorkerConfig" yaml:"secondaryWorkerConfig"`
	// security_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#security_config DataprocWorkflowTemplate#security_config}
	SecurityConfig *DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// software_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#software_config DataprocWorkflowTemplate#software_config}
	SoftwareConfig *DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig `field:"optional" json:"softwareConfig" yaml:"softwareConfig"`
	// Optional.
	//
	// A Cloud Storage bucket used to stage job dependencies, config files, and job driver console output. If you do not specify a staging bucket, Cloud Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's staging bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket (see [Dataproc staging bucket](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)). **This field requires a Cloud Storage bucket name, not a URI to a Cloud Storage bucket.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#staging_bucket DataprocWorkflowTemplate#staging_bucket}
	StagingBucket *string `field:"optional" json:"stagingBucket" yaml:"stagingBucket"`
	// Optional.
	//
	// A Cloud Storage bucket used to store ephemeral cluster and jobs data, such as Spark and MapReduce history files. If you do not specify a temp bucket, Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's temp bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket. The default bucket has a TTL of 90 days, but you can use any TTL (or none) if you specify a bucket. **This field requires a Cloud Storage bucket name, not a URI to a Cloud Storage bucket.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#temp_bucket DataprocWorkflowTemplate#temp_bucket}
	TempBucket *string `field:"optional" json:"tempBucket" yaml:"tempBucket"`
	// worker_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#worker_config DataprocWorkflowTemplate#worker_config}
	WorkerConfig *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig `field:"optional" json:"workerConfig" yaml:"workerConfig"`
}

