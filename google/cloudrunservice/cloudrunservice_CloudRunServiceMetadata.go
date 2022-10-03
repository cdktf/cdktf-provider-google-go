package cloudrunservice


type CloudRunServiceMetadata struct {
	// Annotations is a key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// More
	// info: http://kubernetes.io/docs/user-guide/annotations
	//
	// *Note**: The Cloud Run API may add additional annotations that were not provided in your config.
	// If terraform plan shows a diff where a server-side annotation is added, you can add it to your config
	// or apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.
	//
	// Cloud Run (fully managed) uses the following annotation keys to configure features on a Service:
	//
	// - 'run.googleapis.com/ingress' sets the [ingress settings](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--ingress)
	// for the Service. For example, '"run.googleapis.com/ingress" = "all"'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#annotations CloudRunService#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// May match selectors of replication controllers
	// and routes.
	// More info: http://kubernetes.io/docs/user-guide/labels
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#labels CloudRunService#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// In Cloud Run the namespace must be equal to either the project ID or project number.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_service#namespace CloudRunService#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

