package cloudbuildtrigger


type CloudbuildTriggerSourceToBuild struct {
	// The branch or tag to use. Must start with "refs/" (required).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/cloudbuild_trigger#ref CloudbuildTrigger#ref}
	Ref *string `field:"required" json:"ref" yaml:"ref"`
	// The type of the repo, since it may not be explicit from the repo field (e.g from a URL). Values can be UNKNOWN, CLOUD_SOURCE_REPOSITORIES, GITHUB, BITBUCKET_SERVER Possible values: ["UNKNOWN", "CLOUD_SOURCE_REPOSITORIES", "GITHUB", "BITBUCKET_SERVER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/cloudbuild_trigger#repo_type CloudbuildTrigger#repo_type}
	RepoType *string `field:"required" json:"repoType" yaml:"repoType"`
	// The full resource name of the github enterprise config. Format: projects/{project}/locations/{location}/githubEnterpriseConfigs/{id}. projects/{project}/githubEnterpriseConfigs/{id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/cloudbuild_trigger#github_enterprise_config CloudbuildTrigger#github_enterprise_config}
	GithubEnterpriseConfig *string `field:"optional" json:"githubEnterpriseConfig" yaml:"githubEnterpriseConfig"`
	// The URI of the repo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/cloudbuild_trigger#uri CloudbuildTrigger#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

