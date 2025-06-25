// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apihubplugininstance

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApihubPluginInstanceActionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApihubPluginInstanceActionsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApihubPluginInstanceActionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApihubPluginInstanceActionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApihubPluginInstanceActionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApihubPluginInstanceActionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApihubPluginInstanceActionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApihubPluginInstanceActionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

