// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type RecaptchaEnterpriseKeyIosSettings struct {
	// If set to true, it means allowed_bundle_ids will not be enforced.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/recaptcha_enterprise_key#allow_all_bundle_ids RecaptchaEnterpriseKey#allow_all_bundle_ids}
	AllowAllBundleIds interface{} `field:"optional" json:"allowAllBundleIds" yaml:"allowAllBundleIds"`
	// iOS bundle ids of apps allowed to use the key. Example: 'com.companyname.productname.appname'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/recaptcha_enterprise_key#allowed_bundle_ids RecaptchaEnterpriseKey#allowed_bundle_ids}
	AllowedBundleIds *[]*string `field:"optional" json:"allowedBundleIds" yaml:"allowedBundleIds"`
}

