package cloudbuildbitbucketserverconfig


type CloudbuildBitbucketServerConfigConnectedRepositories struct {
	// Identifier for the project storing the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_bitbucket_server_config#project_key CloudbuildBitbucketServerConfig#project_key}
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// Identifier for the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuild_bitbucket_server_config#repo_slug CloudbuildBitbucketServerConfig#repo_slug}
	RepoSlug *string `field:"required" json:"repoSlug" yaml:"repoSlug"`
}

