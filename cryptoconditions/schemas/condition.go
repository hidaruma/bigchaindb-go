package schemas

import (
		"encoding/asn1"
)

type ConditionTypes struct {
	*asn1.BitString
	NamedValues map[string]int
}

func (ct *ConditionTypes) init(){
	ct.NamedValues = asn1.{
									"preImageSha256":0,
									"prefixSha256": 1,
									"thresholdSha256": 2,
									"rsaSha256": 3,
									"ed25519Sha256":4,
									}
}

