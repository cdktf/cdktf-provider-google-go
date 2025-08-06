// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsBigQueryTargetCadence struct {
	// inspect_template_modified_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/data_loss_prevention_discovery_config#inspect_template_modified_cadence DataLossPreventionDiscoveryConfig#inspect_template_modified_cadence}
	InspectTemplateModifiedCadence *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetCadenceInspectTemplateModifiedCadence `field:"optional" json:"inspectTemplateModifiedCadence" yaml:"inspectTemplateModifiedCadence"`
	// schema_modified_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/data_loss_prevention_discovery_config#schema_modified_cadence DataLossPreventionDiscoveryConfig#schema_modified_cadence}
	SchemaModifiedCadence *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetCadenceSchemaModifiedCadence `field:"optional" json:"schemaModifiedCadence" yaml:"schemaModifiedCadence"`
	// table_modified_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/data_loss_prevention_discovery_config#table_modified_cadence DataLossPreventionDiscoveryConfig#table_modified_cadence}
	TableModifiedCadence *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetCadenceTableModifiedCadence `field:"optional" json:"tableModifiedCadence" yaml:"tableModifiedCadence"`
}

