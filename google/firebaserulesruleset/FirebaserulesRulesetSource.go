package firebaserulesruleset


type FirebaserulesRulesetSource struct {
	// files block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/firebaserules_ruleset#files FirebaserulesRuleset#files}
	Files interface{} `field:"required" json:"files" yaml:"files"`
	// `Language` of the `Source` bundle. If unspecified, the language will default to `FIREBASE_RULES`. Possible values: LANGUAGE_UNSPECIFIED, FIREBASE_RULES, EVENT_FLOW_TRIGGERS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/firebaserules_ruleset#language FirebaserulesRuleset#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
}

