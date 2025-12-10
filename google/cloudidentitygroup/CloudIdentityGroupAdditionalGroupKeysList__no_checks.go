// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudidentitygroup

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudIdentityGroupAdditionalGroupKeysList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudIdentityGroupAdditionalGroupKeysList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudIdentityGroupAdditionalGroupKeysList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudIdentityGroupAdditionalGroupKeysList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudIdentityGroupAdditionalGroupKeysList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudIdentityGroupAdditionalGroupKeysList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudIdentityGroupAdditionalGroupKeysListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

