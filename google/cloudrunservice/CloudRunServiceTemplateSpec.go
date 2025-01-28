// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunservice


type CloudRunServiceTemplateSpec struct {
	// ContainerConcurrency specifies the maximum allowed in-flight (concurrent) requests per container of the Revision.
	//
	// If not specified or 0, defaults to 80 when
	// requested CPU >= 1 and defaults to 1 when requested CPU < 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/cloud_run_service#container_concurrency CloudRunService#container_concurrency}
	ContainerConcurrency *float64 `field:"optional" json:"containerConcurrency" yaml:"containerConcurrency"`
	// containers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/cloud_run_service#containers CloudRunService#containers}
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// Email address of the IAM service account associated with the revision of the service.
	//
	// The service account represents the identity of the running revision,
	// and determines what permissions the revision has. If not provided, the revision
	// will use the project's default service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/cloud_run_service#service_account_name CloudRunService#service_account_name}
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// TimeoutSeconds holds the max duration the instance is allowed for responding to a request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/cloud_run_service#timeout_seconds CloudRunService#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/cloud_run_service#volumes CloudRunService#volumes}
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

