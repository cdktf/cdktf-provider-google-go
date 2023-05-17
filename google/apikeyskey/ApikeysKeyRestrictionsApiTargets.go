package apikeyskey


type ApikeysKeyRestrictionsApiTargets struct {
	// The service for this restriction.
	//
	// It should be the canonical service name, for example: `translate.googleapis.com`. You can use `gcloud services list` to get a list of services that are enabled in the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apikeys_key#service ApikeysKey#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Optional.
	//
	// List of one or more methods that can be called. If empty, all methods for the service are allowed. A wildcard (*) can be used as the last symbol. Valid examples: `google.cloud.translate.v2.TranslateService.GetSupportedLanguage` `TranslateText` `Get*` `translate.googleapis.com.Get*`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apikeys_key#methods ApikeysKey#methods}
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
}

