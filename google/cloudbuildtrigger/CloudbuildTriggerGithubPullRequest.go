package cloudbuildtrigger


type CloudbuildTriggerGithubPullRequest struct {
	// Regex of branches to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#branch CloudbuildTrigger#branch}
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// Whether to block builds on a "/gcbrun" comment from a repository owner or collaborator. Possible values: ["COMMENTS_DISABLED", "COMMENTS_ENABLED", "COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#comment_control CloudbuildTrigger#comment_control}
	CommentControl *string `field:"optional" json:"commentControl" yaml:"commentControl"`
	// If true, branches that do NOT match the git_ref will trigger a build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#invert_regex CloudbuildTrigger#invert_regex}
	InvertRegex interface{} `field:"optional" json:"invertRegex" yaml:"invertRegex"`
}

