package privatecacertificateauthority


type PrivatecaCertificateAuthorityConfigSubjectConfigSubject struct {
	// The common name of the distinguished name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_authority#common_name PrivatecaCertificateAuthority#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
	// The organization of the subject.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_authority#organization PrivatecaCertificateAuthority#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The country code of the subject.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_authority#country_code PrivatecaCertificateAuthority#country_code}
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// The locality or city of the subject.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_authority#locality PrivatecaCertificateAuthority#locality}
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// The organizational unit of the subject.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_authority#organizational_unit PrivatecaCertificateAuthority#organizational_unit}
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// The postal code of the subject.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_authority#postal_code PrivatecaCertificateAuthority#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// The province, territory, or regional state of the subject.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_authority#province PrivatecaCertificateAuthority#province}
	Province *string `field:"optional" json:"province" yaml:"province"`
	// The street address of the subject.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_authority#street_address PrivatecaCertificateAuthority#street_address}
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
}

