package lookerinstance


type LookerInstanceAdminSettings struct {
	// Email domain allowlist for the instance.
	//
	// Define the email domains to which your users can deliver Looker (Google Cloud core) content.
	// Updating this list will restart the instance. Updating the allowed email domains from terraform
	// means the value provided will be considered as the entire list and not an amendment to the
	// existing list of allowed email domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/looker_instance#allowed_email_domains LookerInstance#allowed_email_domains}
	AllowedEmailDomains *[]*string `field:"optional" json:"allowedEmailDomains" yaml:"allowedEmailDomains"`
}

