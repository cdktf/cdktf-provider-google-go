package bigtableappprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigtableAppProfileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_app_profile#app_profile_id BigtableAppProfile#app_profile_id}
	AppProfileId *string `field:"required" json:"appProfileId" yaml:"appProfileId"`
	// Long form description of the use case for this app profile.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_app_profile#description BigtableAppProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_app_profile#id BigtableAppProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, ignore safety checks when deleting/updating the app profile.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_app_profile#ignore_warnings BigtableAppProfile#ignore_warnings}
	IgnoreWarnings interface{} `field:"optional" json:"ignoreWarnings" yaml:"ignoreWarnings"`
	// The name of the instance to create the app profile within.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_app_profile#instance BigtableAppProfile#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
	// The set of clusters to route to.
	//
	// The order is ignored; clusters will be tried in order of distance. If left empty, all clusters are eligible.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_app_profile#multi_cluster_routing_cluster_ids BigtableAppProfile#multi_cluster_routing_cluster_ids}
	MultiClusterRoutingClusterIds *[]*string `field:"optional" json:"multiClusterRoutingClusterIds" yaml:"multiClusterRoutingClusterIds"`
	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available in the event of transient errors or delays.
	//
	// Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
	// consistency to improve availability.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_app_profile#multi_cluster_routing_use_any BigtableAppProfile#multi_cluster_routing_use_any}
	MultiClusterRoutingUseAny interface{} `field:"optional" json:"multiClusterRoutingUseAny" yaml:"multiClusterRoutingUseAny"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_app_profile#project BigtableAppProfile#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// single_cluster_routing block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_app_profile#single_cluster_routing BigtableAppProfile#single_cluster_routing}
	SingleClusterRouting *BigtableAppProfileSingleClusterRouting `field:"optional" json:"singleClusterRouting" yaml:"singleClusterRouting"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_app_profile#timeouts BigtableAppProfile#timeouts}
	Timeouts *BigtableAppProfileTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

