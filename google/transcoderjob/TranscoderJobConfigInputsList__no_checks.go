// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package transcoderjob

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TranscoderJobConfigInputsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TranscoderJobConfigInputsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TranscoderJobConfigInputsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigInputsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigInputsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigInputsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigInputsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTranscoderJobConfigInputsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

