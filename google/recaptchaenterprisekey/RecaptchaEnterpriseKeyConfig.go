package recaptchaenterprisekey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RecaptchaEnterpriseKeyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Human-readable display name of this key. Modifiable by user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/recaptcha_enterprise_key#display_name RecaptchaEnterpriseKey#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// android_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/recaptcha_enterprise_key#android_settings RecaptchaEnterpriseKey#android_settings}
	AndroidSettings *RecaptchaEnterpriseKeyAndroidSettings `field:"optional" json:"androidSettings" yaml:"androidSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/recaptcha_enterprise_key#id RecaptchaEnterpriseKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ios_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/recaptcha_enterprise_key#ios_settings RecaptchaEnterpriseKey#ios_settings}
	IosSettings *RecaptchaEnterpriseKeyIosSettings `field:"optional" json:"iosSettings" yaml:"iosSettings"`
	// See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/recaptcha_enterprise_key#labels RecaptchaEnterpriseKey#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/recaptcha_enterprise_key#project RecaptchaEnterpriseKey#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// testing_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/recaptcha_enterprise_key#testing_options RecaptchaEnterpriseKey#testing_options}
	TestingOptions *RecaptchaEnterpriseKeyTestingOptions `field:"optional" json:"testingOptions" yaml:"testingOptions"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/recaptcha_enterprise_key#timeouts RecaptchaEnterpriseKey#timeouts}
	Timeouts *RecaptchaEnterpriseKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// web_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/recaptcha_enterprise_key#web_settings RecaptchaEnterpriseKey#web_settings}
	WebSettings *RecaptchaEnterpriseKeyWebSettings `field:"optional" json:"webSettings" yaml:"webSettings"`
}

