package apikeyskey


type ApikeysKeyRestrictionsBrowserKeyRestrictions struct {
	// A list of regular expressions for the referrer URLs that are allowed to make API calls with this key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apikeys_key#allowed_referrers ApikeysKey#allowed_referrers}
	AllowedReferrers *[]*string `field:"required" json:"allowedReferrers" yaml:"allowedReferrers"`
}

