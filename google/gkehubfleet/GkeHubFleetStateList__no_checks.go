// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gkehubfleet

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GkeHubFleetStateList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GkeHubFleetStateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GkeHubFleetStateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GkeHubFleetStateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GkeHubFleetStateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGkeHubFleetStateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

