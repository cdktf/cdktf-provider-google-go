package apikeyskey


type ApikeysKeyRestrictionsIosKeyRestrictions struct {
	// A list of bundle IDs that are allowed when making API calls with this key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apikeys_key#allowed_bundle_ids ApikeysKey#allowed_bundle_ids}
	AllowedBundleIds *[]*string `field:"required" json:"allowedBundleIds" yaml:"allowedBundleIds"`
}

