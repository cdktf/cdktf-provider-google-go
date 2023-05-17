package binaryauthorizationattestor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BinaryAuthorizationAttestorConfig struct {
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
	// attestation_authority_note block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/binary_authorization_attestor#attestation_authority_note BinaryAuthorizationAttestor#attestation_authority_note}
	AttestationAuthorityNote *BinaryAuthorizationAttestorAttestationAuthorityNote `field:"required" json:"attestationAuthorityNote" yaml:"attestationAuthorityNote"`
	// The resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/binary_authorization_attestor#name BinaryAuthorizationAttestor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/binary_authorization_attestor#description BinaryAuthorizationAttestor#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/binary_authorization_attestor#id BinaryAuthorizationAttestor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/binary_authorization_attestor#project BinaryAuthorizationAttestor#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/binary_authorization_attestor#timeouts BinaryAuthorizationAttestor#timeouts}
	Timeouts *BinaryAuthorizationAttestorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

