// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package computeroute

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComputeRouteAsPathsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ComputeRouteAsPathsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ComputeRouteAsPathsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ComputeRouteAsPathsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ComputeRouteAsPathsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ComputeRouteAsPathsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewComputeRouteAsPathsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

