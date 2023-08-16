package datalosspreventionstoredinfotype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionStoredInfoTypeConfig struct {
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
	// The parent of the info type in any of the following formats:.
	//
	// 'projects/{{project}}'
	// 'projects/{{project}}/locations/{{location}}'
	// 'organizations/{{organization_id}}'
	// 'organizations/{{organization_id}}/locations/{{location}}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_stored_info_type#parent DataLossPreventionStoredInfoType#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// A description of the info type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_stored_info_type#description DataLossPreventionStoredInfoType#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dictionary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_stored_info_type#dictionary DataLossPreventionStoredInfoType#dictionary}
	Dictionary *DataLossPreventionStoredInfoTypeDictionary `field:"optional" json:"dictionary" yaml:"dictionary"`
	// User set display name of the info type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_stored_info_type#display_name DataLossPreventionStoredInfoType#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_stored_info_type#id DataLossPreventionStoredInfoType#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// large_custom_dictionary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_stored_info_type#large_custom_dictionary DataLossPreventionStoredInfoType#large_custom_dictionary}
	LargeCustomDictionary *DataLossPreventionStoredInfoTypeLargeCustomDictionary `field:"optional" json:"largeCustomDictionary" yaml:"largeCustomDictionary"`
	// regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_stored_info_type#regex DataLossPreventionStoredInfoType#regex}
	Regex *DataLossPreventionStoredInfoTypeRegex `field:"optional" json:"regex" yaml:"regex"`
	// The storedInfoType ID can contain uppercase and lowercase letters, numbers, and hyphens;
	//
	// that is, it must match the regular expression: [a-zA-Z\d-_]+. The maximum length is 100
	// characters. Can be empty to allow the system to generate one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_stored_info_type#stored_info_type_id DataLossPreventionStoredInfoType#stored_info_type_id}
	StoredInfoTypeId *string `field:"optional" json:"storedInfoTypeId" yaml:"storedInfoTypeId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_stored_info_type#timeouts DataLossPreventionStoredInfoType#timeouts}
	Timeouts *DataLossPreventionStoredInfoTypeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

