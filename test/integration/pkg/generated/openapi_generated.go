//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by openapi-gen.go. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package generated

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
	custom "k8s.io/kube-openapi/test/integration/testdata/custom"
	defaults "k8s.io/kube-openapi/test/integration/testdata/defaults"
	enumtype "k8s.io/kube-openapi/test/integration/testdata/enumtype"
	valuevalidation "k8s.io/kube-openapi/test/integration/testdata/valuevalidation"
	ptr "k8s.io/utils/ptr"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"k8s.io/kube-openapi/test/integration/testdata/custom.Bac":                          common.EmbedOpenAPIDefinitionIntoV2Extension(custom.Bac{}.OpenAPIV3Definition(), custom.Bac{}.OpenAPIDefinition()),
		"k8s.io/kube-openapi/test/integration/testdata/custom.Bah":                          schema_test_integration_testdata_custom_Bah(ref),
		"k8s.io/kube-openapi/test/integration/testdata/custom.Bak":                          custom.Bak{}.OpenAPIDefinition(),
		"k8s.io/kube-openapi/test/integration/testdata/custom.Bal":                          custom.Bal{}.OpenAPIV3Definition(),
		"k8s.io/kube-openapi/test/integration/testdata/custom.FooV3OneOf":                   schema_test_integration_testdata_custom_FooV3OneOf(ref),
		"k8s.io/kube-openapi/test/integration/testdata/defaults.Defaulted":                  schema_test_integration_testdata_defaults_Defaulted(ref),
		"k8s.io/kube-openapi/test/integration/testdata/defaults.SubStruct":                  schema_test_integration_testdata_defaults_SubStruct(ref),
		"k8s.io/kube-openapi/test/integration/testdata/dummytype.Bar":                       schema_test_integration_testdata_dummytype_Bar(ref),
		"k8s.io/kube-openapi/test/integration/testdata/dummytype.Baz":                       schema_test_integration_testdata_dummytype_Baz(ref),
		"k8s.io/kube-openapi/test/integration/testdata/dummytype.Foo":                       schema_test_integration_testdata_dummytype_Foo(ref),
		"k8s.io/kube-openapi/test/integration/testdata/dummytype.StatusError":               schema_test_integration_testdata_dummytype_StatusError(ref),
		"k8s.io/kube-openapi/test/integration/testdata/dummytype.Waldo":                     schema_test_integration_testdata_dummytype_Waldo(ref),
		"k8s.io/kube-openapi/test/integration/testdata/enumtype.FruitsBasket":               schema_test_integration_testdata_enumtype_FruitsBasket(ref),
		"k8s.io/kube-openapi/test/integration/testdata/listtype.AtomicList":                 schema_test_integration_testdata_listtype_AtomicList(ref),
		"k8s.io/kube-openapi/test/integration/testdata/listtype.Item":                       schema_test_integration_testdata_listtype_Item(ref),
		"k8s.io/kube-openapi/test/integration/testdata/listtype.MapList":                    schema_test_integration_testdata_listtype_MapList(ref),
		"k8s.io/kube-openapi/test/integration/testdata/listtype.SetList":                    schema_test_integration_testdata_listtype_SetList(ref),
		"k8s.io/kube-openapi/test/integration/testdata/listtype.UntypedList":                schema_test_integration_testdata_listtype_UntypedList(ref),
		"k8s.io/kube-openapi/test/integration/testdata/maptype.AtomicMap":                   schema_test_integration_testdata_maptype_AtomicMap(ref),
		"k8s.io/kube-openapi/test/integration/testdata/maptype.GranularMap":                 schema_test_integration_testdata_maptype_GranularMap(ref),
		"k8s.io/kube-openapi/test/integration/testdata/structtype.AtomicStruct":             schema_test_integration_testdata_structtype_AtomicStruct(ref),
		"k8s.io/kube-openapi/test/integration/testdata/structtype.ContainedStruct":          schema_test_integration_testdata_structtype_ContainedStruct(ref),
		"k8s.io/kube-openapi/test/integration/testdata/structtype.DeclaredAtomicStruct":     schema_test_integration_testdata_structtype_DeclaredAtomicStruct(ref),
		"k8s.io/kube-openapi/test/integration/testdata/structtype.FieldLevelOverrideStruct": schema_test_integration_testdata_structtype_FieldLevelOverrideStruct(ref),
		"k8s.io/kube-openapi/test/integration/testdata/structtype.GranularStruct":           schema_test_integration_testdata_structtype_GranularStruct(ref),
		"k8s.io/kube-openapi/test/integration/testdata/uniontype.InlinedUnion":              schema_test_integration_testdata_uniontype_InlinedUnion(ref),
		"k8s.io/kube-openapi/test/integration/testdata/uniontype.TopLevelUnion":             schema_test_integration_testdata_uniontype_TopLevelUnion(ref),
		"k8s.io/kube-openapi/test/integration/testdata/uniontype.Union":                     schema_test_integration_testdata_uniontype_Union(ref),
		"k8s.io/kube-openapi/test/integration/testdata/uniontype.Union2":                    schema_test_integration_testdata_uniontype_Union2(ref),
		"k8s.io/kube-openapi/test/integration/testdata/valuevalidation.Foo":                 schema_test_integration_testdata_valuevalidation_Foo(ref),
		"k8s.io/kube-openapi/test/integration/testdata/valuevalidation.Foo2":                schema_test_integration_testdata_valuevalidation_Foo2(ref),
		"k8s.io/kube-openapi/test/integration/testdata/valuevalidation.Foo3":                schema_test_integration_testdata_valuevalidation_Foo3(ref),
		"k8s.io/kube-openapi/test/integration/testdata/valuevalidation.Foo4":                valuevalidation.Foo4{}.OpenAPIDefinition(),
		"k8s.io/kube-openapi/test/integration/testdata/valuevalidation.Foo5":                schema_test_integration_testdata_valuevalidation_Foo5(ref),
	}
}

