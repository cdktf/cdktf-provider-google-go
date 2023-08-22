package cloudrunservice


type CloudRunServiceMetadata struct {
	// Annotations is a key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// More
	// info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations
	//
	// *Note**: The Cloud Run API may add additional annotations that were not provided in your config.
	// If terraform plan shows a diff where a server-side annotation is added, you can add it to your config
	// or apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.
	//
	// Annotations with 'run.googleapis.com/' and 'autoscaling.knative.dev' are restricted. Use the following annotation
	// keys to configure features on a Service:
	//
	// - 'run.googleapis.com/binary-authorization-breakglass' sets the [Binary Authorization breakglass](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--breakglass).
	// - 'run.googleapis.com/binary-authorization' sets the [Binary Authorization](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--binary-authorization).
	// - 'run.googleapis.com/client-name' sets the client name calling the Cloud Run API.
	// - 'run.googleapis.com/custom-audiences' sets the [custom audiences](https://cloud.google.com/sdk/gcloud/reference/alpha/run/deploy#--add-custom-audiences)
	// that can be used in the audience field of ID token for authenticated requests.
	// - 'run.googleapis.com/description' sets a user defined description for the Service.
	// - 'run.googleapis.com/ingress' sets the [ingress settings](https://cloud.google.com/sdk/gcloud/reference/run/deploy#--ingress)
	// for the Service. For example, '"run.googleapis.com/ingress" = "all"'.
	// - 'run.googleapis.com/launch-stage' sets the [launch stage](https://cloud.google.com/run/docs/troubleshooting#launch-stage-validation)
	// when a preview feature is used. For example, '"run.googleapis.com/launch-stage": "BETA"'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#annotations CloudRunService#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// May match selectors of replication controllers
	// and routes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#labels CloudRunService#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// In Cloud Run the namespace must be equal to either the project ID or project number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#namespace CloudRunService#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

