// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package recaptchaenterprisekey


type RecaptchaEnterpriseKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/recaptcha_enterprise_key#create RecaptchaEnterpriseKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/recaptcha_enterprise_key#delete RecaptchaEnterpriseKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/recaptcha_enterprise_key#update RecaptchaEnterpriseKey#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

