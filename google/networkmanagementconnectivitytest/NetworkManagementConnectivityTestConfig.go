package networkmanagementconnectivitytest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkManagementConnectivityTestConfig struct {
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
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_management_connectivity_test#destination NetworkManagementConnectivityTest#destination}
	Destination *NetworkManagementConnectivityTestDestination `field:"required" json:"destination" yaml:"destination"`
	// Unique name for the connectivity test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_management_connectivity_test#name NetworkManagementConnectivityTest#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_management_connectivity_test#source NetworkManagementConnectivityTest#source}
	Source *NetworkManagementConnectivityTestSource `field:"required" json:"source" yaml:"source"`
	// The user-supplied description of the Connectivity Test. Maximum of 512 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_management_connectivity_test#description NetworkManagementConnectivityTest#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_management_connectivity_test#id NetworkManagementConnectivityTest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels to represent user-provided metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_management_connectivity_test#labels NetworkManagementConnectivityTest#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_management_connectivity_test#project NetworkManagementConnectivityTest#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_management_connectivity_test#protocol NetworkManagementConnectivityTest#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Other projects that may be relevant for reachability analysis. This is applicable to scenarios where a test can cross project boundaries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_management_connectivity_test#related_projects NetworkManagementConnectivityTest#related_projects}
	RelatedProjects *[]*string `field:"optional" json:"relatedProjects" yaml:"relatedProjects"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_management_connectivity_test#timeouts NetworkManagementConnectivityTest#timeouts}
	Timeouts *NetworkManagementConnectivityTestTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

