package privatecacertificateauthority

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatecaCertificateAuthorityConfig struct {
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
	// The user provided Resource ID for this Certificate Authority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#certificate_authority_id PrivatecaCertificateAuthority#certificate_authority_id}
	CertificateAuthorityId *string `field:"required" json:"certificateAuthorityId" yaml:"certificateAuthorityId"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#config PrivatecaCertificateAuthority#config}
	Config *PrivatecaCertificateAuthorityConfigA `field:"required" json:"config" yaml:"config"`
	// key_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#key_spec PrivatecaCertificateAuthority#key_spec}
	KeySpec *PrivatecaCertificateAuthorityKeySpec `field:"required" json:"keySpec" yaml:"keySpec"`
	// Location of the CertificateAuthority. A full list of valid locations can be found by running 'gcloud privateca locations list'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#location PrivatecaCertificateAuthority#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the CaPool this Certificate Authority belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#pool PrivatecaCertificateAuthority#pool}
	Pool *string `field:"required" json:"pool" yaml:"pool"`
	// Whether or not to allow Terraform to destroy the CertificateAuthority.
	//
	// Unless this field is set to false
	// in Terraform state, a 'terraform destroy' or 'terraform apply' that would delete the instance will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#deletion_protection PrivatecaCertificateAuthority#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Desired state of the CertificateAuthority. Set this field to 'STAGED' to create a 'STAGED' root CA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#desired_state PrivatecaCertificateAuthority#desired_state}
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs.
	//
	// This must be a bucket name, without any prefixes
	// (such as 'gs://') or suffixes (such as '.googleapis.com'). For example, to use a bucket named
	// my-bucket, you would simply specify 'my-bucket'. If not specified, a managed bucket will be
	// created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#gcs_bucket PrivatecaCertificateAuthority#gcs_bucket}
	GcsBucket *string `field:"optional" json:"gcsBucket" yaml:"gcsBucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#id PrivatecaCertificateAuthority#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// This field allows the CA to be deleted even if the CA has active certs.
	//
	// Active certs include both unrevoked and unexpired certs.
	// Use with care. Defaults to 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#ignore_active_certificates_on_deletion PrivatecaCertificateAuthority#ignore_active_certificates_on_deletion}
	IgnoreActiveCertificatesOnDeletion interface{} `field:"optional" json:"ignoreActiveCertificatesOnDeletion" yaml:"ignoreActiveCertificatesOnDeletion"`
	// Labels with user-defined metadata.
	//
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#labels PrivatecaCertificateAuthority#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The desired lifetime of the CA certificate.
	//
	// Used to create the "notBeforeTime" and
	// "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
	// fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#lifetime PrivatecaCertificateAuthority#lifetime}
	Lifetime *string `field:"optional" json:"lifetime" yaml:"lifetime"`
	// The signed CA certificate issued from the subordinated CA's CSR.
	//
	// This is needed when activating the subordiante CA with a third party issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#pem_ca_certificate PrivatecaCertificateAuthority#pem_ca_certificate}
	PemCaCertificate *string `field:"optional" json:"pemCaCertificate" yaml:"pemCaCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#project PrivatecaCertificateAuthority#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// If this flag is set, the Certificate Authority will be deleted as soon as possible without a 30-day grace period where undeletion would have been allowed.
	//
	// If you proceed, there will be no way to recover this CA.
	// Use with care. Defaults to 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#skip_grace_period PrivatecaCertificateAuthority#skip_grace_period}
	SkipGracePeriod interface{} `field:"optional" json:"skipGracePeriod" yaml:"skipGracePeriod"`
	// subordinate_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#subordinate_config PrivatecaCertificateAuthority#subordinate_config}
	SubordinateConfig *PrivatecaCertificateAuthoritySubordinateConfig `field:"optional" json:"subordinateConfig" yaml:"subordinateConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#timeouts PrivatecaCertificateAuthority#timeouts}
	Timeouts *PrivatecaCertificateAuthorityTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The Type of this CertificateAuthority.
	//
	// ~> **Note:** For 'SUBORDINATE' Certificate Authorities, they need to
	// be activated before they can issue certificates. Default value: "SELF_SIGNED" Possible values: ["SELF_SIGNED", "SUBORDINATE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_authority#type PrivatecaCertificateAuthority#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

