package privatecacertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatecaCertificateConfig struct {
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
	// Location of the Certificate. A full list of valid locations can be found by running 'gcloud privateca locations list'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate#location PrivatecaCertificate#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name for this Certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate#name PrivatecaCertificate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the CaPool this Certificate belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate#pool PrivatecaCertificate#pool}
	Pool *string `field:"required" json:"pool" yaml:"pool"`
	// The Certificate Authority ID that should issue the certificate.
	//
	// For example, to issue a Certificate from
	// a Certificate Authority with resource name 'projects/my-project/locations/us-central1/caPools/my-pool/certificateAuthorities/my-ca',
	// argument 'pool' should be set to 'projects/my-project/locations/us-central1/caPools/my-pool', argument 'certificate_authority'
	// should be set to 'my-ca'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate#certificate_authority PrivatecaCertificate#certificate_authority}
	CertificateAuthority *string `field:"optional" json:"certificateAuthority" yaml:"certificateAuthority"`
	// The resource name for a CertificateTemplate used to issue this certificate, in the format 'projects/*\/locations/*\/certificateTemplates/*'.
	//
	// If this is specified,
	// the caller must have the necessary permission to use this template. If this is
	// omitted, no template will be used. This template must be in the same location
	// as the Certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate#certificate_template PrivatecaCertificate#certificate_template}
	CertificateTemplate *string `field:"optional" json:"certificateTemplate" yaml:"certificateTemplate"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate#config PrivatecaCertificate#config}
	Config *PrivatecaCertificateConfigA `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate#id PrivatecaCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels with user-defined metadata to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate#labels PrivatecaCertificate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The desired lifetime of the CA certificate.
	//
	// Used to create the "notBeforeTime" and
	// "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
	// fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate#lifetime PrivatecaCertificate#lifetime}
	Lifetime *string `field:"optional" json:"lifetime" yaml:"lifetime"`
	// Immutable. A pem-encoded X.509 certificate signing request (CSR).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate#pem_csr PrivatecaCertificate#pem_csr}
	PemCsr *string `field:"optional" json:"pemCsr" yaml:"pemCsr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate#project PrivatecaCertificate#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_certificate#timeouts PrivatecaCertificate#timeouts}
	Timeouts *PrivatecaCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

