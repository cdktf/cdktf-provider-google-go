// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chronicledataaccessscope


type ChronicleDataAccessScopeDeniedDataAccessLabels struct {
	// The asset namespace configured in the forwarder of the customer's events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/chronicle_data_access_scope#asset_namespace ChronicleDataAccessScope#asset_namespace}
	AssetNamespace *string `field:"optional" json:"assetNamespace" yaml:"assetNamespace"`
	// The name of the data access label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/chronicle_data_access_scope#data_access_label ChronicleDataAccessScope#data_access_label}
	DataAccessLabel *string `field:"optional" json:"dataAccessLabel" yaml:"dataAccessLabel"`
	// ingestion_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/chronicle_data_access_scope#ingestion_label ChronicleDataAccessScope#ingestion_label}
	IngestionLabel *ChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel `field:"optional" json:"ingestionLabel" yaml:"ingestionLabel"`
	// The name of the log type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/chronicle_data_access_scope#log_type ChronicleDataAccessScope#log_type}
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
}

