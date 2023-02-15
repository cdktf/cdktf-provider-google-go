package apikeyskey


type ApikeysKeyRestrictionsIosKeyRestrictions struct {
	// A list of bundle IDs that are allowed when making API calls with this key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apikeys_key#allowed_bundle_ids ApikeysKey#allowed_bundle_ids}
	AllowedBundleIds *[]*string `field:"required" json:"allowedBundleIds" yaml:"allowedBundleIds"`
}

