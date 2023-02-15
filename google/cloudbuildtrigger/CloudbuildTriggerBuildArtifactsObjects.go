package cloudbuildtrigger


type CloudbuildTriggerBuildArtifactsObjects struct {
	// Cloud Storage bucket and optional object path, in the form "gs://bucket/path/to/somewhere/".
	//
	// Files in the workspace matching any path pattern will be uploaded to Cloud Storage with
	// this location as a prefix.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#location CloudbuildTrigger#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Path globs used to match files in the build's workspace.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#paths CloudbuildTrigger#paths}
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
}

