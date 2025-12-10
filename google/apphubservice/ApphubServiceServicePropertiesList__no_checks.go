// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apphubservice

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApphubServiceServicePropertiesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApphubServiceServicePropertiesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApphubServiceServicePropertiesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApphubServiceServicePropertiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApphubServiceServicePropertiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApphubServiceServicePropertiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApphubServiceServicePropertiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

