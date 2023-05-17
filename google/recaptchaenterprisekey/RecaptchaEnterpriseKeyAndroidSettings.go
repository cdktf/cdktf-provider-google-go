package recaptchaenterprisekey


type RecaptchaEnterpriseKeyAndroidSettings struct {
	// If set to true, it means allowed_package_names will not be enforced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/recaptcha_enterprise_key#allow_all_package_names RecaptchaEnterpriseKey#allow_all_package_names}
	AllowAllPackageNames interface{} `field:"optional" json:"allowAllPackageNames" yaml:"allowAllPackageNames"`
	// Android package names of apps allowed to use the key. Example: 'com.companyname.appname'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/recaptcha_enterprise_key#allowed_package_names RecaptchaEnterpriseKey#allowed_package_names}
	AllowedPackageNames *[]*string `field:"optional" json:"allowedPackageNames" yaml:"allowedPackageNames"`
}

