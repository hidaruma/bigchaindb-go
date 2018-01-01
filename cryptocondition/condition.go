package cryptocondition

import (
	"encoding/asn1"
)

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

