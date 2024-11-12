// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package memorystoreinstance

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MemorystoreInstanceDiscoveryEndpointsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MemorystoreInstanceDiscoveryEndpointsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MemorystoreInstanceDiscoveryEndpointsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MemorystoreInstanceDiscoveryEndpointsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MemorystoreInstanceDiscoveryEndpointsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MemorystoreInstanceDiscoveryEndpointsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMemorystoreInstanceDiscoveryEndpointsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

