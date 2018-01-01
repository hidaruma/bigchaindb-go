package common

import (
	"github.com/itchyny/base58-go"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/"
	"log"
	"strings"
	"os/exec"
	"net/http"
	"encoding/base64"
	"encoding/asn1"
	"go/types"
	"regexp"
	"net/url"
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

type Fulfills struct {
	*TransactionLink
}

type Input struct {
	Fulfillment Fulfillment
	OwnersBefore []string
	Fulfills Fulfills
}


func (i *Input) Eq(other *Input) bool {
	if i.ToDict() == other.ToDict() {
		return true
	} else {
		return false
	}
}

func (i *Input) ToDict() error {
	var fulfillment
	fulfillment, err := i.Fulfillment.SerializeUrl()
	if err != nil {
		switch err.(type) {
			case exceptions.TypeError, exceptions.AttributeError, exceptions.ASN1EncodeError:
				fulfillment = fulfillmentToDetails(i.Fulfillment)
		}
	}
	var fulfills
	fulfills, err = i.Fulfills.ToDict()
	if err != nil {
		
	
	}
	var input Input
	input = Input{}

	return input
}

func (i *Input) Generate(publicKeys []string) *Input {

	var output
	output = Output.Generate(publicKeys, 1)
	i.Fulfillment = output.Fulfillment
	i.OwnerBefore = publicKeys
	i.Fulfills = {}
	return i
}

func (i *Input) FromDict(data [string]string) *Input {
	var fulfillment
	fulfillment = data["fulfillment"]

	if fulfillment.(type) != Fulfillment {
		fulfillment, err := Fulfillment.FromUrl(data["fulfillment"])
		if err != nil {
			switch err {
				case exceptions.ASN1DecodeError:
					log.Println(err)
				case exceptions.TypeError:
					fulfillment = FulfillmentFromDetails(data["fulfillment"])
			}
		}
	
	}
	var fulfills
	fulfills = TransactionLink.FromDict(data["fulfills"])
	i.Fulfills = fulfills
	i.OwnerBefore = data["owner_before"]
	i.Fulfillment = fulfillment
	return i
}

func FulfillmentToDatails(fulfillment ) ([string]interface{}, error) {
	switch fulfillment.typeName {
		case :
			return {"type": "ed25519-sha-256",
				"public_key": base58.b58encode(fulfillment.PublicKey)
				}, nil
		case "threshold-sha-256":
			var subconditions = map([string]string)
			for _, cond := range fulfillment.subconditions {
				subconditions = append(subconditions, FulfillmentToDetails(cond["body"]))
			}
			return {
				"type":"threshold-sha-256",
				"threshold": fulfillment.threshold,
				"subconditions": subconditions,
			}, nil
		default:
			return nil, exceptions.UnsupportedTypeError()
	}
}

func FulfillmentFromDetails(data tx.Output[].Condition.Details, depth int) (, error){
	if depth == 100 {
		return nil, exceptions.ThresholdTooDeep()
	}

	switch data["type"] {
		case "ed25519-sha-256":
			var publicKey string
			publicKey = base58.b58decode(data["public_key"])
			return ed25519sha256(publicKey), nil
		case "threshold-sha-256":
			var threshold 
			threshold = thresholdsha256(data["threshold"])
			for _, cond := range data["subconditions"] {
				cond, err := FulfillmentFromDetails(cond, depth + 1)
				if err != nil {
					log.Println(err)
				}
			}
			threshold.AddSubfulfillment(cond)
			return threshold, nil
		default:
			return nil, exceptions.UnsupportedTypeError(data.(type))
			}
}

type TransactionLink struct {
	TxID string
	Output int
}

func (tl *TransactionLink) Bool() {
	
}

func (tl *TransactionLink) Eq() {
	
}

func (tl *TransactionLink) Hash() {
	
}

func (tl *TransactionLink) FromDict(link [string]string) *transactionLink{
	
}

func (tl *TransactionLink) ToDict() {
	
}

func (tl *TransactionLink) ToURL(path string) [string]string {
	if tl.TxID != nil && tl.Output != nil {
		return nil
	} else {
		return path + "/transactions" + tl.TxID + "/outputs" + strings.Itoa(tl.Output)
	}
}

type Output struct {
	Fulfillment
	PublicKeys []string
}

func (o *Output) MaxAmount() int{
	return 9 * 10 *18
}

func (o *Output) Eq(other ) bool {
	
}

func (o *Output) ToDict() Output {
	var conditions
}

func (o *Output) Generate(publicKeys []string, amount int) Output {
	
}

func (o *Output) genCondition(initial, newPublicKeys) {
	
}

func (o *Output) FromDict(data [string]string) {
	
}

type Asset [string]string
type Metadata [string]string
type Transaction struct {
	Operation string
	Inputs []Input
	Outputs []Output
	Asset Asset
	Metadata Metadata
	Version string
}

func (t *Transaction) CREATE() string {
	return "CREATE"
}

func (t *Transaction) TRANSFER() string {
	return "TRANSFER"
}

func (t *Transaction) GENESIS() string {
	return "GENESIS"
}

func (t *Transaction) ALLOWEDOPERATIONS() (string, string, string) {
	return "CREATE", "TRANSFER", "GENESIS"
}

func (t *Transaction) VERSION() string {
	return "1.0"
}

func (t *Transaction) Serialized() string{
	
}

func (t *Transaction) hash() {
	t.id = HashData(t.serialized())
}

func (t *Transaction) Create(txSigners []string, recipients , metadata Metadata, asset Asset) *Transaction{

	t.Operation = t.CREATE()
	t.Inputs = inputs
	t.Asset = {"data": asset}
	t.Outputs = outputs
	t.Metadata = metadata
	return t
}

func (t *Transaction) Transfer(inputs []Input, recipients ,assetID string, metadata Metadata) *Transaction{

	t.Operation = t.TRANSFER()
	t.Asset = append(t.Asset, {"id": assetID})
	t.Inputs = inputs
	t.Outputs = outputs
	t.Metadata = metadata
	return t
}

func (t *transaction) Eq(other ) bool {
	
}

func (t *Transaction) ToInputs(indices []int) []Input {
	var 
}

func (t *Transaction) AddInputs(input Input) error {
	
}

func (t *Transaction) AddOutput(output Output) error {
	
}

func (t *Transaction) Sign(privateKeys []string) *Transaction {
	if privateKeys == nil || privateKeys.(type) != []string {
		log.Println(exceptions.TypeError())
	}
	var keyPairs = map[string]string{}
	for _, privateKey := range privateKeys {
		keyPairs[t.Sign.GenpublicKey(PrivateKey(privateKey))] = PrivateKey(privateKey)
	}
	var txDict []string
	txDict = t.ToDict()
	txDict = Transaction.removeSignatures(txDict)
	var txSerialized string
	txSerialized = Transaction.toStr(txDict)
	for i, input := range t.Inputs {
		y.Inputs[i] = t.signInput(input, txSerialized, keyPairs)
	}
	t.hash()

	return t
}

func (s *Sign) GenpublicKey(privateKey crypto.PrivateKey) string {
	var publicKey
	publicKey = privateKey

	return publicKey.Decode()
}

func (t *Transaction) signInput(input Input, message string, keyPairs [string]string) {
	switch input.Fulfillment {
	case "ed25519sha256":
		t.signSimpleSignatureFulfillment(input, message, keyPairs)
	case "threshold-sha-256":
		t.signThresholdSignatureFulfillment(input, message, keyPairs)
	default:
		log.Println(exceptions.ValueError())
	}
}

func (t *Transaction) signSimpleSignatureFulfillment(input Input, message string, keyPairs [string]string) Input {
	var input_ Input
	input_ = input
	var publicKey string
	publicKey = input_.OwnersBefore[0]
	err := input_.Fulfillment.Sign(message.Encode(), base58.b58decode(keyPairs[publicKey].Encode()))
	if err != nil {
		log.Println(exceptions.KeypairMismatchException())
	}
	return input_
}

func (t *Transaction) signThresholdSignatureFulfillment(input Input, message string, keyPairs [string]string) Input {

}

func (t *Transaction) inputsValid(outputConditionURLs []string) {
	if len(t.Inputs) != len(outputConditionURLs) {
		log.Println(exceptions.ValueError())
	}

	return
}

func (t *Transaction) validate(i int, outputConditionURL string, txSerialized string, outputConditionURLs []string) {
	return t.inputValid(t.Inputs[i], t.Operation, txSerialized, outputConditionURLs)
}

func (t *Transaction) inputValid(in/home/hidarumaput Input, operation string, txSerialized string, outputConditionURL string) bool {
	var ccffill Fulfillment
	ccffill = input.Fulfillment
	var parsedFfill
	parsedFfill, err  := Fulfillment.FromURL(ccffill.SerializeURL())
	if err != nil {
		switch err {
		case exceptions.TypeError, exceptions.ASN1DecodeError, exceptions.ASN1EncodeError, exceptions.ValueError, exceptions.ParsingError:
			return false
		}
	}
	var outputValid bool
	switch operation {
	case Transaction.CREATE(), Transaction.GENESIS:
		outputValid = true
	default:
		if outputConditionURL == ccffill.ConditionURL {
			outputValid = true
			} else {
				outputValid = false
		}
	}
	var ffillValid bool
	ffillValid = parsedFfill.validate(message=txSerialized.Encode())
	return outputValid & ffillValid
}

func (t *Transaction) ToDict() [string]interface{}{

}

func (t *Transaction) removeSignatures(txDict [string]string) [string]string {

}

func (t *Transaction) toHash(value string) string {
	return crypto.HashData(value)
}

func (t *Transaction) ID() {
	return t.id
}

func (t *Transaction) ToHash() {
	return t.ToDict()["id"]
}

func (t *Transaction) toStr(value string) string {
	return utils.Serialize(value)
}

func (t *Transaction) str() string {
	var tx
	tx = transaction.removeSignatures(t.ToDict())
	return Transaction.toStr(tx)
}

func (t *Transaction) GetAssetID(transactions []Transaction) string {
	if transactions.(type) != []Transaction {
		transactions = [transactions]
	}
	var assetIDs = make([]string)
	for _, tx := range transactions {
		if tx.Operation == Transaction.CREATE() {
			assetIDs = append(assetIDs, tx.ID)
		} else {
			assetIDs = append(assetIDs, tx.Asset["id"])
		}
	}
	if len(assetIDs) > 1 {
		log.Println(exceptions.AssetIdMismatch())
	}
	return assetIDs.Pop()
}

func (t *Transaction) ValidateID(txBody) {

}

func (t *Transaction) FromDict(tx *Transaction) *Transaction{
	var inputs []Input
	var outputs []Output
	for _, input := range tx["inputs"] {
		inputs = append(inputs, Input.FromDict(input))
	}
	for _, output := range tx["outputs"] {
		outputs = append(outputs, Output.FromDict(output))
	}

	t.Operation = tx["operation"]
	t.Asset = tx["asset"]
	t.Inputs = inputs
	t.Outputs = outputs
	t.Metadata = tx["metadata"]
	t.Version = tx["version"]
	t.HashID = tx["id"]
	return t
}