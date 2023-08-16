package datagoogleprivatecacertificateauthority

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGooglePrivatecaCertificateAuthorityConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/privateca_certificate_authority#certificate_authority_id DataGooglePrivatecaCertificateAuthority#certificate_authority_id}
	CertificateAuthorityId *string `field:"optional" json:"certificateAuthorityId" yaml:"certificateAuthorityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/privateca_certificate_authority#id DataGooglePrivatecaCertificateAuthority#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Location of the CertificateAuthority. A full list of valid locations can be found by running 'gcloud privateca locations list'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/privateca_certificate_authority#location DataGooglePrivatecaCertificateAuthority#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The name of the CaPool this Certificate Authority belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/privateca_certificate_authority#pool DataGooglePrivatecaCertificateAuthority#pool}
	Pool *string `field:"optional" json:"pool" yaml:"pool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/data-sources/privateca_certificate_authority#project DataGooglePrivatecaCertificateAuthority#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

