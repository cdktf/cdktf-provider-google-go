package computesecuritypolicy


type ComputeSecurityPolicyRuleHeaderAction struct {
	// request_headers_to_adds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_security_policy#request_headers_to_adds ComputeSecurityPolicy#request_headers_to_adds}
	RequestHeadersToAdds interface{} `field:"required" json:"requestHeadersToAdds" yaml:"requestHeadersToAdds"`
}

