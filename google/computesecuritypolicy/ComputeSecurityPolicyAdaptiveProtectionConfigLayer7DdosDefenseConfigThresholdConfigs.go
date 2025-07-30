// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicy


type ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfigThresholdConfigs struct {
	// The name must be 1-63 characters long, and comply with RFC1035.
	//
	// The name must be unique within the security policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_security_policy#name ComputeSecurityPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_security_policy#auto_deploy_confidence_threshold ComputeSecurityPolicy#auto_deploy_confidence_threshold}.
	AutoDeployConfidenceThreshold *float64 `field:"optional" json:"autoDeployConfidenceThreshold" yaml:"autoDeployConfidenceThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_security_policy#auto_deploy_expiration_sec ComputeSecurityPolicy#auto_deploy_expiration_sec}.
	AutoDeployExpirationSec *float64 `field:"optional" json:"autoDeployExpirationSec" yaml:"autoDeployExpirationSec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_security_policy#auto_deploy_impacted_baseline_threshold ComputeSecurityPolicy#auto_deploy_impacted_baseline_threshold}.
	AutoDeployImpactedBaselineThreshold *float64 `field:"optional" json:"autoDeployImpactedBaselineThreshold" yaml:"autoDeployImpactedBaselineThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_security_policy#auto_deploy_load_threshold ComputeSecurityPolicy#auto_deploy_load_threshold}.
	AutoDeployLoadThreshold *float64 `field:"optional" json:"autoDeployLoadThreshold" yaml:"autoDeployLoadThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_security_policy#detection_absolute_qps ComputeSecurityPolicy#detection_absolute_qps}.
	DetectionAbsoluteQps *float64 `field:"optional" json:"detectionAbsoluteQps" yaml:"detectionAbsoluteQps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_security_policy#detection_load_threshold ComputeSecurityPolicy#detection_load_threshold}.
	DetectionLoadThreshold *float64 `field:"optional" json:"detectionLoadThreshold" yaml:"detectionLoadThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_security_policy#detection_relative_to_baseline_qps ComputeSecurityPolicy#detection_relative_to_baseline_qps}.
	DetectionRelativeToBaselineQps *float64 `field:"optional" json:"detectionRelativeToBaselineQps" yaml:"detectionRelativeToBaselineQps"`
	// traffic_granularity_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_security_policy#traffic_granularity_configs ComputeSecurityPolicy#traffic_granularity_configs}
	TrafficGranularityConfigs interface{} `field:"optional" json:"trafficGranularityConfigs" yaml:"trafficGranularityConfigs"`
}

