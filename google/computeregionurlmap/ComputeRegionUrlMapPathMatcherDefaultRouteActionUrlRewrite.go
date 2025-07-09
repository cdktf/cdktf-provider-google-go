// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionurlmap


type ComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewrite struct {
	// Prior to forwarding the request to the selected service, the request's host header is replaced with contents of hostRewrite.
	//
	// The value must be between 1 and 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/compute_region_url_map#host_rewrite ComputeRegionUrlMap#host_rewrite}
	HostRewrite *string `field:"optional" json:"hostRewrite" yaml:"hostRewrite"`
	// Prior to forwarding the request to the selected backend service, the matching portion of the request's path is replaced by pathPrefixRewrite.
	//
	// The value must be between 1 and 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/compute_region_url_map#path_prefix_rewrite ComputeRegionUrlMap#path_prefix_rewrite}
	PathPrefixRewrite *string `field:"optional" json:"pathPrefixRewrite" yaml:"pathPrefixRewrite"`
	// If specified, the pattern rewrites the URL path (based on the :path header) using the HTTP template syntax.
	//
	// A corresponding pathTemplateMatch must be specified. Any template variables must exist in the pathTemplateMatch field.
	//
	// * At least one variable must be specified in the pathTemplateMatch field
	// * You can omit variables from the rewritten URL
	// * The * and ** operators cannot be matched unless they have a corresponding variable name - e.g. {format=*} or {var=**}.
	//
	// For example, a pathTemplateMatch of /static/{format=**} could be rewritten as /static/content/{format} to prefix
	// /content to the URL. Variables can also be re-ordered in a rewrite, so that /{country}/{format}/{suffix=**} can be
	// rewritten as /content/{format}/{country}/{suffix}.
	//
	// At least one non-empty routeRules[].matchRules[].path_template_match is required.
	//
	// Only one of pathPrefixRewrite or pathTemplateRewrite may be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/compute_region_url_map#path_template_rewrite ComputeRegionUrlMap#path_template_rewrite}
	PathTemplateRewrite *string `field:"optional" json:"pathTemplateRewrite" yaml:"pathTemplateRewrite"`
}

