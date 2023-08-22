package cloudbuildtrigger


type CloudbuildTriggerRepositoryEventConfig struct {
	// pull_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#pull_request CloudbuildTrigger#pull_request}
	PullRequest *CloudbuildTriggerRepositoryEventConfigPullRequest `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// push block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#push CloudbuildTrigger#push}
	Push *CloudbuildTriggerRepositoryEventConfigPush `field:"optional" json:"push" yaml:"push"`
	// The resource name of the Repo API resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#repository CloudbuildTrigger#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}

