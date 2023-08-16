package firebaserulesruleset


type FirebaserulesRulesetSourceFiles struct {
	// Textual Content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/firebaserules_ruleset#content FirebaserulesRuleset#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// File name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/firebaserules_ruleset#name FirebaserulesRuleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Fingerprint (e.g. github sha) associated with the `File`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/firebaserules_ruleset#fingerprint FirebaserulesRuleset#fingerprint}
	Fingerprint *string `field:"optional" json:"fingerprint" yaml:"fingerprint"`
}

