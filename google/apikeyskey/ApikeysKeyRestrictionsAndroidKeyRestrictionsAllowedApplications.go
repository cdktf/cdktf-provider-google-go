package apikeyskey


type ApikeysKeyRestrictionsAndroidKeyRestrictionsAllowedApplications struct {
	// The package name of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/apikeys_key#package_name ApikeysKey#package_name}
	PackageName *string `field:"required" json:"packageName" yaml:"packageName"`
	// The SHA1 fingerprint of the application.
	//
	// For example, both sha1 formats are acceptable : DA:39:A3:EE:5E:6B:4B:0D:32:55:BF:EF:95:60:18:90:AF:D8:07:09 or DA39A3EE5E6B4B0D3255BFEF95601890AFD80709. Output format is the latter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/apikeys_key#sha1_fingerprint ApikeysKey#sha1_fingerprint}
	Sha1Fingerprint *string `field:"required" json:"sha1Fingerprint" yaml:"sha1Fingerprint"`
}

