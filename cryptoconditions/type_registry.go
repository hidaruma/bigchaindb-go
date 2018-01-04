package cryptoconditions

import (
		"math"
		"log"
)

const (
		"MAXSAFEINTERGERJS" int = int(math.Pow(float64(2),  float64(53))) - 1
)

type RegisterdType map[string]interface{}

type TypeRegistry struct {
	RegisteredTypes []RegisterdType
}

func (tr *TypeRegistry) FindByTypeID(typeID int) RegisterdType {
	if typeID > MAXSAFEINTERGERJS {
		log.Println(UnsupportedError() + "%d", typeID)
		return nil
	}
	for _, registeredType := range tr.RegisteredTypes {
		if typeID == registeredType["type_id"].(int) {
			return registeredType
		}
	}
	log.Println(UnsupportedError() + "%d", typeID)
	return nil
}

func (tr *TypeRegistry) FinfByName(name string) RegisterdType {
	for _, registeredType := range tr.RegisteredTypes {
		if name == registeredType["name"].(string) {
			return registeredType
		}
	}
	log.Println(UnsupportedError() +"%s", name)
	return nil
}

func (tr *TypeRegistry) FindByAsn1Type(asn1Type string) RegisterdType{
	for _, registeredType := range tr.RegisteredTypes {
		if asn1Type == registeredType["asn1"] {
			return registeredType
		}
	}
	log.Println(UnsupportedError())
	return nil
}

func (tr *TypeRegistry) FindByAsn1ConditionType(asn1Type string) RegisterdType {
	for _, registeredType := range tr.RegisteredTypes {
		if asn1Type == registeredType["asn1_condition"] {
			return registeredType
		}
	}
	log.Println(UnsupportedError())
	return nil
}

func (tr *TypeRegistry) FndByAsn1FulfillmentType(asn1Type string) RegisterdType {
	for _, registeredType := range tr.RegisteredTypes {
		if asn1Type == registeredType["asn1_condition"] {
			return registeredType
		}
	}
	log.Println(UnsupportedError())
	return nil
}

func (tr *TypeRegistry) RegisterType() {
	var registerType RegisterdType
	registerType["type_id"] = tr.TYPEID()
	registerType["name"] = tr.TYPENAME()
	registerType["asn1"] = tr.TYPEASN1()
	registerType["asn1_condition"] = tr.TYPEASN1CONDITION()
	registerType["ans1_fulfillment"] = tr.TYPEASN1FULFILLMENT()

	tr.RegisteredTypes = append(tr.RegisteredTypes, registerType)
}