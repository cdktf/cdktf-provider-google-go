// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebucket


type StorageBucketLifecycleRuleCondition struct {
	// Minimum age of an object in days to satisfy this condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#age StorageBucket#age}
	Age *float64 `field:"optional" json:"age" yaml:"age"`
	// Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#created_before StorageBucket#created_before}
	CreatedBefore *string `field:"optional" json:"createdBefore" yaml:"createdBefore"`
	// Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#custom_time_before StorageBucket#custom_time_before}
	CustomTimeBefore *string `field:"optional" json:"customTimeBefore" yaml:"customTimeBefore"`
	// Number of days elapsed since the user-specified timestamp set on an object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#days_since_custom_time StorageBucket#days_since_custom_time}
	DaysSinceCustomTime *float64 `field:"optional" json:"daysSinceCustomTime" yaml:"daysSinceCustomTime"`
	// Number of days elapsed since the noncurrent timestamp of an object. This 										condition is relevant only for versioned objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#days_since_noncurrent_time StorageBucket#days_since_noncurrent_time}
	DaysSinceNoncurrentTime *float64 `field:"optional" json:"daysSinceNoncurrentTime" yaml:"daysSinceNoncurrentTime"`
	// One or more matching name prefixes to satisfy this condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#matches_prefix StorageBucket#matches_prefix}
	MatchesPrefix *[]*string `field:"optional" json:"matchesPrefix" yaml:"matchesPrefix"`
	// Storage Class of objects to satisfy this condition. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, STANDARD, DURABLE_REDUCED_AVAILABILITY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#matches_storage_class StorageBucket#matches_storage_class}
	MatchesStorageClass *[]*string `field:"optional" json:"matchesStorageClass" yaml:"matchesStorageClass"`
	// One or more matching name suffixes to satisfy this condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#matches_suffix StorageBucket#matches_suffix}
	MatchesSuffix *[]*string `field:"optional" json:"matchesSuffix" yaml:"matchesSuffix"`
	// Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#noncurrent_time_before StorageBucket#noncurrent_time_before}
	NoncurrentTimeBefore *string `field:"optional" json:"noncurrentTimeBefore" yaml:"noncurrentTimeBefore"`
	// Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#num_newer_versions StorageBucket#num_newer_versions}
	NumNewerVersions *float64 `field:"optional" json:"numNewerVersions" yaml:"numNewerVersions"`
	// While set true, age value will be sent in the request even for zero value of the field.
	//
	// This field is only useful for setting 0 value to the age field. It can be used alone or together with age.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#send_age_if_zero StorageBucket#send_age_if_zero}
	SendAgeIfZero interface{} `field:"optional" json:"sendAgeIfZero" yaml:"sendAgeIfZero"`
	// While set true, days_since_custom_time value will be sent in the request even for zero value of the field.
	//
	// This field is only useful for setting 0 value to the days_since_custom_time field. It can be used alone or together with days_since_custom_time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#send_days_since_custom_time_if_zero StorageBucket#send_days_since_custom_time_if_zero}
	SendDaysSinceCustomTimeIfZero interface{} `field:"optional" json:"sendDaysSinceCustomTimeIfZero" yaml:"sendDaysSinceCustomTimeIfZero"`
	// While set true, days_since_noncurrent_time value will be sent in the request even for zero value of the field.
	//
	// This field is only useful for setting 0 value to the days_since_noncurrent_time field. It can be used alone or together with days_since_noncurrent_time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#send_days_since_noncurrent_time_if_zero StorageBucket#send_days_since_noncurrent_time_if_zero}
	SendDaysSinceNoncurrentTimeIfZero interface{} `field:"optional" json:"sendDaysSinceNoncurrentTimeIfZero" yaml:"sendDaysSinceNoncurrentTimeIfZero"`
	// While set true, num_newer_versions value will be sent in the request even for zero value of the field.
	//
	// This field is only useful for setting 0 value to the num_newer_versions field. It can be used alone or together with num_newer_versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#send_num_newer_versions_if_zero StorageBucket#send_num_newer_versions_if_zero}
	SendNumNewerVersionsIfZero interface{} `field:"optional" json:"sendNumNewerVersionsIfZero" yaml:"sendNumNewerVersionsIfZero"`
	// Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: "LIVE", "ARCHIVED", "ANY".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/storage_bucket#with_state StorageBucket#with_state}
	WithState *string `field:"optional" json:"withState" yaml:"withState"`
}

