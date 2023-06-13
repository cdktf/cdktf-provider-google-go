package recaptchaenterprisekey


type RecaptchaEnterpriseKeyTestingOptions struct {
	// For challenge-based keys only (CHECKBOX, INVISIBLE), all challenge requests for this site will return nocaptcha if NOCAPTCHA, or an unsolvable challenge if UNSOLVABLE_CHALLENGE.
	//
	// Possible values: TESTING_CHALLENGE_UNSPECIFIED, NOCAPTCHA, UNSOLVABLE_CHALLENGE
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/recaptcha_enterprise_key#testing_challenge RecaptchaEnterpriseKey#testing_challenge}
	TestingChallenge *string `field:"optional" json:"testingChallenge" yaml:"testingChallenge"`
	// All assessments for this Key will return this score.
	//
	// Must be between 0 (likely not legitimate) and 1 (likely legitimate) inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/recaptcha_enterprise_key#testing_score RecaptchaEnterpriseKey#testing_score}
	TestingScore *float64 `field:"optional" json:"testingScore" yaml:"testingScore"`
}

