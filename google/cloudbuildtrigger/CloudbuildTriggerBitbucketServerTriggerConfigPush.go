package cloudbuildtrigger


type CloudbuildTriggerBitbucketServerTriggerConfigPush struct {
	// Regex of branches to match.  Specify only one of branch or tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#branch CloudbuildTrigger#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// When true, only trigger a build if the revision regex does NOT match the gitRef regex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#invert_regex CloudbuildTrigger#invert_regex}
	InvertRegex interface{} `field:"optional" json:"invertRegex" yaml:"invertRegex"`
	// Regex of tags to match.  Specify only one of branch or tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuild_trigger#tag CloudbuildTrigger#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

