// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apihubplugin

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApihubPluginActionsConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApihubPluginActionsConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApihubPluginActionsConfigList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApihubPluginActionsConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApihubPluginActionsConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApihubPluginActionsConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApihubPluginActionsConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApihubPluginActionsConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

