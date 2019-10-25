package amino

import "reflect"

type TypesJson struct {
	ConcreteTypes  []TypeJson `json:"concrete_types"`
	InterfaceTypes []TypeJson `json:"interface_types"`
}

type TypeJson struct {
	// GoName corresponds to the reflect.Type Name()
	GoName string `json:"go_name"`
	// AminoName corresponds to the name passed to RegisterConcrete for concrete types
	AminoName string       `json:"amino_name"`
	PkgPath   string       `json:"pkg_path"`
	String    string       `json:"string"`
	Kind      reflect.Kind `json:"kind"`
	// ImplementingTypes lists the amino names of concrete types implementing
	// this interface for interface types
	//ImplementingTypes []string          `json:"implementing_types"`
	Fields []StructFieldJson `json:"fields"`
	Key    TypeJson          `json:"key"`
	Elem   TypeJson          `json:"elem"`
	Len    int               `json:"len"`
}

type StructFieldJson struct {
	Name             string `json:"name"`
	ProtoFieldNumber uint32 `json:"proto_field_number"`
	Type             TypeJson
}

type reflector struct {
}

func (r *reflector) findAminoName(typ reflect.Type) string {
	panic("TODO")
}

func (r *reflector) storeAminoName(aminoName string, typ reflect.Type) {
	panic("TODO")
}

func (r *reflector) typeToJson(aminoName string, typ reflect.Type) TypeJson {
	if aminoName == "" {
		aminoName = r.findAminoName(typ)
	} else {
		r.storeAminoName(aminoName, typ)
	}
	kind := typ.Kind()
	if kind == reflect.Ptr {
		return r.typeToJson(aminoName, typ.Elem())
	}
	tj := TypeJson{
		AminoName: aminoName,
		GoName:    typ.Name(),
		PkgPath:   typ.PkgPath(),
		String:    typ.String(),
		Kind:      kind,
	}
	if kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map {
		tj.Elem = r.typeToJson("", typ.Elem())
	}
	if kind == reflect.Map {
		tj.Key = r.typeToJson("", typ.Key())

	}
	if kind == reflect.Struct {
		n := typ.NumField()
		for i := 0; i < n; i++ {
			f := typ.Field()
			sj := StructFieldJson{Name:}
			tj.Fields = append(tj.Fields, )
		}
	}
	if kind == reflect.Chan || kind == reflect.Func {
		panic("unexpected")
	}
	return tj
}

func (cdc *Codec) DumpTypes() {
}
