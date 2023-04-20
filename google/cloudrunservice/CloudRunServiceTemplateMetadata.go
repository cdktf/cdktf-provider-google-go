package cloudrunservice


type CloudRunServiceTemplateMetadata struct {
	// Annotations is a key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// *Note**: The Cloud Run API may add additional annotations that were not provided in your config.
	// If terraform plan shows a diff where a server-side annotation is added, you can add it to your config
	// or apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.1/docs/resources/cloud_run_service#annotations CloudRunService#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.1/docs/resources/cloud_run_service#labels CloudRunService#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name must be unique within a Google Cloud project and region.
	//
	// Is required when creating resources. Name is primarily intended
	// for creation idempotence and configuration definition. Cannot be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.1/docs/resources/cloud_run_service#name CloudRunService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// In Cloud Run the namespace must be equal to either the project ID or project number.
	//
	// It will default to the resource's project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.1/docs/resources/cloud_run_service#namespace CloudRunService#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

