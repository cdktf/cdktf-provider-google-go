// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline


type DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironment struct {
	// Additional experiment flags for the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#additional_experiments DataPipelinePipeline#additional_experiments}
	AdditionalExperiments *[]*string `field:"optional" json:"additionalExperiments" yaml:"additionalExperiments"`
	// Additional user labels to be specified for the job.
	//
	// Keys and values should follow the restrictions specified in the labeling restrictions page. An object containing a list of key/value pairs.
	// 'Example: { "name": "wrench", "mass": "1kg", "count": "3" }.'
	// 'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#additional_user_labels DataPipelinePipeline#additional_user_labels}
	AdditionalUserLabels *map[string]*string `field:"optional" json:"additionalUserLabels" yaml:"additionalUserLabels"`
	// Whether to bypass the safety checks for the job's temporary directory. Use with caution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#bypass_temp_dir_validation DataPipelinePipeline#bypass_temp_dir_validation}
	BypassTempDirValidation interface{} `field:"optional" json:"bypassTempDirValidation" yaml:"bypassTempDirValidation"`
	// Whether to enable Streaming Engine for the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#enable_streaming_engine DataPipelinePipeline#enable_streaming_engine}
	EnableStreamingEngine interface{} `field:"optional" json:"enableStreamingEngine" yaml:"enableStreamingEngine"`
	// Configuration for VM IPs. https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#WorkerIPAddressConfiguration Possible values: ["WORKER_IP_UNSPECIFIED", "WORKER_IP_PUBLIC", "WORKER_IP_PRIVATE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#ip_configuration DataPipelinePipeline#ip_configuration}
	IpConfiguration *string `field:"optional" json:"ipConfiguration" yaml:"ipConfiguration"`
	// 'Name for the Cloud KMS key for the job. The key format is: projects//locations//keyRings//cryptoKeys/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#kms_key_name DataPipelinePipeline#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// The machine type to use for the job. Defaults to the value from the template if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#machine_type DataPipelinePipeline#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The maximum number of Compute Engine instances to be made available to your pipeline during execution, from 1 to 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#max_workers DataPipelinePipeline#max_workers}
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// Network to which VMs will be assigned. If empty or unspecified, the service will use the network "default".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#network DataPipelinePipeline#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The initial number of Compute Engine instances for the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#num_workers DataPipelinePipeline#num_workers}
	NumWorkers *float64 `field:"optional" json:"numWorkers" yaml:"numWorkers"`
	// The email address of the service account to run the job as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#service_account_email DataPipelinePipeline#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// Subnetwork to which VMs will be assigned, if desired.
	//
	// You can specify a subnetwork using either a complete URL or an abbreviated path. Expected to be of the form "https://www.googleapis.com/compute/v1/projects/HOST_PROJECT_ID/regions/REGION/subnetworks/SUBNETWORK" or "regions/REGION/subnetworks/SUBNETWORK". If the subnetwork is located in a Shared VPC network, you must use the complete URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#subnetwork DataPipelinePipeline#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// The Cloud Storage path to use for temporary files. Must be a valid Cloud Storage URL, beginning with gs://.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#temp_location DataPipelinePipeline#temp_location}
	TempLocation *string `field:"optional" json:"tempLocation" yaml:"tempLocation"`
	// The Compute Engine region (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in which worker processing should occur, e.g. "us-west1". Mutually exclusive with workerZone. If neither workerRegion nor workerZone is specified, default to the control plane's region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#worker_region DataPipelinePipeline#worker_region}
	WorkerRegion *string `field:"optional" json:"workerRegion" yaml:"workerRegion"`
	// The Compute Engine zone (https://cloud.google.com/compute/docs/regions-zones/regions-zones) in which worker processing should occur, e.g. "us-west1-a". Mutually exclusive with workerRegion. If neither workerRegion nor workerZone is specified, a zone in the control plane's region is chosen based on available capacity. If both workerZone and zone are set, workerZone takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#worker_zone DataPipelinePipeline#worker_zone}
	WorkerZone *string `field:"optional" json:"workerZone" yaml:"workerZone"`
	// The Compute Engine availability zone for launching worker instances to run your pipeline.
	//
	// In the future, workerZone will take precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.3.0/docs/resources/data_pipeline_pipeline#zone DataPipelinePipeline#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

