package gameservicesgameservercluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GameServicesGameServerClusterConfig struct {
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
	// Required. The resource name of the game server cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/game_services_game_server_cluster#cluster_id GameServicesGameServerCluster#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// connection_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/game_services_game_server_cluster#connection_info GameServicesGameServerCluster#connection_info}
	ConnectionInfo *GameServicesGameServerClusterConnectionInfo `field:"required" json:"connectionInfo" yaml:"connectionInfo"`
	// The realm id of the game server realm.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/game_services_game_server_cluster#realm_id GameServicesGameServerCluster#realm_id}
	RealmId *string `field:"required" json:"realmId" yaml:"realmId"`
	// Human readable description of the cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/game_services_game_server_cluster#description GameServicesGameServerCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/game_services_game_server_cluster#id GameServicesGameServerCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels associated with this game server cluster. Each label is a key-value pair.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/game_services_game_server_cluster#labels GameServicesGameServerCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Location of the Cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/game_services_game_server_cluster#location GameServicesGameServerCluster#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/game_services_game_server_cluster#project GameServicesGameServerCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/game_services_game_server_cluster#timeouts GameServicesGameServerCluster#timeouts}
	Timeouts *GameServicesGameServerClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

