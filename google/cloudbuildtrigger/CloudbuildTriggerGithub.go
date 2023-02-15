package cloudbuildtrigger


type CloudbuildTriggerGithub struct {
	// Name of the repository. For example: The name for https://github.com/googlecloudplatform/cloud-builders is "cloud-builders".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#name CloudbuildTrigger#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Owner of the repository. For example: The owner for https://github.com/googlecloudplatform/cloud-builders is "googlecloudplatform".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#owner CloudbuildTrigger#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// pull_request block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#pull_request CloudbuildTrigger#pull_request}
	PullRequest *CloudbuildTriggerGithubPullRequest `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// push block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#push CloudbuildTrigger#push}
	Push *CloudbuildTriggerGithubPush `field:"optional" json:"push" yaml:"push"`
}

