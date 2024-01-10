// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package securesourcemanagerinstance

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecureSourceManagerInstanceHostConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SecureSourceManagerInstanceHostConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecureSourceManagerInstanceHostConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecureSourceManagerInstanceHostConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecureSourceManagerInstanceHostConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecureSourceManagerInstanceHostConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecureSourceManagerInstanceHostConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

