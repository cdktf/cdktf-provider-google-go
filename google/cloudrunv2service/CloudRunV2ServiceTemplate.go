// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2service


type CloudRunV2ServiceTemplate struct {
	// Unstructured key value map that may be set by external tools to store and arbitrary metadata.
	//
	// They are not queryable and should be preserved when modifying objects.
	//
	// Cloud Run API v2 does not support annotations with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected.
	// All system annotations in v1 now have a corresponding field in v2 RevisionTemplate.
	//
	// This field follows Kubernetes annotations' namespacing, limits, and rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#annotations CloudRunV2Service#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// containers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#containers CloudRunV2Service#containers}
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// A reference to a customer managed encryption key (CMEK) to use to encrypt this container image.
	//
	// For more information, go to https://cloud.google.com/run/docs/securing/using-cmek
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#encryption_key CloudRunV2Service#encryption_key}
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The sandbox environment to host this Revision. Possible values: ["EXECUTION_ENVIRONMENT_GEN1", "EXECUTION_ENVIRONMENT_GEN2"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#execution_environment CloudRunV2Service#execution_environment}
	ExecutionEnvironment *string `field:"optional" json:"executionEnvironment" yaml:"executionEnvironment"`
	// True if GPU zonal redundancy is disabled on this revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#gpu_zonal_redundancy_disabled CloudRunV2Service#gpu_zonal_redundancy_disabled}
	GpuZonalRedundancyDisabled interface{} `field:"optional" json:"gpuZonalRedundancyDisabled" yaml:"gpuZonalRedundancyDisabled"`
	// Unstructured key value map that can be used to organize and categorize objects.
	//
	// User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc.
	// For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels.
	//
	// Cloud Run API v2 does not support labels with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected.
	// All system labels in v1 now have a corresponding field in v2 RevisionTemplate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#labels CloudRunV2Service#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Sets the maximum number of requests that each serving instance can receive.
	//
	// If not specified or 0, defaults to 80 when requested CPU >= 1 and defaults to 1 when requested CPU < 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#max_instance_request_concurrency CloudRunV2Service#max_instance_request_concurrency}
	MaxInstanceRequestConcurrency *float64 `field:"optional" json:"maxInstanceRequestConcurrency" yaml:"maxInstanceRequestConcurrency"`
	// node_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#node_selector CloudRunV2Service#node_selector}
	NodeSelector *CloudRunV2ServiceTemplateNodeSelector `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	// The unique name for the revision.
	//
	// If this field is omitted, it will be automatically generated based on the Service name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#revision CloudRunV2Service#revision}
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
	// scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#scaling CloudRunV2Service#scaling}
	Scaling *CloudRunV2ServiceTemplateScaling `field:"optional" json:"scaling" yaml:"scaling"`
	// Email address of the IAM service account associated with the revision of the service.
	//
	// The service account represents the identity of the running revision, and determines what permissions the revision has. If not provided, the revision will use the project's default service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#service_account CloudRunV2Service#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Enables session affinity. For more information, go to https://cloud.google.com/run/docs/configuring/session-affinity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#session_affinity CloudRunV2Service#session_affinity}
	SessionAffinity interface{} `field:"optional" json:"sessionAffinity" yaml:"sessionAffinity"`
	// Max allowed time for an instance to respond to a request.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#timeout CloudRunV2Service#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#volumes CloudRunV2Service#volumes}
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
	// vpc_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/cloud_run_v2_service#vpc_access CloudRunV2Service#vpc_access}
	VpcAccess *CloudRunV2ServiceTemplateVpcAccess `field:"optional" json:"vpcAccess" yaml:"vpcAccess"`
}

