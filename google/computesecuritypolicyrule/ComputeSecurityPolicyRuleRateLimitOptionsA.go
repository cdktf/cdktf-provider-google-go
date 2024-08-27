// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicyrule


type ComputeSecurityPolicyRuleRateLimitOptionsA struct {
	// Can only be specified if the action for the rule is "rate_based_ban".
	//
	// If specified, determines the time (in seconds) the traffic will continue to be banned by the rate limit after the rate falls below the threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/compute_security_policy_rule#ban_duration_sec ComputeSecurityPolicyRuleA#ban_duration_sec}
	BanDurationSec *float64 `field:"optional" json:"banDurationSec" yaml:"banDurationSec"`
	// ban_threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/compute_security_policy_rule#ban_threshold ComputeSecurityPolicyRuleA#ban_threshold}
	BanThreshold *ComputeSecurityPolicyRuleRateLimitOptionsBanThresholdA `field:"optional" json:"banThreshold" yaml:"banThreshold"`
	// Action to take for requests that are under the configured rate limit threshold. Valid option is "allow" only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/compute_security_policy_rule#conform_action ComputeSecurityPolicyRuleA#conform_action}
	ConformAction *string `field:"optional" json:"conformAction" yaml:"conformAction"`
	// Determines the key to enforce the rateLimitThreshold on.
	//
	// Possible values are:
	// * ALL: A single rate limit threshold is applied to all the requests matching this rule. This is the default value if "enforceOnKey" is not configured.
	// * IP: The source IP address of the request is the key. Each IP has this limit enforced separately.
	// * HTTP_HEADER: The value of the HTTP header whose name is configured under "enforceOnKeyName". The key value is truncated to the first 128 bytes of the header value. If no such header is present in the request, the key type defaults to ALL.
	// * XFF_IP: The first IP address (i.e. the originating client IP address) specified in the list of IPs under X-Forwarded-For HTTP header. If no such header is present or the value is not a valid IP, the key defaults to the source IP address of the request i.e. key type IP.
	// * HTTP_COOKIE: The value of the HTTP cookie whose name is configured under "enforceOnKeyName". The key value is truncated to the first 128 bytes of the cookie value. If no such cookie is present in the request, the key type defaults to ALL.
	// * HTTP_PATH: The URL path of the HTTP request. The key value is truncated to the first 128 bytes.
	// * SNI: Server name indication in the TLS session of the HTTPS request. The key value is truncated to the first 128 bytes. The key type defaults to ALL on a HTTP session.
	// * REGION_CODE: The country/region from which the request originates.
	// * TLS_JA3_FINGERPRINT: JA3 TLS/SSL fingerprint if the client connects using HTTPS, HTTP/2 or HTTP/3. If not available, the key type defaults to ALL.
	// * USER_IP: The IP address of the originating client, which is resolved based on "userIpRequestHeaders" configured with the security policy. If there is no "userIpRequestHeaders" configuration or an IP address cannot be resolved from it, the key type defaults to IP. Possible values: ["ALL", "IP", "HTTP_HEADER", "XFF_IP", "HTTP_COOKIE", "HTTP_PATH", "SNI", "REGION_CODE", "TLS_JA3_FINGERPRINT", "USER_IP"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/compute_security_policy_rule#enforce_on_key ComputeSecurityPolicyRuleA#enforce_on_key}
	EnforceOnKey *string `field:"optional" json:"enforceOnKey" yaml:"enforceOnKey"`
	// enforce_on_key_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/compute_security_policy_rule#enforce_on_key_configs ComputeSecurityPolicyRuleA#enforce_on_key_configs}
	EnforceOnKeyConfigs interface{} `field:"optional" json:"enforceOnKeyConfigs" yaml:"enforceOnKeyConfigs"`
	// Rate limit key name applicable only for the following key types: HTTP_HEADER -- Name of the HTTP header whose value is taken as the key value.
	//
	// HTTP_COOKIE -- Name of the HTTP cookie whose value is taken as the key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/compute_security_policy_rule#enforce_on_key_name ComputeSecurityPolicyRuleA#enforce_on_key_name}
	EnforceOnKeyName *string `field:"optional" json:"enforceOnKeyName" yaml:"enforceOnKeyName"`
	// Action to take for requests that are above the configured rate limit threshold, to either deny with a specified HTTP response code, or redirect to a different endpoint.
	//
	// Valid options are deny(STATUS), where valid values for STATUS are 403, 404, 429, and 502.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/compute_security_policy_rule#exceed_action ComputeSecurityPolicyRuleA#exceed_action}
	ExceedAction *string `field:"optional" json:"exceedAction" yaml:"exceedAction"`
	// exceed_redirect_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/compute_security_policy_rule#exceed_redirect_options ComputeSecurityPolicyRuleA#exceed_redirect_options}
	ExceedRedirectOptions *ComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsA `field:"optional" json:"exceedRedirectOptions" yaml:"exceedRedirectOptions"`
	// rate_limit_threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/compute_security_policy_rule#rate_limit_threshold ComputeSecurityPolicyRuleA#rate_limit_threshold}
	RateLimitThreshold *ComputeSecurityPolicyRuleRateLimitOptionsRateLimitThresholdA `field:"optional" json:"rateLimitThreshold" yaml:"rateLimitThreshold"`
}

