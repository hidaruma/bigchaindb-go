package cryptoconditions

import (
	"encoding/asn1"
	"strings"
	"regexp"
	"net/url"
	"encoding/base64"
	"strconv"
)

const (
	CONDITIONREGEX string = `^ni:\/\/\/sha-256;([a-zA-Z0-9_-]{0,86})\?(.+)$`
	CONDITIONREGEXSTRICT string = CONDITIONREGEX
	INTEGERREGEX string = `^0|[1-9]\d*$`
	CONDITIONURISCHEME string = `ni`
)

type Condition struct {
	TypeID int
	TypeName string
	hash string
	cost int
	subtypes string
}

type ConditionType map[string]interface{}

func (ct *ConditionType) TYPECATEGORY() string {

}

func (c *Condition) SUPPORTEDSUBTYPES() []string {
	supportedSubtypes := []string{"preimage-sha-256","prefix-sha-256", "threshold-sha-256","rsa-sha-256", "ed25519-sha-256"}
	return supportedSubtypes
}

func (c *Condition) MAXSAFESUBTYPES() int {
	return len(c.SUPPORTEDSUBTYPES())
}

func (c *Condition) MAXCOST() int {
	return 2097152
}

func (c *Condition) REGEX() string {
	return CONDITIONREGEX
}

func (c *Condition) REGEXSTRICT() string {
	return CONDITIONREGEXSTRICT
}

func (c *Condition) eq(other *Condition) bool {
	return c.SerializeBinary() == other.SerializeBinary()
}

func (c *Condition) lt (other *Condition) bool {
	return c.SerializeBinary() < other.SerializeBinary()
}

func (c *Condition) FromURI(serializedCondititon string) *Condition {
	switch serializedCondititon.(type) {
	//	case *Condition:
	//		return serializedCondititon
	case string:
	default:
		log.Println(TypeError())
	}
	var pieces []string
	pieces = strings.Split(serializedCondititon, ":")
	if pieces[0] != CONDITIONURISCHEME {
		log.Println(PrefixError("Serialized condition must start with" + CONDITIONURISCHEME + "."))
	}
	var regexMatch bool
	regexMatch, err := regexp.MatchString(CONDITIONREGEXSTRICT, serializedCondititon)
	if err != nil {
		log.Println(err)
	}
	if !regexMatch {
		log.Println(ParsingError)
	}
	re := regexp.MustCompile(CONDITIONREGEXSTRICT)
	var regexMatchGroup [][]string
	regexMatchGroup = re.FindAllStringSubmatch(serializedCondititon, -1)
	u, err := url.Parse(regexMatchGroup[1][2])
	if err != nil {
		log.Println(err)
	}
	qsDict, _ := url.ParseQuery(u.RawQuery)
	var fingerprintType string
	fingerprintType = qsDict["fpt"][0]
	var conditionType map[string]string
	conditionType = TypeRegistry.FindByName(fingerprintType)
	var cost string
	cost = qsDict["cost"][0]
	if !regexp.MatchString(INTEGERREGEX, cost) {
		log.Println(ParsingError())
	}
	var fingerprint string
	fingerprint = regexMatchGroup[1][1]
	var condition *Condition
	condition.TypeID = conditionType["type_id"]
	if conditionType["class"].TYPECATEGORY == "compound" {
		condition.subtypes.Update(strings.Split(qsDict["subtypes"][0], ","))
	}
	condition.hash, err := base64.URLEncoding.DecodeString(fingerprint)
	if err != nil {
		log.Println(err)
	}
	condition.cost = strconv.Atoi(cost)
	return condition
}

func (c *Condition) FromBinary(data []byte) {
	var asn1ConditionObj map[string]string
	var residue string
	asn1ConditionObj, residue = asn1.Unmarshal
	var asn1ConditionDict map[string]string

	return Condition.FromAsn1Dict(asn1ConditionDict)
}

type Fulfillment struct {

}

func (f *Fulfillment) TYPEID() int {
	return
}

func (f *Fulfillment) TYPENAME() string {
	return ""
}

func (f *Fulfillment) FromURI(selializedFullfilment string) {
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

func (f *Fulfillment) ToDict() {

}

func (f *Fulfillment) Asn1Dict(data [string]string) Fulfillment {

}

func (f *Fulfillment) ParseAsn1DictPayload(data [string]string) {

}

func (f *Fulfillment) Validate(args []int, kwargs []string) bool {

}

type ConditionTypes struct {
	*asn1.BitString
	NamedValues map[string]int
}

func (ct *ConditionTypes) init() {
	ct.NamedValues = map[string]int{
		"preImageSha256": 0,
		"prefixSha256": 1,
		"thresholdSha256":2,
		"ed25519Sha256":4,
	}
}

type NamedType struct {
	string
	*Condition
	ImplicitTag int
}

type NamedTypes struct {
	NamedType []NamedType
}

type Condition struct {
	Fingerprint string
	Cost int
}

type CompoundSha256Condition struct {
	*Condition
	Subtypes *ConditionTypes
}

type SimpleSha256Conditon struct {
	*Condition
}

