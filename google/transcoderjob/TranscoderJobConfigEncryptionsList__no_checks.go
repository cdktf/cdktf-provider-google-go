// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package transcoderjob

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TranscoderJobConfigEncryptionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTranscoderJobConfigEncryptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