func schema_test_integration_testdata_custom_Bah(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.EmbedOpenAPIDefinitionIntoV2Extension(custom.Bah{}.OpenAPIV3Definition(), common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:   custom.Bah{}.OpenAPISchemaType(),
				Format: custom.Bah{}.OpenAPISchemaFormat(),
			},
		},
	})
}

func schema_test_integration_testdata_custom_FooV3OneOf(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.EmbedOpenAPIDefinitionIntoV2Extension(common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FooV3OneOf has an OpenAPIV3OneOfTypes method",
				OneOf:       common.GenerateOpenAPIV3OneOfSchema(custom.FooV3OneOf{}.OpenAPIV3OneOfTypes()),
				Format:      custom.FooV3OneOf{}.OpenAPISchemaFormat(),
			},
		},
	}, common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FooV3OneOf has an OpenAPIV3OneOfTypes method",
				Type:        custom.FooV3OneOf{}.OpenAPISchemaType(),
				Format:      custom.FooV3OneOf{}.OpenAPISchemaFormat(),
			},
		},
	})
}

func schema_test_integration_testdata_defaults_Defaulted(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"Field": {
						SchemaProps: spec.SchemaProps{
							Default: "bar",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"OtherField": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
					"List": {
						SchemaProps: spec.SchemaProps{
							Default: []interface{}{"foo", "bar"},
							Type:    []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"Sub": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{"i": 5, "s": "foo"},
							Ref:     ref("k8s.io/kube-openapi/test/integration/testdata/defaults.SubStruct"),
						},
					},
					"OtherSub": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/kube-openapi/test/integration/testdata/defaults.SubStruct"),
						},
					},
					"Map": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{"foo": "bar"},
							Type:    []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"localSymbolReference": {
						SchemaProps: spec.SchemaProps{
							Default: defaults.ConstantValue,
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"fullyQualifiedSymbolReference": {
						SchemaProps: spec.SchemaProps{
							Default: defaults.ConstantValue,
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"externalSymbolReference": {
						SchemaProps: spec.SchemaProps{
							Default: enumtype.FruitApple,
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"pointerConversionSymbolReference": {
						SchemaProps: spec.SchemaProps{
							Default: enumtype.FruitApple,
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"defaultedAliasSymbolReference": {
						SchemaProps: spec.SchemaProps{
							Default: defaults.ConstantValue,
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"OtherField", "List", "Sub", "OtherSub", "Map"},
			},
		},
		Dependencies: []string{
			"k8s.io/kube-openapi/test/integration/testdata/defaults.SubStruct"},
	}
}

func schema_test_integration_testdata_defaults_SubStruct(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"S": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"I": {
						SchemaProps: spec.SchemaProps{
							Default: 1,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
				},
				Required: []string{"S"},
			},
		},
	}
}

func schema_test_integration_testdata_dummytype_Bar(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"ViolationBehind": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"Violation": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
				},
				Required: []string{"ViolationBehind", "Violation"},
			},
		},
	}
}

func schema_test_integration_testdata_dummytype_Baz(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"Violation": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"ViolationBehind": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
				},
				Required: []string{"Violation", "ViolationBehind"},
			},
		},
	}
}

