package amino

import (
	"fmt"
	"reflect"
)

type TypesJson struct {
	ConcreteTypes  []ConcreteJson  `json:"concrete_types"`
	InterfaceTypes []InterfaceJson `json:"interface_types"`
}

type ConcreteJson struct {
	*TypeJson
	Fields []StructFieldJson
}

type InterfaceJson struct {
	Name              string   `json:"name"`
	ImplementingTypes []string `json:"implementing_types"`
}

type TypeJson struct {
	Name string       `json:"name"`
	Kind reflect.Kind `json:"kind"`
	Key  *TypeJson    `json:"key"`
	Elem *TypeJson    `json:"elem"`
	Len  int          `json:"len"`
}

type StructFieldJson struct {
	Name             string   `json:"name"`
	ProtoFieldNumber uint32   `json:"proto_field_number"`
	Type             TypeJson `json:"type"`
}

type reflector struct {
	aminoNames map[string]string
}

func typeKey(typ reflect.Type) string  {
	return fmt.Sprintf("%s/%s", typ.PkgPath(), typ.Name())
}

func (r *reflector) findAminoName(typ reflect.Type) string {
	return r.aminoNames[typeKey(typ)]
}

func (r *reflector) storeAminoName(aminoName string, ti TypeInfo) {
	r.aminoNames[typeKey(ti.Type)] = aminoName
}

func (r *reflector) typeToJson(typ reflect.Type) *TypeJson {
	kind := typ.Kind()
	tj := &TypeJson{
		Name: r.findAminoName(typ),
		Kind: kind,
	}
	if kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map {
		tj.Elem = r.typeToJson(typ)
	}
	if kind == reflect.Array {
		tj.Len = typ.Len()
	}
	if kind == reflect.Map {
		tj.Key = r.typeToJson(typ)
	}
	return tj
}

func (r *reflector) concreteToJson(ti TypeInfo) ConcreteJson {
	cj := ConcreteJson{TypeJson: r.typeToJson(ti.Type)}
	name := ti.Name
	cj.Name = name
	if cj.Kind == reflect.Struct {

	}
	return cj
}

func (cdc *Codec) DumpTypes() {
	// register concrete amino names first!
	// dump types
	// dump interfaces
}
