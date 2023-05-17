package privatecacertificatetemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatecaCertificateTemplateConfig struct {
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
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_template#location PrivatecaCertificateTemplate#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name for this CertificateTemplate in the format `projects/*\/locations/*\/certificateTemplates/*`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_template#name PrivatecaCertificateTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional. A human-readable description of scenarios this template is intended for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_template#description PrivatecaCertificateTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_template#id PrivatecaCertificateTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_template#identity_constraints PrivatecaCertificateTemplate#identity_constraints}
	IdentityConstraints *PrivatecaCertificateTemplateIdentityConstraints `field:"optional" json:"identityConstraints" yaml:"identityConstraints"`
	// Optional. Labels with user-defined metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_template#labels PrivatecaCertificateTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// passthrough_extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_template#passthrough_extensions PrivatecaCertificateTemplate#passthrough_extensions}
	PassthroughExtensions *PrivatecaCertificateTemplatePassthroughExtensions `field:"optional" json:"passthroughExtensions" yaml:"passthroughExtensions"`
	// predefined_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_template#predefined_values PrivatecaCertificateTemplate#predefined_values}
	PredefinedValues *PrivatecaCertificateTemplatePredefinedValues `field:"optional" json:"predefinedValues" yaml:"predefinedValues"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_template#project PrivatecaCertificateTemplate#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/privateca_certificate_template#timeouts PrivatecaCertificateTemplate#timeouts}
	Timeouts *PrivatecaCertificateTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

