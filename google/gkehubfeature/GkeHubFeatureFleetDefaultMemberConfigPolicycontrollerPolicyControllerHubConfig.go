// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfig struct {
	// Configures the mode of the Policy Controller installation Possible values: ["INSTALL_SPEC_UNSPECIFIED", "INSTALL_SPEC_NOT_INSTALLED", "INSTALL_SPEC_ENABLED", "INSTALL_SPEC_SUSPENDED", "INSTALL_SPEC_DETACHED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/gke_hub_feature#install_spec GkeHubFeature#install_spec}
	InstallSpec *string `field:"required" json:"installSpec" yaml:"installSpec"`
	// Interval for Policy Controller Audit scans (in seconds). When set to 0, this disables audit functionality altogether.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/gke_hub_feature#audit_interval_seconds GkeHubFeature#audit_interval_seconds}
	AuditIntervalSeconds *float64 `field:"optional" json:"auditIntervalSeconds" yaml:"auditIntervalSeconds"`
	// The maximum number of audit violations to be stored in a constraint.
	//
	// If not set, the internal default of 20 will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/gke_hub_feature#constraint_violation_limit GkeHubFeature#constraint_violation_limit}
	ConstraintViolationLimit *float64 `field:"optional" json:"constraintViolationLimit" yaml:"constraintViolationLimit"`
	// deployment_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/gke_hub_feature#deployment_configs GkeHubFeature#deployment_configs}
	DeploymentConfigs interface{} `field:"optional" json:"deploymentConfigs" yaml:"deploymentConfigs"`
	// The set of namespaces that are excluded from Policy Controller checks.
	//
	// Namespaces do not need to currently exist on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/gke_hub_feature#exemptable_namespaces GkeHubFeature#exemptable_namespaces}
	ExemptableNamespaces *[]*string `field:"optional" json:"exemptableNamespaces" yaml:"exemptableNamespaces"`
	// Logs all denies and dry run failures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/gke_hub_feature#log_denies_enabled GkeHubFeature#log_denies_enabled}
	LogDeniesEnabled interface{} `field:"optional" json:"logDeniesEnabled" yaml:"logDeniesEnabled"`
	// monitoring block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/gke_hub_feature#monitoring GkeHubFeature#monitoring}
	Monitoring *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigMonitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Enables the ability to mutate resources using Policy Controller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/gke_hub_feature#mutation_enabled GkeHubFeature#mutation_enabled}
	MutationEnabled interface{} `field:"optional" json:"mutationEnabled" yaml:"mutationEnabled"`
	// policy_content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/gke_hub_feature#policy_content GkeHubFeature#policy_content}
	PolicyContent *GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContent `field:"optional" json:"policyContent" yaml:"policyContent"`
	// Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/gke_hub_feature#referential_rules_enabled GkeHubFeature#referential_rules_enabled}
	ReferentialRulesEnabled interface{} `field:"optional" json:"referentialRulesEnabled" yaml:"referentialRulesEnabled"`
}

