// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package transcoderjob

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TranscoderJobConfigManifestsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TranscoderJobConfigManifestsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TranscoderJobConfigManifestsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigManifestsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigManifestsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigManifestsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigManifestsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTranscoderJobConfigManifestsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

