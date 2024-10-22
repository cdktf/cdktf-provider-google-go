// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetGenerationCadence struct {
	// inspect_template_modified_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/data_loss_prevention_discovery_config#inspect_template_modified_cadence DataLossPreventionDiscoveryConfig#inspect_template_modified_cadence}
	InspectTemplateModifiedCadence *DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetGenerationCadenceInspectTemplateModifiedCadence `field:"optional" json:"inspectTemplateModifiedCadence" yaml:"inspectTemplateModifiedCadence"`
	// Data changes (non-schema changes) in Cloud SQL tables can't trigger reprofiling.
	//
	// If you set this field, profiles are refreshed at this frequency regardless of whether the underlying tables have changes. Defaults to never. Possible values: ["UPDATE_FREQUENCY_NEVER", "UPDATE_FREQUENCY_DAILY", "UPDATE_FREQUENCY_MONTHLY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/data_loss_prevention_discovery_config#refresh_frequency DataLossPreventionDiscoveryConfig#refresh_frequency}
	RefreshFrequency *string `field:"optional" json:"refreshFrequency" yaml:"refreshFrequency"`
	// schema_modified_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/data_loss_prevention_discovery_config#schema_modified_cadence DataLossPreventionDiscoveryConfig#schema_modified_cadence}
	SchemaModifiedCadence *DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetGenerationCadenceSchemaModifiedCadence `field:"optional" json:"schemaModifiedCadence" yaml:"schemaModifiedCadence"`
}

