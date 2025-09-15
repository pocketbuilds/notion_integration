package notion

import (
	"reflect"
)

type Property interface {
	GetId() string
	GetType() string
	GetValue() any
	SetValue(v any)
	Editable() bool
}

var propertyRegistry = map[string]Property{}

func RegisterProperty(p Property) {
	propertyRegistry[p.GetType()] = p
}

func NewProperty(propertyType string) Property {
	p, ok := propertyRegistry[propertyType]
	if !ok {
		return nil
	}
	return reflect.New(reflect.TypeOf(p).Elem()).Interface().(Property)
}

type PropertyConfig interface {
	GetType() string
	GetId() string
}

var propertyConfigRegistry = map[string]PropertyConfig{}

func RegisterPropertyConfig(pc PropertyConfig) {
	propertyConfigRegistry[pc.GetType()] = pc
}

func NewPropertyConfig(propertyType string) PropertyConfig {
	p, ok := propertyConfigRegistry[propertyType]
	if !ok {
		return nil
	}
	return reflect.New(reflect.TypeOf(p).Elem()).Interface().(PropertyConfig)
}
