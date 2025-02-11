// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package memorystoreinstance

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MemorystoreInstanceEndpointsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MemorystoreInstanceEndpointsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MemorystoreInstanceEndpointsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MemorystoreInstanceEndpointsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MemorystoreInstanceEndpointsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MemorystoreInstanceEndpointsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMemorystoreInstanceEndpointsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

