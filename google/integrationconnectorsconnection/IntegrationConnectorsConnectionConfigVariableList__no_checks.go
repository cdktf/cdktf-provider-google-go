// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package integrationconnectorsconnection

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IntegrationConnectorsConnectionConfigVariableList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IntegrationConnectorsConnectionConfigVariableList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIntegrationConnectorsConnectionConfigVariableListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

