package dataplexdatascan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexDatascanConfig struct {
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
	// data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#data DataplexDatascan#data}
	Data *DataplexDatascanData `field:"required" json:"data" yaml:"data"`
	// DataScan identifier.
	//
	// Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#data_scan_id DataplexDatascan#data_scan_id}
	DataScanId *string `field:"required" json:"dataScanId" yaml:"dataScanId"`
	// execution_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#execution_spec DataplexDatascan#execution_spec}
	ExecutionSpec *DataplexDatascanExecutionSpec `field:"required" json:"executionSpec" yaml:"executionSpec"`
	// The location where the data scan should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#location DataplexDatascan#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// data_profile_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#data_profile_spec DataplexDatascan#data_profile_spec}
	DataProfileSpec *DataplexDatascanDataProfileSpec `field:"optional" json:"dataProfileSpec" yaml:"dataProfileSpec"`
	// data_quality_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#data_quality_spec DataplexDatascan#data_quality_spec}
	DataQualitySpec *DataplexDatascanDataQualitySpec `field:"optional" json:"dataQualitySpec" yaml:"dataQualitySpec"`
	// Description of the scan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#description DataplexDatascan#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User friendly display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#display_name DataplexDatascan#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#id DataplexDatascan#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels for the scan. A list of key->value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#labels DataplexDatascan#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#project DataplexDatascan#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#timeouts DataplexDatascan#timeouts}
	Timeouts *DataplexDatascanTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

