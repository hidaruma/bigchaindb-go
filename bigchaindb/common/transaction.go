package common

import (
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
	"github.com/hidaruma/bigchaindb-go/cryptoconditions"
	"github.com/hidaruma/bigchaindb-go/bigchaindb/common"
	base58 "github.com/itchyny/base58-go"
)


type Recipient map[string]string

type Fulfill struct {
	*TransactionLink
}

type Input struct {
	Fulfillment *cryptoconditions.Fulfillment
	OwnersBefore []string
	Fulfills []*Fulfill
}


func (i *Input) eq(other *Input) bool {
	if i == other {
		return true
	} else {
		return false
	}
}
/*
func (i *Input) ToDict() map[string]interface{} {
	var fulfillment *cryptoconditions.Fulfillment
	fulfillment, err := i.Fulfillment.SerializeUrl()
	if err != nil {
		switch err.(type) {
			case TypeError, AttributeError, ASN1EncodeError:
				fulfillment = FulfillmentToDetails(i.Fulfillment)
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
*/
func (i *Input) Generate(publicKeys []string) *Input {

	var output *Output
	output = Output.Generate(publicKeys, 1)
	i.Fulfillment = output.Fulfillment
	i.OwnersBefore = publicKeys
	return i
}
/*

func (i *Input) FromDict(data [string]string) *Input {
	var fulfillment
	fulfillment = data["fulfillment"]

	if fulfillment.(type) != cryptocondition.Fulfillment {
		fulfillment, err := cryptocondition.Fulfillment.FromUrl(data["fulfillment"])
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
*/
func fulfillmentToDatails(fulfillment *cryptoconditions.Fulfillment) map[string]interface{} {
	var details map[string]interface{}
	switch fulfillment.TypeName() {
		case "ed25519-sha-256":
			details["type"] = "ed25519-sha-256"
			pubKey, err  := base58.BitcoinEncoding.Encode(fulfillment.PublicKey)
			if err != nil {
				log.Println()
			}
			details["publick_key"] = string(pubKey)
			return details
		case "threshold-sha-256":
			var subconditions map[string]string
			for _, cond := range fulfillment.Subconditions {
				subconditions = append(subconditions, fulfillmentToDetails(cond["body"]))
			}
			details["type"] = "threshold-sha-256"
			details["thresold"] = fulfillment.Threshold
			details["subconditions"] = subconditions
			return details
		default:
			return nil
	}
}

func fulfillmentFromDetails(data map[string]interface{}, depth int) *Fulfillment {
	if depth == 100 {
		log.Fatal(ThresholdTooDeep)
	}

	switch data["type"].(string) {
		case "ed25519-sha-256":
			var publicKey []byte
			var decodePubKeyString []byte
			decodePubKeyString = []byte(data["public_key"].(string))
			publicKey, err := base58.BitcoinEncoding.Decode(decodePubKeyString)
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

type Asset map[string]string
type Metadata map[string]string
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

func (t *Transaction) Transfer(inputs []Input, recipients []Recipient, assetID string, metadata Metadata) *Transaction{
	if len(inputs) == 0 {
		log.Println(ValueError())
	}
	if len(recipients) == 0 {
		log.Println(ValueError())
	}
	var outputs []Output
	for _, recipient := range recipients {
		for pubKeys, amount := recipient {
			outputs = append(outputs, Output.Generate(pubKeys, amount))
		}
	}
	t.Operation = t.TRANSFER()
	t.Asset["id"] = assetID
	t.Inputs = inputs
	t.Outputs = outputs
	t.Metadata = metadata
	return t
}

func (t *transaction) eq(other *Transaction) bool {
	
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

func (t *Transaction) signSimpleSignatureFulfillment(input Input, message string, keyPairs map[string]string) Input {
	var input_ Input
	input_ = input
	var publicKey string
	publicKey = input_.OwnersBefore[0]
	err := input_.Fulfillment.Sign(message, base58.b58decode(string(keyPairs[publicKey]))
	if err != nil {
		log.Println(KeypairMismatchException())
	}
	return input_
}

func (t *Transaction) signThresholdSignatureFulfillment(input Input, message string, keyPairs map[string]string) Input {

	for _, ownerBefore := range input.OwnersBefore {
		var ccffill *cryptoconditions.Fulfillment
		ccffill = input.Fulfillment
		var subfills *[]Fulfill
		subfills = ccffill.GetSubconditionFromVk(base58.b58Encode(ownerBefore))
		if subfills == nil {
			log.Println(KeypairMicmatchException())
		}
		var privateKey string
		privateKey = keyPairs[ownerBefore]
		if privateKey.(type) != string {
			log.Println(KeypairMismatchException())
		}

	}
}

func (t *Transaction) InputsValid(outputs []Output) {
	if t.Operation == Transaction.CREATE() || t.Operation == Transaction.GENESIS() {
		outputs []Output
	}

		return t.inputsValid()
	}
}



func (t *Transaction) inputsValid(outputConditionURLs []string) bool {
	if len(t.Inputs) != len(outputConditionURLs) {
		log.Println(ValueError())
	}
	var txDict map[string]interface{}
	txDict = t.ToDict()
	txDict = Transaction.RemoveSignatures(txDict)
	txDict["id"] = nil
	var txSerialized string
	txSerialized = Transaction.toStr(txDict)
	var icond map[i]struct{ string;  }
	for i, cond := range outputConditionURLs {
		icond[i] = cond
	}
	return common.All(icond, t.validate)
}

func (t *Transaction) validate(i int, outputConditionURL string, txSerialized string) bool {
	return t.inputValid(t.Inputs[i], t.Operation, txSerialized, outputConditionURL)
}

func (t *Transaction) inputValid(input Input, operation string, txSerialized string, outputConditionURL string) bool {
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

func (t *Transaction) removeSignatures(txDict map[string]string) map[string]string {

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
	var tx *Transaction
	tx = Transaction.removeSignatures(t.ToDict())
	return Transaction.toStr(tx)
}

func (t *Transaction) GetAssetID(transactions []*Transaction) string {

	var assetIDs []string
	for _, tx := range transactions {
		if tx.Operation == Transaction.CREATE() {
			assetIDs = append(assetIDs, tx.ID)
		} else {
			assetIDs = append(assetIDs, tx.Asset["id"])
		}
	}
	if len(assetIDs) > 1 {
		log.Println(AssetIDMismatch())
	}
	return assetIDs[0]
}

func (t *Transaction) ValidateID(txBody) {

}

func (t *Transaction) FromDict(tx map[string]interface{}) *Transaction{
	var inputs []*Input
	var outputs []*Output
	switch tx["inputs"].(type) {
	case []*Input :
	for _, input := range tx["inputs"] {
		inputs = append(inputs, Input.FromDict(input))
	}
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