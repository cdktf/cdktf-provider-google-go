// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package integrationconnectorsconnection

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegrationConnectorsConnectionStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IntegrationConnectorsConnectionStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IntegrationConnectorsConnectionStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IntegrationConnectorsConnectionStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IntegrationConnectorsConnectionStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIntegrationConnectorsConnectionStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

