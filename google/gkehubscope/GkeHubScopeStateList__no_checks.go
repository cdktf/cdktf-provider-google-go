// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gkehubscope

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GkeHubScopeStateList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GkeHubScopeStateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GkeHubScopeStateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GkeHubScopeStateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GkeHubScopeStateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGkeHubScopeStateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

