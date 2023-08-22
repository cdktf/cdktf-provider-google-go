package monitoringgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitoringGroupConfig struct {
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
	// A user-assigned name for this group, used only for display purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/monitoring_group#display_name MonitoringGroup#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The filter used to determine which monitored resources belong to this group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/monitoring_group#filter MonitoringGroup#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/monitoring_group#id MonitoringGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, the members of this group are considered to be a cluster.
	//
	// The system can perform additional analysis on
	// groups that are clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/monitoring_group#is_cluster MonitoringGroup#is_cluster}
	IsCluster interface{} `field:"optional" json:"isCluster" yaml:"isCluster"`
	// The name of the group's parent, if it has one.
	//
	// The format is
	// "projects/{project_id_or_number}/groups/{group_id}". For
	// groups with no parent, parentName is the empty string, "".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/monitoring_group#parent_name MonitoringGroup#parent_name}
	ParentName *string `field:"optional" json:"parentName" yaml:"parentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/monitoring_group#project MonitoringGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/monitoring_group#timeouts MonitoringGroup#timeouts}
	Timeouts *MonitoringGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

