package datalosspreventioninspecttemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionInspectTemplateConfig struct {
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
	// The parent of the inspect template in any of the following formats:.
	//
	// 'projects/{{project}}'
	// 'projects/{{project}}/locations/{{location}}'
	// 'organizations/{{organization_id}}'
	// 'organizations/{{organization_id}}/locations/{{location}}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_inspect_template#parent DataLossPreventionInspectTemplate#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// A description of the inspect template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_inspect_template#description DataLossPreventionInspectTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User set display name of the inspect template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_inspect_template#display_name DataLossPreventionInspectTemplate#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_inspect_template#id DataLossPreventionInspectTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// inspect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_inspect_template#inspect_config DataLossPreventionInspectTemplate#inspect_config}
	InspectConfig *DataLossPreventionInspectTemplateInspectConfig `field:"optional" json:"inspectConfig" yaml:"inspectConfig"`
	// The template id can contain uppercase and lowercase letters, numbers, and hyphens;
	//
	// that is, it must match the regular expression: [a-zA-Z\d-_]+. The maximum length is
	// 100 characters. Can be empty to allow the system to generate one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_inspect_template#template_id DataLossPreventionInspectTemplate#template_id}
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_inspect_template#timeouts DataLossPreventionInspectTemplate#timeouts}
	Timeouts *DataLossPreventionInspectTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

