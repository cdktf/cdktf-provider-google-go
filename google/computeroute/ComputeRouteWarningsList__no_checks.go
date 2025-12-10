// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package computeroute

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComputeRouteWarningsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ComputeRouteWarningsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ComputeRouteWarningsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ComputeRouteWarningsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ComputeRouteWarningsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ComputeRouteWarningsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewComputeRouteWarningsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

