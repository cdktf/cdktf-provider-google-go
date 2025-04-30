// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunservice


type CloudRunServiceTemplateMetadata struct {
	// Annotations is a key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations
	//
	// **Note**: The Cloud Run API may add additional annotations that were not provided in your config.
	// If terraform plan shows a diff where a server-side annotation is added, you can add it to your config
	// or apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.
	//
	// Annotations with 'run.googleapis.com/' and 'autoscaling.knative.dev' are restricted. Use the following annotation
	// keys to configure features on a Revision template:
	//
	// - 'autoscaling.knative.dev/maxScale' sets the [maximum number of container
	//   instances](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--max-instances) of the Revision to run.
	// - 'autoscaling.knative.dev/minScale' sets the [minimum number of container
	//   instances](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--min-instances) of the Revision to run.
	// - 'run.googleapis.com/client-name' sets the client name calling the Cloud Run API.
	// - 'run.googleapis.com/cloudsql-instances' sets the [Cloud SQL
	//   instances](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--add-cloudsql-instances) the Revision connects to.
	// - 'run.googleapis.com/cpu-throttling' sets whether to throttle the CPU when the container is not actively serving
	//   requests. See https://cloud.google.com/sdk/gcloud/reference/run/deploy#--[no-]cpu-throttling.
	// - 'run.googleapis.com/encryption-key-shutdown-hours' sets the number of hours to wait before an automatic shutdown
	//   server after CMEK key revocation is detected.
	// - 'run.googleapis.com/encryption-key' sets the [CMEK key](https://cloud.google.com/run/docs/securing/using-cmek)
	//   reference to encrypt the container with.
	// - 'run.googleapis.com/execution-environment' sets the [execution
	//   environment](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--execution-environment)
	//   where the application will run.
	// - 'run.googleapis.com/post-key-revocation-action-type' sets the
	//   [action type](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--post-key-revocation-action-type)
	//   after CMEK key revocation.
	// - 'run.googleapis.com/secrets' sets a list of key-value pairs to set as
	//   [secrets](https://cloud.google.com/run/docs/configuring/secrets#yaml).
	// - 'run.googleapis.com/sessionAffinity' sets whether to enable
	//   [session affinity](https://cloud.google.com/sdk/gcloud/reference/beta/run/deploy#--[no-]session-affinity)
	//   for connections to the Revision.
	// - 'run.googleapis.com/startup-cpu-boost' sets whether to allocate extra CPU to containers on startup.
	//   See https://cloud.google.com/sdk/gcloud/reference/run/deploy#--[no-]cpu-boost.
	// - 'run.googleapis.com/network-interfaces' sets [Direct VPC egress](https://cloud.google.com/run/docs/configuring/vpc-direct-vpc#yaml)
	//   for the Revision.
	// - 'run.googleapis.com/vpc-access-connector' sets a [VPC connector](https://cloud.google.com/run/docs/configuring/connecting-vpc#terraform_1)
	//   for the Revision.
	// - 'run.googleapis.com/vpc-access-egress' sets the outbound traffic to send through the VPC connector for this resource.
	//   See https://cloud.google.com/sdk/gcloud/reference/run/deploy#--vpc-egress.
	// - 'run.googleapis.com/gpu-zonal-redundancy-disabled' sets
	//   [GPU zonal redundancy](https://cloud.google.com/run/docs/configuring/services/gpu-zonal-redundancy) for the Revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/cloud_run_service#annotations CloudRunService#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/cloud_run_service#labels CloudRunService#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name must be unique within a Google Cloud project and region.
	//
	// Is required when creating resources. Name is primarily intended
	// for creation idempotence and configuration definition. Cannot be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/cloud_run_service#name CloudRunService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// In Cloud Run the namespace must be equal to either the project ID or project number.
	//
	// It will default to the resource's project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/cloud_run_service#namespace CloudRunService#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

