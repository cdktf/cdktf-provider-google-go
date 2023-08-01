package privatecacertificateauthority


type PrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltName struct {
	// Contains only valid, fully-qualified host names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate_authority#dns_names PrivatecaCertificateAuthority#dns_names}
	DnsNames *[]*string `field:"optional" json:"dnsNames" yaml:"dnsNames"`
	// Contains only valid RFC 2822 E-mail addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate_authority#email_addresses PrivatecaCertificateAuthority#email_addresses}
	EmailAddresses *[]*string `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
	// Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate_authority#ip_addresses PrivatecaCertificateAuthority#ip_addresses}
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
	// Contains only valid RFC 3986 URIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate_authority#uris PrivatecaCertificateAuthority#uris}
	Uris *[]*string `field:"optional" json:"uris" yaml:"uris"`
}

