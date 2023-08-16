package privatecacertificate


type PrivatecaCertificateConfigX509ConfigNameConstraints struct {
	// Indicates whether or not the name constraints are marked critical.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate#critical PrivatecaCertificate#critical}
	Critical interface{} `field:"required" json:"critical" yaml:"critical"`
	// Contains excluded DNS names.
	//
	// Any DNS name that can be
	// constructed by simply adding zero or more labels to
	// the left-hand side of the name satisfies the name constraint.
	// For example, 'example.com', 'www.example.com', 'www.sub.example.com'
	// would satisfy 'example.com' while 'example1.com' does not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate#excluded_dns_names PrivatecaCertificate#excluded_dns_names}
	ExcludedDnsNames *[]*string `field:"optional" json:"excludedDnsNames" yaml:"excludedDnsNames"`
	// Contains the excluded email addresses.
	//
	// The value can be a particular
	// email address, a hostname to indicate all email addresses on that host or
	// a domain with a leading period (e.g. '.example.com') to indicate
	// all email addresses in that domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate#excluded_email_addresses PrivatecaCertificate#excluded_email_addresses}
	ExcludedEmailAddresses *[]*string `field:"optional" json:"excludedEmailAddresses" yaml:"excludedEmailAddresses"`
	// Contains the excluded IP ranges.
	//
	// For IPv4 addresses, the ranges
	// are expressed using CIDR notation as specified in RFC 4632.
	// For IPv6 addresses, the ranges are expressed in similar encoding as IPv4
	// addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate#excluded_ip_ranges PrivatecaCertificate#excluded_ip_ranges}
	ExcludedIpRanges *[]*string `field:"optional" json:"excludedIpRanges" yaml:"excludedIpRanges"`
	// Contains the excluded URIs that apply to the host part of the name.
	//
	// The value can be a hostname or a domain with a
	// leading period (like '.example.com')
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate#excluded_uris PrivatecaCertificate#excluded_uris}
	ExcludedUris *[]*string `field:"optional" json:"excludedUris" yaml:"excludedUris"`
	// Contains permitted DNS names.
	//
	// Any DNS name that can be
	// constructed by simply adding zero or more labels to
	// the left-hand side of the name satisfies the name constraint.
	// For example, 'example.com', 'www.example.com', 'www.sub.example.com'
	// would satisfy 'example.com' while 'example1.com' does not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate#permitted_dns_names PrivatecaCertificate#permitted_dns_names}
	PermittedDnsNames *[]*string `field:"optional" json:"permittedDnsNames" yaml:"permittedDnsNames"`
	// Contains the permitted email addresses.
	//
	// The value can be a particular
	// email address, a hostname to indicate all email addresses on that host or
	// a domain with a leading period (e.g. '.example.com') to indicate
	// all email addresses in that domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate#permitted_email_addresses PrivatecaCertificate#permitted_email_addresses}
	PermittedEmailAddresses *[]*string `field:"optional" json:"permittedEmailAddresses" yaml:"permittedEmailAddresses"`
	// Contains the permitted IP ranges.
	//
	// For IPv4 addresses, the ranges
	// are expressed using CIDR notation as specified in RFC 4632.
	// For IPv6 addresses, the ranges are expressed in similar encoding as IPv4
	// addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate#permitted_ip_ranges PrivatecaCertificate#permitted_ip_ranges}
	PermittedIpRanges *[]*string `field:"optional" json:"permittedIpRanges" yaml:"permittedIpRanges"`
	// Contains the permitted URIs that apply to the host part of the name.
	//
	// The value can be a hostname or a domain with a
	// leading period (like '.example.com')
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/privateca_certificate#permitted_uris PrivatecaCertificate#permitted_uris}
	PermittedUris *[]*string `field:"optional" json:"permittedUris" yaml:"permittedUris"`
}

