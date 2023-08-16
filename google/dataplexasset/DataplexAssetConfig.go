package dataplexasset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexAssetConfig struct {
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
	// The zone for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_asset#dataplex_zone DataplexAsset#dataplex_zone}
	DataplexZone *string `field:"required" json:"dataplexZone" yaml:"dataplexZone"`
	// discovery_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_asset#discovery_spec DataplexAsset#discovery_spec}
	DiscoverySpec *DataplexAssetDiscoverySpec `field:"required" json:"discoverySpec" yaml:"discoverySpec"`
	// The lake for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_asset#lake DataplexAsset#lake}
	Lake *string `field:"required" json:"lake" yaml:"lake"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_asset#location DataplexAsset#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_asset#name DataplexAsset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_asset#resource_spec DataplexAsset#resource_spec}
	ResourceSpec *DataplexAssetResourceSpec `field:"required" json:"resourceSpec" yaml:"resourceSpec"`
	// Optional. Description of the asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_asset#description DataplexAsset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional. User friendly display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_asset#display_name DataplexAsset#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_asset#id DataplexAsset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. User defined labels for the asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_asset#labels DataplexAsset#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_asset#project DataplexAsset#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_asset#timeouts DataplexAsset#timeouts}
	Timeouts *DataplexAssetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

