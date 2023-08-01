package gameservicesgameserverconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GameServicesGameServerConfigConfig struct {
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
	// A unique id for the deployment config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/game_services_game_server_config#config_id GameServicesGameServerConfig#config_id}
	ConfigId *string `field:"required" json:"configId" yaml:"configId"`
	// A unique id for the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/game_services_game_server_config#deployment_id GameServicesGameServerConfig#deployment_id}
	DeploymentId *string `field:"required" json:"deploymentId" yaml:"deploymentId"`
	// fleet_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/game_services_game_server_config#fleet_configs GameServicesGameServerConfig#fleet_configs}
	FleetConfigs interface{} `field:"required" json:"fleetConfigs" yaml:"fleetConfigs"`
	// The description of the game server config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/game_services_game_server_config#description GameServicesGameServerConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/game_services_game_server_config#id GameServicesGameServerConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels associated with this game server config. Each label is a key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/game_services_game_server_config#labels GameServicesGameServerConfig#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Location of the Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/game_services_game_server_config#location GameServicesGameServerConfig#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/game_services_game_server_config#project GameServicesGameServerConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// scaling_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/game_services_game_server_config#scaling_configs GameServicesGameServerConfig#scaling_configs}
	ScalingConfigs interface{} `field:"optional" json:"scalingConfigs" yaml:"scalingConfigs"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/game_services_game_server_config#timeouts GameServicesGameServerConfig#timeouts}
	Timeouts *GameServicesGameServerConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

