// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership


type GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfig struct {
	// Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gke_hub_feature_membership#audit_interval_seconds GkeHubFeatureMembership#audit_interval_seconds}
	AuditIntervalSeconds *float64 `field:"optional" json:"auditIntervalSeconds" yaml:"auditIntervalSeconds"`
	// The maximum number of audit violations to be stored in a constraint.
	//
	// If not set, the internal default of 20 will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gke_hub_feature_membership#constraint_violation_limit GkeHubFeatureMembership#constraint_violation_limit}
	ConstraintViolationLimit *float64 `field:"optional" json:"constraintViolationLimit" yaml:"constraintViolationLimit"`
	// deployment_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gke_hub_feature_membership#deployment_configs GkeHubFeatureMembership#deployment_configs}
	DeploymentConfigs interface{} `field:"optional" json:"deploymentConfigs" yaml:"deploymentConfigs"`
	// The set of namespaces that are excluded from Policy Controller checks.
	//
	// Namespaces do not need to currently exist on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gke_hub_feature_membership#exemptable_namespaces GkeHubFeatureMembership#exemptable_namespaces}
	ExemptableNamespaces *[]*string `field:"optional" json:"exemptableNamespaces" yaml:"exemptableNamespaces"`
	// Configures the mode of the Policy Controller installation. Possible values: INSTALL_SPEC_UNSPECIFIED, INSTALL_SPEC_NOT_INSTALLED, INSTALL_SPEC_ENABLED, INSTALL_SPEC_SUSPENDED, INSTALL_SPEC_DETACHED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gke_hub_feature_membership#install_spec GkeHubFeatureMembership#install_spec}
	InstallSpec *string `field:"optional" json:"installSpec" yaml:"installSpec"`
	// Logs all denies and dry run failures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gke_hub_feature_membership#log_denies_enabled GkeHubFeatureMembership#log_denies_enabled}
	LogDeniesEnabled interface{} `field:"optional" json:"logDeniesEnabled" yaml:"logDeniesEnabled"`
	// monitoring block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gke_hub_feature_membership#monitoring GkeHubFeatureMembership#monitoring}
	Monitoring *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Enables the ability to mutate resources using Policy Controller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gke_hub_feature_membership#mutation_enabled GkeHubFeatureMembership#mutation_enabled}
	MutationEnabled interface{} `field:"optional" json:"mutationEnabled" yaml:"mutationEnabled"`
	// policy_content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gke_hub_feature_membership#policy_content GkeHubFeatureMembership#policy_content}
	PolicyContent *GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent `field:"optional" json:"policyContent" yaml:"policyContent"`
	// Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/gke_hub_feature_membership#referential_rules_enabled GkeHubFeatureMembership#referential_rules_enabled}
	ReferentialRulesEnabled interface{} `field:"optional" json:"referentialRulesEnabled" yaml:"referentialRulesEnabled"`
}

