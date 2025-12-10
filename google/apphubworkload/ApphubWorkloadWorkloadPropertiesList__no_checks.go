// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apphubworkload

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApphubWorkloadWorkloadPropertiesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApphubWorkloadWorkloadPropertiesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApphubWorkloadWorkloadPropertiesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApphubWorkloadWorkloadPropertiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApphubWorkloadWorkloadPropertiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApphubWorkloadWorkloadPropertiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApphubWorkloadWorkloadPropertiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

