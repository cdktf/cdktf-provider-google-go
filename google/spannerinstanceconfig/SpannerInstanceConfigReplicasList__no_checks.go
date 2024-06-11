// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package spannerinstanceconfig

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SpannerInstanceConfigReplicasList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SpannerInstanceConfigReplicasList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SpannerInstanceConfigReplicasList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SpannerInstanceConfigReplicasList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SpannerInstanceConfigReplicasList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SpannerInstanceConfigReplicasList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SpannerInstanceConfigReplicasList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSpannerInstanceConfigReplicasListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

