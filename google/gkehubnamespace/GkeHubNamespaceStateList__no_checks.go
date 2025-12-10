// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gkehubnamespace

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GkeHubNamespaceStateList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GkeHubNamespaceStateList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GkeHubNamespaceStateList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GkeHubNamespaceStateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GkeHubNamespaceStateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GkeHubNamespaceStateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGkeHubNamespaceStateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

