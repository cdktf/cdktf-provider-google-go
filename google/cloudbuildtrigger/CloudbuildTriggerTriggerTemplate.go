package cloudbuildtrigger


type CloudbuildTriggerTriggerTemplate struct {
	// Name of the branch to build.
	//
	// Exactly one a of branch name, tag, or commit SHA must be provided.
	// This field is a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#branch_name CloudbuildTrigger#branch_name}
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
	// Explicit commit SHA to build. Exactly one of a branch name, tag, or commit SHA must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#commit_sha CloudbuildTrigger#commit_sha}
	CommitSha *string `field:"optional" json:"commitSha" yaml:"commitSha"`
	// Directory, relative to the source root, in which to run the build.
	//
	// This must be a relative path. If a step's dir is specified and
	// is an absolute path, this value is ignored for that step's
	// execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#dir CloudbuildTrigger#dir}
	Dir *string `field:"optional" json:"dir" yaml:"dir"`
	// Only trigger a build if the revision regex does NOT match the revision regex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#invert_regex CloudbuildTrigger#invert_regex}
	InvertRegex interface{} `field:"optional" json:"invertRegex" yaml:"invertRegex"`
	// ID of the project that owns the Cloud Source Repository. If omitted, the project ID requesting the build is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#project_id CloudbuildTrigger#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Name of the Cloud Source Repository. If omitted, the name "default" is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#repo_name CloudbuildTrigger#repo_name}
	RepoName *string `field:"optional" json:"repoName" yaml:"repoName"`
	// Name of the tag to build.
	//
	// Exactly one of a branch name, tag, or commit SHA must be provided.
	// This field is a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#tag_name CloudbuildTrigger#tag_name}
	TagName *string `field:"optional" json:"tagName" yaml:"tagName"`
}