func schema_test_integration_testdata_dummytype_Foo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"Second": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"First": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
				},
				Required: []string{"Second", "First"},
			},
		},
	}
}

func schema_test_integration_testdata_dummytype_StatusError(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"Code": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
					"Message": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"Code", "Message"},
			},
		},
	}
}

func schema_test_integration_testdata_dummytype_Waldo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"First": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
					"Second": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"First", "Second"},
			},
		},
	}
}

func schema_test_integration_testdata_enumtype_FruitsBasket(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FruitsBasket is the type that contains the enum type.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"content": {
						SchemaProps: spec.SchemaProps{
							Description: "Possible enum values:\n - `\"apple\"` is the Apple\n - `\"banana\"` is the Banana\n - `\"onigiri\"` is the Rice ball that does not seem to belong to a fruits basket but has a long comment that is so long that it spans multiple lines",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
							Enum:        []interface{}{"apple", "banana", "onigiri"},
						},
					},
					"count": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
				},
				Required: []string{"content", "count"},
			},
		},
	}
}

func schema_test_integration_testdata_listtype_AtomicList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"Field": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"Field"},
			},
		},
	}
}

func schema_test_integration_testdata_listtype_Item(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"Protocol": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"Port": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
					"a": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
					"b": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"c": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
				},
				Required: []string{"Protocol", "Port"},
			},
		},
	}
}

func schema_test_integration_testdata_listtype_MapList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"Field": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"port",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/kube-openapi/test/integration/testdata/listtype.Item"),
									},
								},
							},
						},
					},
				},
				Required: []string{"Field"},
			},
		},
		Dependencies: []string{
			"k8s.io/kube-openapi/test/integration/testdata/listtype.Item"},
	}
}

func schema_test_integration_testdata_listtype_SetList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"Field": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"Field"},
			},
		},
	}
}

func schema_test_integration_testdata_listtype_UntypedList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"Field": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"Field"},
			},
		},
	}
}

func schema_test_integration_testdata_maptype_AtomicMap(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"KeyValue": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-map-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"KeyValue"},
			},
		},
	}
}

func schema_test_integration_testdata_maptype_GranularMap(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"KeyValue": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-map-type": "granular",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"KeyValue"},
			},
		},
	}
}

func schema_test_integration_testdata_structtype_AtomicStruct(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"Field": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-map-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/kube-openapi/test/integration/testdata/structtype.ContainedStruct"),
						},
					},
					"OtherField": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
				},
				Required: []string{"Field", "OtherField"},
			},
		},
		Dependencies: []string{
			"k8s.io/kube-openapi/test/integration/testdata/structtype.ContainedStruct"},
	}
}

func schema_test_integration_testdata_structtype_ContainedStruct(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
			},
		},
	}
}

func schema_test_integration_testdata_structtype_DeclaredAtomicStruct(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"Field": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
				},
				Required: []string{"Field"},
			},
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					"x-kubernetes-map-type": "atomic",
				},
			},
		},
	}
}

func schema_test_integration_testdata_structtype_FieldLevelOverrideStruct(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"Field": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-map-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/kube-openapi/test/integration/testdata/structtype.DeclaredAtomicStruct"),
						},
					},
					"OtherField": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
				},
				Required: []string{"Field", "OtherField"},
			},
		},
		Dependencies: []string{
			"k8s.io/kube-openapi/test/integration/testdata/structtype.DeclaredAtomicStruct"},
	}
}

func schema_test_integration_testdata_structtype_GranularStruct(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"Field": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-map-type": "granular",
							},
						},
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/kube-openapi/test/integration/testdata/structtype.ContainedStruct"),
						},
					},
					"OtherField": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int32",
						},
					},
				},
				Required: []string{"Field", "OtherField"},
			},
		},
		Dependencies: []string{
			"k8s.io/kube-openapi/test/integration/testdata/structtype.ContainedStruct"},
	}
}

func schema_test_integration_testdata_uniontype_InlinedUnion(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"field1": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"field2": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"unionType": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"fieldA": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"fieldB": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"type": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"alpha": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"beta": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
				Required: []string{"name", "type"},
			},
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					"x-kubernetes-unions": []interface{}{
						map[string]interface{}{
							"discriminator": "unionType",
							"fields-to-discriminateBy": map[string]interface{}{
								"fieldA": "FieldA",
								"fieldB": "FieldB",
							},
						},
						map[string]interface{}{
							"discriminator": "type",
							"fields-to-discriminateBy": map[string]interface{}{
								"alpha": "Alpha",
								"beta":  "Beta",
							},
						},
						map[string]interface{}{
							"fields-to-discriminateBy": map[string]interface{}{
								"field1": "Field1",
								"field2": "Field2",
							},
						},
					},
				},
			},
		},
	}
}

