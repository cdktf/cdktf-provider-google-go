// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2workerpool


type CloudRunV2WorkerPoolTemplate struct {
	// Unstructured key value map that may be set by external tools to store and arbitrary metadata.
	//
	// They are not queryable and should be preserved when modifying objects.
	//
	// Cloud Run API v2 does not support annotations with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected.
	// All system annotations in v1 now have a corresponding field in v2 WorkerPoolRevisionTemplate.
	//
	// This field follows Kubernetes annotations' namespacing, limits, and rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_worker_pool#annotations CloudRunV2WorkerPool#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// containers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_worker_pool#containers CloudRunV2WorkerPool#containers}
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// A reference to a customer managed encryption key (CMEK) to use to encrypt this container image.
	//
	// For more information, go to https://cloud.google.com/run/docs/securing/using-cmek
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_worker_pool#encryption_key CloudRunV2WorkerPool#encryption_key}
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The action to take if the encryption key is revoked. Possible values: ["PREVENT_NEW", "SHUTDOWN"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_worker_pool#encryption_key_revocation_action CloudRunV2WorkerPool#encryption_key_revocation_action}
	EncryptionKeyRevocationAction *string `field:"optional" json:"encryptionKeyRevocationAction" yaml:"encryptionKeyRevocationAction"`
	// If encryptionKeyRevocationAction is SHUTDOWN, the duration before shutting down all instances. The minimum increment is 1 hour.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_worker_pool#encryption_key_shutdown_duration CloudRunV2WorkerPool#encryption_key_shutdown_duration}
	EncryptionKeyShutdownDuration *string `field:"optional" json:"encryptionKeyShutdownDuration" yaml:"encryptionKeyShutdownDuration"`
	// True if GPU zonal redundancy is disabled on this revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_worker_pool#gpu_zonal_redundancy_disabled CloudRunV2WorkerPool#gpu_zonal_redundancy_disabled}
	GpuZonalRedundancyDisabled interface{} `field:"optional" json:"gpuZonalRedundancyDisabled" yaml:"gpuZonalRedundancyDisabled"`
	// Unstructured key value map that can be used to organize and categorize objects.
	//
	// User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc.
	// For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels.
	//
	// Cloud Run API v2 does not support labels with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected.
	// All system labels in v1 now have a corresponding field in v2 WorkerPoolRevisionTemplate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_worker_pool#labels CloudRunV2WorkerPool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// node_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_worker_pool#node_selector CloudRunV2WorkerPool#node_selector}
	NodeSelector *CloudRunV2WorkerPoolTemplateNodeSelector `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	// The unique name for the revision.
	//
	// If this field is omitted, it will be automatically generated based on the WorkerPool name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_worker_pool#revision CloudRunV2WorkerPool#revision}
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
	// Email address of the IAM service account associated with the revision of the WorkerPool.
	//
	// The service account represents the identity of the running revision, and determines what permissions the revision has. If not provided, the revision will use the project's default service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_worker_pool#service_account CloudRunV2WorkerPool#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_worker_pool#volumes CloudRunV2WorkerPool#volumes}
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
	// vpc_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_worker_pool#vpc_access CloudRunV2WorkerPool#vpc_access}
	VpcAccess *CloudRunV2WorkerPoolTemplateVpcAccess `field:"optional" json:"vpcAccess" yaml:"vpcAccess"`
}

