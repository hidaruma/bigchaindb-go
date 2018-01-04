package cryptoconditions

import (
	"encoding/asn1"
	"strings"
	"regexp"
	"net/url"
	"encoding/base64"
	"strconv"
	"log"
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
	Fingerprint string
	Hash string
	Cost int
	Subtypes string
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

func (c *Condition) FromURI(serializedCondition interface{}) *Condition {
	var serializedConditionString string
	switch serializedCondition.(type) {
	case *Condition:
		return serializedCondition.(*Condition)
	case string:
		serializedConditionString = serializedCondition.(string)
	default:
		log.Println(TypeError())
	}
	var pieces []string
	pieces = strings.Split(serializedConditionString, ":")
	if pieces[0] != CONDITIONURISCHEME {
		log.Println(PrefixError("Serialized condition must start with" + CONDITIONURISCHEME + "."))
	}
	var regexMatch bool
	re := regexp.MustCompile(CONDITIONREGEXSTRICT)
	regexMatch = re.MatchString(serializedConditionString)
	if !regexMatch {
		log.Println(ParsingError)
	}
	var regexMatchGroup [][]string
	regexMatchGroup = re.FindAllStringSubmatch(serializedConditionString, -1)
	u, err := url.Parse(regexMatchGroup[0][2])
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
	var reInt = regexp.MustCompile(INTEGERREGEX)
	if reInt.MatchString(cost) {
		log.Println(ParsingError())
	}
	var fingerprint string
	fingerprint = regexMatchGroup[1][1]
	var condition *Condition
	condition.TypeID = conditionType["type_id"]
	if conditionType["class"].TYPECATEGORY == "compound" {
		condition.subtypes.Update(strings.Split(qsDict["subtypes"][0], ","))
	}
	var conditionHash []byte
	conditionHash, err = base64.URLEncoding.DecodeString(fingerprint)
	if err != nil {
		log.Println(err)
	}
	condition.Hash = string(conditionHash)
	condition.Cost, err = strconv.Atoi(cost)
	if err != nil {
		log.Println(err)
	}
	return condition
}

func (c *Condition) FromBinary(data []byte) {
	var asn1ConditionObj map[string]string
	var residue string
	asn1ConditionObj, residue = asn1.Unmarshal
	var asn1ConditionDict map[string]string

	return Condition.FromAsn1Dict(asn1ConditionDict)
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


type CompoundSha256Condition struct {
	*Condition
	Subtypes *ConditionTypes
}

type SimpleSha256Conditon struct {
	*Condition
}

