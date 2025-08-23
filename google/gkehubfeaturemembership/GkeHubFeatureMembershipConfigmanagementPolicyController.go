// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership


type GkeHubFeatureMembershipConfigmanagementPolicyController struct {
	// Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/gke_hub_feature_membership#audit_interval_seconds GkeHubFeatureMembership#audit_interval_seconds}
	AuditIntervalSeconds *string `field:"optional" json:"auditIntervalSeconds" yaml:"auditIntervalSeconds"`
	// Enables the installation of Policy Controller. If false, the rest of PolicyController fields take no effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/gke_hub_feature_membership#enabled GkeHubFeatureMembership#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The set of namespaces that are excluded from Policy Controller checks.
	//
	// Namespaces do not need to currently exist on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/gke_hub_feature_membership#exemptable_namespaces GkeHubFeatureMembership#exemptable_namespaces}
	ExemptableNamespaces *[]*string `field:"optional" json:"exemptableNamespaces" yaml:"exemptableNamespaces"`
	// Logs all denies and dry run failures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/gke_hub_feature_membership#log_denies_enabled GkeHubFeatureMembership#log_denies_enabled}
	LogDeniesEnabled interface{} `field:"optional" json:"logDeniesEnabled" yaml:"logDeniesEnabled"`
	// monitoring block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/gke_hub_feature_membership#monitoring GkeHubFeatureMembership#monitoring}
	Monitoring *GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Enable or disable mutation in policy controller.
	//
	// If true, mutation CRDs, webhook and controller deployment will be deployed to the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/gke_hub_feature_membership#mutation_enabled GkeHubFeatureMembership#mutation_enabled}
	MutationEnabled interface{} `field:"optional" json:"mutationEnabled" yaml:"mutationEnabled"`
	// Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/gke_hub_feature_membership#referential_rules_enabled GkeHubFeatureMembership#referential_rules_enabled}
	ReferentialRulesEnabled interface{} `field:"optional" json:"referentialRulesEnabled" yaml:"referentialRulesEnabled"`
	// Installs the default template library along with Policy Controller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/gke_hub_feature_membership#template_library_installed GkeHubFeatureMembership#template_library_installed}
	TemplateLibraryInstalled interface{} `field:"optional" json:"templateLibraryInstalled" yaml:"templateLibraryInstalled"`
}

