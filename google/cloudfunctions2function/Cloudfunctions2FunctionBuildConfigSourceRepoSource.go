package cloudfunctions2function


type Cloudfunctions2FunctionBuildConfigSourceRepoSource struct {
	// Regex matching branches to build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions2_function#branch_name Cloudfunctions2Function#branch_name}
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
	// Regex matching tags to build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions2_function#commit_sha Cloudfunctions2Function#commit_sha}
	CommitSha *string `field:"optional" json:"commitSha" yaml:"commitSha"`
	// Directory, relative to the source root, in which to run the build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions2_function#dir Cloudfunctions2Function#dir}
	Dir *string `field:"optional" json:"dir" yaml:"dir"`
	// Only trigger a build if the revision regex does NOT match the revision regex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions2_function#invert_regex Cloudfunctions2Function#invert_regex}
	InvertRegex interface{} `field:"optional" json:"invertRegex" yaml:"invertRegex"`
	// ID of the project that owns the Cloud Source Repository. If omitted, the project ID requesting the build is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions2_function#project_id Cloudfunctions2Function#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Name of the Cloud Source Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions2_function#repo_name Cloudfunctions2Function#repo_name}
	RepoName *string `field:"optional" json:"repoName" yaml:"repoName"`
	// Regex matching tags to build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions2_function#tag_name Cloudfunctions2Function#tag_name}
	TagName *string `field:"optional" json:"tagName" yaml:"tagName"`
}

