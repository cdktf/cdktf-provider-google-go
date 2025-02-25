// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrundomainmapping


type CloudRunDomainMappingMetadata struct {
	// In Cloud Run the namespace must be equal to either the project ID or project number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/cloud_run_domain_mapping#namespace CloudRunDomainMapping#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Annotations is a key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations
	//
	// **Note**: The Cloud Run API may add additional annotations that were not provided in your config.
	// If terraform plan shows a diff where a server-side annotation is added, you can add it to your config
	// or apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/cloud_run_domain_mapping#annotations CloudRunDomainMapping#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// May match selectors of replication controllers
	// and routes.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/cloud_run_domain_mapping#labels CloudRunDomainMapping#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

