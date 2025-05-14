// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacapool


type PrivatecaCaPoolIssuancePolicy struct {
	// allowed_issuance_modes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/privateca_ca_pool#allowed_issuance_modes PrivatecaCaPool#allowed_issuance_modes}
	AllowedIssuanceModes *PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes `field:"optional" json:"allowedIssuanceModes" yaml:"allowedIssuanceModes"`
	// allowed_key_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/privateca_ca_pool#allowed_key_types PrivatecaCaPool#allowed_key_types}
	AllowedKeyTypes interface{} `field:"optional" json:"allowedKeyTypes" yaml:"allowedKeyTypes"`
	// The duration to backdate all certificates issued from this CaPool.
	//
	// If not set, the
	// certificates will be issued with a not_before_time of the issuance time (i.e. the current
	// time). If set, the certificates will be issued with a not_before_time of the issuance
	// time minus the backdate_duration. The not_after_time will be adjusted to preserve the
	// requested lifetime. The backdate_duration must be less than or equal to 48 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/privateca_ca_pool#backdate_duration PrivatecaCaPool#backdate_duration}
	BackdateDuration *string `field:"optional" json:"backdateDuration" yaml:"backdateDuration"`
	// baseline_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/privateca_ca_pool#baseline_values PrivatecaCaPool#baseline_values}
	BaselineValues *PrivatecaCaPoolIssuancePolicyBaselineValues `field:"optional" json:"baselineValues" yaml:"baselineValues"`
	// identity_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/privateca_ca_pool#identity_constraints PrivatecaCaPool#identity_constraints}
	IdentityConstraints *PrivatecaCaPoolIssuancePolicyIdentityConstraints `field:"optional" json:"identityConstraints" yaml:"identityConstraints"`
	// The maximum lifetime allowed for issued Certificates.
	//
	// Note that if the issuing CertificateAuthority
	// expires before a Certificate's requested maximumLifetime, the effective lifetime will be explicitly truncated to match it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/privateca_ca_pool#maximum_lifetime PrivatecaCaPool#maximum_lifetime}
	MaximumLifetime *string `field:"optional" json:"maximumLifetime" yaml:"maximumLifetime"`
}

