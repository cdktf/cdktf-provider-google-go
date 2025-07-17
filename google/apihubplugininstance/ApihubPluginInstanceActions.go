// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubplugininstance


type ApihubPluginInstanceActions struct {
	// This should map to one of the action id specified in actions_config in the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/apihub_plugin_instance#action_id ApihubPluginInstance#action_id}
	ActionId *string `field:"required" json:"actionId" yaml:"actionId"`
	// curation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/apihub_plugin_instance#curation_config ApihubPluginInstance#curation_config}
	CurationConfig *ApihubPluginInstanceActionsCurationConfig `field:"optional" json:"curationConfig" yaml:"curationConfig"`
	// The schedule for this plugin instance action.
	//
	// This can only be set if the
	// plugin supports API_HUB_SCHEDULE_TRIGGER mode for this action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/apihub_plugin_instance#schedule_cron_expression ApihubPluginInstance#schedule_cron_expression}
	ScheduleCronExpression *string `field:"optional" json:"scheduleCronExpression" yaml:"scheduleCronExpression"`
	// The time zone for the schedule cron expression. If not provided, UTC will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/apihub_plugin_instance#schedule_time_zone ApihubPluginInstance#schedule_time_zone}
	ScheduleTimeZone *string `field:"optional" json:"scheduleTimeZone" yaml:"scheduleTimeZone"`
}

