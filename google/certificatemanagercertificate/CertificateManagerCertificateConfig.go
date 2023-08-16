package certificatemanagercertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CertificateManagerCertificateConfig struct {
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
	// A user-defined name of the certificate.
	//
	// Certificate names must be unique
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#name CertificateManagerCertificate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#description CertificateManagerCertificate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#id CertificateManagerCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the Certificate resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#labels CertificateManagerCertificate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The Certificate Manager location. If not specified, "global" is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#location CertificateManagerCertificate#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// managed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#managed CertificateManagerCertificate#managed}
	Managed *CertificateManagerCertificateManaged `field:"optional" json:"managed" yaml:"managed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#project CertificateManagerCertificate#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The scope of the certificate.
	//
	// DEFAULT: Certificates with default scope are served from core Google data centers.
	// If unsure, choose this option.
	//
	// EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
	// served from non-core Google data centers.
	// Currently allowed only for managed certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#scope CertificateManagerCertificate#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// self_managed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#self_managed CertificateManagerCertificate#self_managed}
	SelfManaged *CertificateManagerCertificateSelfManaged `field:"optional" json:"selfManaged" yaml:"selfManaged"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/certificate_manager_certificate#timeouts CertificateManagerCertificate#timeouts}
	Timeouts *CertificateManagerCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

