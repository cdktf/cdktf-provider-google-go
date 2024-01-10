// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package kmscryptokey

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KmsCryptoKeyPrimaryList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KmsCryptoKeyPrimaryList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KmsCryptoKeyPrimaryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KmsCryptoKeyPrimaryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KmsCryptoKeyPrimaryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KmsCryptoKeyPrimaryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKmsCryptoKeyPrimaryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

