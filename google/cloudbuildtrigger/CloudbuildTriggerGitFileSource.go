package cloudbuildtrigger


type CloudbuildTriggerGitFileSource struct {
	// The path of the file, with the repo root as the root of the path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#path CloudbuildTrigger#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The type of the repo, since it may not be explicit from the repo field (e.g from a URL). Values can be UNKNOWN, CLOUD_SOURCE_REPOSITORIES, GITHUB, BITBUCKET_SERVER Possible values: ["UNKNOWN", "CLOUD_SOURCE_REPOSITORIES", "GITHUB", "BITBUCKET_SERVER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#repo_type CloudbuildTrigger#repo_type}
	RepoType *string `field:"required" json:"repoType" yaml:"repoType"`
	// The full resource name of the bitbucket server config. Format: projects/{project}/locations/{location}/bitbucketServerConfigs/{id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#bitbucket_server_config CloudbuildTrigger#bitbucket_server_config}
	BitbucketServerConfig *string `field:"optional" json:"bitbucketServerConfig" yaml:"bitbucketServerConfig"`
	// The full resource name of the github enterprise config. Format: projects/{project}/locations/{location}/githubEnterpriseConfigs/{id}. projects/{project}/githubEnterpriseConfigs/{id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#github_enterprise_config CloudbuildTrigger#github_enterprise_config}
	GithubEnterpriseConfig *string `field:"optional" json:"githubEnterpriseConfig" yaml:"githubEnterpriseConfig"`
	// The fully qualified resource name of the Repo API repository.
	//
	// The fully qualified resource name of the Repo API repository.
	// If unspecified, the repo from which the trigger invocation originated is assumed to be the repo from which to read the specified path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#repository CloudbuildTrigger#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// The branch, tag, arbitrary ref, or SHA version of the repo to use when resolving the filename (optional).
	//
	// This field respects the same syntax/resolution as described here: https://git-scm.com/docs/gitrevisions
	// If unspecified, the revision from which the trigger invocation originated is assumed to be the revision from which to read the specified path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#revision CloudbuildTrigger#revision}
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
	// The URI of the repo (optional).
	//
	// If unspecified, the repo from which the trigger
	// invocation originated is assumed to be the repo from which to read the specified path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#uri CloudbuildTrigger#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

