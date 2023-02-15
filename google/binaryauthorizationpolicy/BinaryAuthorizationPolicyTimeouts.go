package binaryauthorizationpolicy


type BinaryAuthorizationPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/binary_authorization_policy#create BinaryAuthorizationPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/binary_authorization_policy#delete BinaryAuthorizationPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/binary_authorization_policy#update BinaryAuthorizationPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

