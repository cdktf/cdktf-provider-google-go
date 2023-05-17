package apigeekeystoresaliaseskeycertfile


type ApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfo struct {
	// X.509 basic constraints extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apigee_keystores_aliases_key_cert_file#basic_constraints ApigeeKeystoresAliasesKeyCertFile#basic_constraints}
	BasicConstraints *string `field:"optional" json:"basicConstraints" yaml:"basicConstraints"`
	// X.509 notAfter validity period in milliseconds since epoch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apigee_keystores_aliases_key_cert_file#expiry_date ApigeeKeystoresAliasesKeyCertFile#expiry_date}
	ExpiryDate *string `field:"optional" json:"expiryDate" yaml:"expiryDate"`
	// X.509 issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apigee_keystores_aliases_key_cert_file#issuer ApigeeKeystoresAliasesKeyCertFile#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Flag that specifies whether the certificate is valid.
	//
	// Flag is set to Yes if the certificate is valid, No if expired, or Not yet if not yet valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apigee_keystores_aliases_key_cert_file#is_valid ApigeeKeystoresAliasesKeyCertFile#is_valid}
	IsValid *string `field:"optional" json:"isValid" yaml:"isValid"`
	// Public key component of the X.509 subject public key info.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apigee_keystores_aliases_key_cert_file#public_key ApigeeKeystoresAliasesKeyCertFile#public_key}
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
	// X.509 serial number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apigee_keystores_aliases_key_cert_file#serial_number ApigeeKeystoresAliasesKeyCertFile#serial_number}
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// X.509 signatureAlgorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apigee_keystores_aliases_key_cert_file#sig_alg_name ApigeeKeystoresAliasesKeyCertFile#sig_alg_name}
	SigAlgName *string `field:"optional" json:"sigAlgName" yaml:"sigAlgName"`
	// X.509 subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apigee_keystores_aliases_key_cert_file#subject ApigeeKeystoresAliasesKeyCertFile#subject}
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
	// X.509 subject alternative names (SANs) extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apigee_keystores_aliases_key_cert_file#subject_alternative_names ApigeeKeystoresAliasesKeyCertFile#subject_alternative_names}
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// X.509 notBefore validity period in milliseconds since epoch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apigee_keystores_aliases_key_cert_file#valid_from ApigeeKeystoresAliasesKeyCertFile#valid_from}
	ValidFrom *string `field:"optional" json:"validFrom" yaml:"validFrom"`
	// X.509 version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/apigee_keystores_aliases_key_cert_file#version ApigeeKeystoresAliasesKeyCertFile#version}
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

