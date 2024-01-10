// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package tpunode

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TpuNodeNetworkEndpointsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TpuNodeNetworkEndpointsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TpuNodeNetworkEndpointsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TpuNodeNetworkEndpointsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TpuNodeNetworkEndpointsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TpuNodeNetworkEndpointsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTpuNodeNetworkEndpointsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

