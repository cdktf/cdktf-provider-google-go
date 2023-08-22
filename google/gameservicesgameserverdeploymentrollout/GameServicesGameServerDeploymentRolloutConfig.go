package gameservicesgameserverdeploymentrollout

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GameServicesGameServerDeploymentRolloutConfig struct {
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
	// This field points to the game server config that is applied by default to all realms and clusters. For example,.
	//
	// 'projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/game_services_game_server_deployment_rollout#default_game_server_config GameServicesGameServerDeploymentRollout#default_game_server_config}
	DefaultGameServerConfig *string `field:"required" json:"defaultGameServerConfig" yaml:"defaultGameServerConfig"`
	// The deployment to rollout the new config to. Only 1 rollout must be associated with each deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/game_services_game_server_deployment_rollout#deployment_id GameServicesGameServerDeploymentRollout#deployment_id}
	DeploymentId *string `field:"required" json:"deploymentId" yaml:"deploymentId"`
	// game_server_config_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/game_services_game_server_deployment_rollout#game_server_config_overrides GameServicesGameServerDeploymentRollout#game_server_config_overrides}
	GameServerConfigOverrides interface{} `field:"optional" json:"gameServerConfigOverrides" yaml:"gameServerConfigOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/game_services_game_server_deployment_rollout#id GameServicesGameServerDeploymentRollout#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/game_services_game_server_deployment_rollout#project GameServicesGameServerDeploymentRollout#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/game_services_game_server_deployment_rollout#timeouts GameServicesGameServerDeploymentRollout#timeouts}
	Timeouts *GameServicesGameServerDeploymentRolloutTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

