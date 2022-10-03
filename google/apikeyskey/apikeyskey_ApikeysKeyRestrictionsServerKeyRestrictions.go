package apikeyskey


type ApikeysKeyRestrictionsServerKeyRestrictions struct {
	// A list of the caller IP addresses that are allowed to make API calls with this key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apikeys_key#allowed_ips ApikeysKey#allowed_ips}
	AllowedIps *[]*string `field:"required" json:"allowedIps" yaml:"allowedIps"`
}

