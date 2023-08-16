package cloudbuildtrigger


type CloudbuildTriggerBuildSourceStorageSource struct {
	// Google Cloud Storage bucket containing the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#bucket CloudbuildTrigger#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Google Cloud Storage object containing the source. This object must be a gzipped archive file (.tar.gz) containing source to build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#object CloudbuildTrigger#object}
	Object *string `field:"required" json:"object" yaml:"object"`
	// Google Cloud Storage generation for the object. If the generation is omitted, the latest generation will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#generation CloudbuildTrigger#generation}
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
}

