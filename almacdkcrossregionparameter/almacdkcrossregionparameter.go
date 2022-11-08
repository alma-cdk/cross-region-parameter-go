package almacdkcrossregionparameter

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@alma-cdk/cross-region-parameter.CrossRegionParameter",
		reflect.TypeOf((*CrossRegionParameter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CrossRegionParameter{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@alma-cdk/cross-region-parameter.CrossRegionParameterProps",
		reflect.TypeOf((*CrossRegionParameterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@alma-cdk/cross-region-parameter.TagProp",
		reflect.TypeOf((*TagProp)(nil)).Elem(),
	)
}
