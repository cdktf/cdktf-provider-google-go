// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbinstance


type AlloydbInstanceQueryInsightsConfig struct {
	// Number of query execution plans captured by Insights per minute for all queries combined.
	//
	// The default value is 5. Any integer between 0 and 20 is considered valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/alloydb_instance#query_plans_per_minute AlloydbInstance#query_plans_per_minute}
	QueryPlansPerMinute *float64 `field:"optional" json:"queryPlansPerMinute" yaml:"queryPlansPerMinute"`
	// Query string length. The default value is 1024. Any integer between 256 and 4500 is considered valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/alloydb_instance#query_string_length AlloydbInstance#query_string_length}
	QueryStringLength *float64 `field:"optional" json:"queryStringLength" yaml:"queryStringLength"`
	// Record application tags for an instance. This flag is turned "on" by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/alloydb_instance#record_application_tags AlloydbInstance#record_application_tags}
	RecordApplicationTags interface{} `field:"optional" json:"recordApplicationTags" yaml:"recordApplicationTags"`
	// Record client address for an instance. Client address is PII information. This flag is turned "on" by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/alloydb_instance#record_client_address AlloydbInstance#record_client_address}
	RecordClientAddress interface{} `field:"optional" json:"recordClientAddress" yaml:"recordClientAddress"`
}

