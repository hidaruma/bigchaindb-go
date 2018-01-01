package schema

import (
	"fmt"
	"encoding/json"
	"github.com/graphql-go/graphql"
)

type OutputType struct {
	Name string
	Description string

	Condition 
	PublicKeys
	Amount
}

func (ot *OutputType) FromJson(cls, output ) {
	return cls(** output)
}

type InputType struct {
	Name string
	Description string

	OwnerBefore
	Fulfillment
	Fulfills
}

func (it *InputType) FromJson(cls, input ) {
	it.Fulfills = 
	return 
}


type TransactionType struct {
	Name string
	Description string

	ID string
	Operation string
	Version string
	Asset 
	Metadata
	Inputs []Inputtype
	Outputs []OutputType
}

func (tt *TransactionType) FromJson(cls, retrievedTx) {
	
	
}

type FulfillsType struct {
	Name string
	Description string
	OutputIndex int
	Transaction
}

func (ft *FulfillsType) FromJson(cls, fulfills) {
	
}

type QueryType struct {
	Name string
	Description string
	Transaction
	Transactions
	Outputs
}


func (qt *QueryType) ResolveTransaction(args, context, info ) {
	
}

func (qt *QueryType) ResolveTransactions(args, context, info) {
	
}

func (qt *QueryType) ResolveOutputs(args, context, info) {
	
}

func (qt *QueryType) Schema() {
	var query QueryType
	
}