func schema_test_integration_testdata_uniontype_TopLevelUnion(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"unionType": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"fieldA": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"fieldB": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
				Required: []string{"name"},
			},
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					"x-kubernetes-unions": []interface{}{
						map[string]interface{}{
							"discriminator": "unionType",
							"fields-to-discriminateBy": map[string]interface{}{
								"fieldA": "FieldA",
								"fieldB": "FieldB",
							},
						},
					},
				},
			},
		},
	}
}

func schema_test_integration_testdata_uniontype_Union(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"unionType": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"fieldA": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"fieldB": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
			},
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					"x-kubernetes-unions": []interface{}{
						map[string]interface{}{
							"discriminator": "unionType",
							"fields-to-discriminateBy": map[string]interface{}{
								"fieldA": "FieldA",
								"fieldB": "FieldB",
							},
						},
					},
				},
			},
		},
	}
}

func schema_test_integration_testdata_uniontype_Union2(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"alpha": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"beta": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
				Required: []string{"type"},
			},
			VendorExtensible: spec.VendorExtensible{
				Extensions: spec.Extensions{
					"x-kubernetes-unions": []interface{}{
						map[string]interface{}{
							"discriminator": "type",
							"fields-to-discriminateBy": map[string]interface{}{
								"alpha": "Alpha",
								"beta":  "Beta",
							},
						},
					},
				},
			},
		},
	}
}

func schema_test_integration_testdata_valuevalidation_Foo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:          []string{"object"},
				MinProperties: ptr.To[int64](1),
				MaxProperties: ptr.To[int64](5),
				Properties: map[string]spec.Schema{
					"StringValue": {
						SchemaProps: spec.SchemaProps{
							Default:   "",
							MinLength: ptr.To[int64](1),
							MaxLength: ptr.To[int64](5),
							Pattern:   "^a.*b$",
							Type:      []string{"string"},
							Format:    "",
						},
					},
					"NumberValue": {
						SchemaProps: spec.SchemaProps{
							Default:          0,
							Minimum:          ptr.To[float64](1),
							Maximum:          ptr.To[float64](5),
							ExclusiveMinimum: true,
							ExclusiveMaximum: true,
							MultipleOf:       ptr.To[float64](2),
							Type:             []string{"number"},
							Format:           "double",
						},
					},
					"ArrayValue": {
						SchemaProps: spec.SchemaProps{
							MinItems:    ptr.To[int64](1),
							MaxItems:    ptr.To[int64](5),
							UniqueItems: true,
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"MapValue": {
						SchemaProps: spec.SchemaProps{
							MinProperties: ptr.To[int64](1),
							MaxProperties: ptr.To[int64](5),
							Type:          []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"StringValue", "NumberValue", "ArrayValue", "MapValue"},
			},
		},
	}
}

func schema_test_integration_testdata_valuevalidation_Foo2(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.EmbedOpenAPIDefinitionIntoV2Extension(valuevalidation.Foo2{}.OpenAPIV3Definition(), common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description:   "This one has an open API v3 definition",
				Type:          valuevalidation.Foo2{}.OpenAPISchemaType(),
				Format:        valuevalidation.Foo2{}.OpenAPISchemaFormat(),
				MaxProperties: ptr.To[int64](5),
			},
		},
	})
}

func schema_test_integration_testdata_valuevalidation_Foo3(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.EmbedOpenAPIDefinitionIntoV2Extension(common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description:   "This one has a OneOf",
				OneOf:         common.GenerateOpenAPIV3OneOfSchema(valuevalidation.Foo3{}.OpenAPIV3OneOfTypes()),
				Format:        valuevalidation.Foo3{}.OpenAPISchemaFormat(),
				MaxProperties: ptr.To[int64](5),
			},
		},
	}, common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description:   "This one has a OneOf",
				Type:          valuevalidation.Foo3{}.OpenAPISchemaType(),
				Format:        valuevalidation.Foo3{}.OpenAPISchemaFormat(),
				MaxProperties: ptr.To[int64](5),
			},
		},
	})
}

func schema_test_integration_testdata_valuevalidation_Foo5(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.EmbedOpenAPIDefinitionIntoV2Extension(valuevalidation.Foo5{}.OpenAPIV3Definition(), common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:          valuevalidation.Foo5{}.OpenAPISchemaType(),
				Format:        valuevalidation.Foo5{}.OpenAPISchemaFormat(),
				MinProperties: ptr.To[int64](1),
				MaxProperties: ptr.To[int64](5),
			},
		},
	})
}
