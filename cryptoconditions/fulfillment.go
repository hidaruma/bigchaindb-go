package cryptoconditions

import (
		"encoding/base64"
		"log"
		"encoding/asn1"
)

type Fulfillment struct {

}

func (f *Fulfillment) TYPEID() int {
	return
}

func (f *Fulfillment) TYPENAME() string {
	return ""
}

func (f *Fulfillment) FromURI(selializedFulfillment string) {
	if serializedFulfillment.(type) != string {
		log.Println(exceptions.TypeError())
	}
	return
}

func (f *Fulfillment) FromBinary(data []byte) {

}

func (f *Fulfillment) FromAsn1Dict(asn1Dict ) {

}

func (f *Fulfillment) FromDict(data []byte) {

}

func (f *Fulfillment) FromJson(data []byte) {

}

func (f *Fulfillment) TypeID() int {
	return f.TYPEID
}

func (f *Fulfillment) TypeName() string {
	return f.TYPENAME()
}

func (f *Fulfillment) Subtypes() []interface{} {
	return []interface{}
}

func (f *Fulfillment) Condition() *Condition {
	var condition *Condition
	condition.TypeID = f.TypeID
	condition.hash = f.GenerateHash()
	condition.cost = f.CalculateCost()
	condition.subtypes = f.Subtypes()
	return condition
}

func (f *Fulfillment) ConditionURI() string {
	return f.Condition.SerializeURI()
}

func (f *Fulfillment) ConditionBinary() []byte{
	return f.Condition.SerializeBinary()
}

func (f *Fulfillment) GenerateHash() []byte {

}

func (f *Fulfillment) CalculateCost() int {

}

func (f *Fulfillment) SerializeURI() string {
	var url string
	url, err := base64.RawURLEncoding.DecodeString(f.serializeBinary())
	if err != nil {
		log.Println(err)
	}
	return url
}

func (f *Fulfillment) SerializeBinary() []byte {
	var asn1Dict = map[string]string
	asn1Dict = {f.TYPEASN1: f.Asn1DictPayload}
	var asn1
	asn1 = natDecode(asn1Dict, Asn1Fulfillment())
	var binObj []byte
	binObj, err := derEnecode(asn1)
	if err != nil {
		log.Println(err)
	}
	return binObj
}


func (f *Fulfillment) Asn1Dict(data map[string]interface{}) {

}

func (f *Fulfillment) ParseDict

func (f *Fulfillment) ParseAsn1DictPayload(data map[string]interface{}) {

}

func (f *Fulfillment) Validate(args []int, kwargs []string) bool {

}