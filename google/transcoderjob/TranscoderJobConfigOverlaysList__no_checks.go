// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package transcoderjob

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TranscoderJobConfigOverlaysList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTranscoderJobConfigOverlaysListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

