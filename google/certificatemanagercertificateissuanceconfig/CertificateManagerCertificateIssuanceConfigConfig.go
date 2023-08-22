package certificatemanagercertificateissuanceconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CertificateManagerCertificateIssuanceConfigConfig struct {
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
	// certificate_authority_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/certificate_manager_certificate_issuance_config#certificate_authority_config CertificateManagerCertificateIssuanceConfig#certificate_authority_config}
	CertificateAuthorityConfig *CertificateManagerCertificateIssuanceConfigCertificateAuthorityConfig `field:"required" json:"certificateAuthorityConfig" yaml:"certificateAuthorityConfig"`
	// Key algorithm to use when generating the private key. Possible values: ["RSA_2048", "ECDSA_P256"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/certificate_manager_certificate_issuance_config#key_algorithm CertificateManagerCertificateIssuanceConfig#key_algorithm}
	KeyAlgorithm *string `field:"required" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Lifetime of issued certificates.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'.
	// Example: "1814400s". Valid values are from 21 days (1814400s) to 30 days (2592000s)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/certificate_manager_certificate_issuance_config#lifetime CertificateManagerCertificateIssuanceConfig#lifetime}
	Lifetime *string `field:"required" json:"lifetime" yaml:"lifetime"`
	// A user-defined name of the certificate issuance config. CertificateIssuanceConfig names must be unique globally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/certificate_manager_certificate_issuance_config#name CertificateManagerCertificateIssuanceConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// It specifies the percentage of elapsed time of the certificate lifetime to wait before renewing the certificate.
	//
	// Must be a number between 1-99, inclusive.
	// You must set the rotation window percentage in relation to the certificate lifetime so that certificate renewal occurs at least 7 days after
	// the certificate has been issued and at least 7 days before it expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/certificate_manager_certificate_issuance_config#rotation_window_percentage CertificateManagerCertificateIssuanceConfig#rotation_window_percentage}
	RotationWindowPercentage *float64 `field:"required" json:"rotationWindowPercentage" yaml:"rotationWindowPercentage"`
	// One or more paragraphs of text description of a CertificateIssuanceConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/certificate_manager_certificate_issuance_config#description CertificateManagerCertificateIssuanceConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/certificate_manager_certificate_issuance_config#id CertificateManagerCertificateIssuanceConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// 'Set of label tags associated with the CertificateIssuanceConfig resource.
	//
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/certificate_manager_certificate_issuance_config#labels CertificateManagerCertificateIssuanceConfig#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The Certificate Manager location. If not specified, "global" is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/certificate_manager_certificate_issuance_config#location CertificateManagerCertificateIssuanceConfig#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/certificate_manager_certificate_issuance_config#project CertificateManagerCertificateIssuanceConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/certificate_manager_certificate_issuance_config#timeouts CertificateManagerCertificateIssuanceConfig#timeouts}
	Timeouts *CertificateManagerCertificateIssuanceConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

