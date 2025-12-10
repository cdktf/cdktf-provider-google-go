// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apigeeapi

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApigeeApiMetaDataList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApigeeApiMetaDataList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApigeeApiMetaDataList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApigeeApiMetaDataList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApigeeApiMetaDataList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApigeeApiMetaDataList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApigeeApiMetaDataListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

