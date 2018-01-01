package common

import (
	"github.com/hidaruma/bigchaindb-go/bigchaindb"
)

type ConfigurationError *bigchaindb.BigchainDBError

type DatabaseAlreadyExists *bigchaindb.BigchainDBError

type DatabaseDoesNotExist *bigchaindb.BigchainDBError

type StartupError *bigchaindb.BigchainDBError

type CyclicBlockchainError *bigchaindb.BigchainDBError

type KeypairNotFoundException *bigchaindb.BigchainDBError

type OperationError *bigchaindb.BigchainDBError

type ValidationError *bigchaindb.BigchainDBError

type DoubleSpend *ValidationError

type InvalidHash *ValidationError

type SchemaValidationError *ValidationError

type InvalidSignature *ValidationError

type ImproperVotesError *ValidationError

type MultipleVotesError *ValidationError

type TransactionNotInValidBlock *ValidationError

type AssetIdMismatch *ValidationError

type AmountError *ValidationError

type InputDoesNotExist *ValidationError

type TransactionOwnerError *ValidationError

type SybilError *ValidationError

type DuplicateTransaction *ValidationError

type ThresholdTooDeep *ValidationError

type GenesisBlockAlreadyExistsError *ValidationError