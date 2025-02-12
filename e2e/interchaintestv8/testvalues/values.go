package testvalues

import (
	"time"

	"cosmossdk.io/math"

	ibctm "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"

	"github.com/strangelove-ventures/interchaintest/v8/chain/ethereum"
)

const (
	// StartingTokenAmount is the amount of tokens to give to each user at the start of the test.
	StartingTokenAmount int64 = 10_000_000_000

	// StartingERC20TokenAmount is the amount of ERC20 tokens to give to each user at the start of the test.
	StartingERC20TokenAmount int64 = 1_000_000_000_000_000_000

	// ERC20TransferAmount is the default transfer amount on the ERC20 side of things
	ERC20TransferAmount int64 = 1_000_000_000_000_000

	// SDKTransferAmount is the default transfer amount on the SDK side of things
	SDKTransferAmount int64 = 1_000

	// EnvKeyTendermintRPC Tendermint RPC URL.
	EnvKeyTendermintRPC = "TENDERMINT_RPC_URL"
	// EnvKeyEthRPC Ethereum RPC URL.
	EnvKeyEthRPC = "RPC_URL"
	// EnvKeyOperatorPrivateKey Private key used to submit transactions by the operator.
	EnvKeyOperatorPrivateKey = "PRIVATE_KEY"
	// EnvKeySp1Prover The prover type (local|network|mock).
	EnvKeySp1Prover = "SP1_PROVER"
	// EnvKeySp1PrivateKey Private key for the prover network.
	EnvKeySp1PrivateKey = "SP1_PRIVATE_KEY"
	// EnvKeyGenerateFixtures Generate fixtures for the solidity tests if set to true.
	EnvKeyGenerateFixtures = "GENERATE_FIXTURES"
	// The log level for the Rust logger.
	EnvKeyRustLog = "RUST_LOG"

	// Log level for the Rust logger.
	EnvValueRustLog_Info = "info"
	// EnvValueSp1Prover_Network is the prover type for the network prover.
	EnvValueSp1Prover_Network = "network"
	// EnvValueSp1Prover_Mock is the prover type for the mock prover.
	EnvValueSp1Prover_Mock = "mock"
	// EnvValueGenerateFixtures_True is the value to set to generate fixtures for the solidity tests.
	EnvValueGenerateFixtures_True = "true"

	// Sp1GenesisFilePath is the path to the genesis file for the SP1 chain.
	// This file is generated and then deleted by the test.
	Sp1GenesisFilePath = "e2e/genesis.json"
	// FixturesDir is the directory where the Solidity fixtures are stored.
	FixturesDir = "test/fixtures/"

	// FaucetPrivateKey is the private key of the faucet account.
	// '0x' prefix is trimmed.
	FaucetPrivateKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
)

var (
	// MaxDepositPeriod Maximum period to deposit on a proposal.
	// This value overrides the default value in the gov module using the `modifyGovV1AppState` function.
	MaxDepositPeriod = time.Second * 10
	// VotingPeriod Duration of the voting period.
	// This value overrides the default value in the gov module using the `modifyGovV1AppState` function.
	VotingPeriod = time.Second * 30

	// StartingEthBalance is the amount of ETH to give to each user at the start of the test.
	StartingEthBalance = math.NewInt(2 * ethereum.ETHER)

	// DefaultTrustLevel is the trust level used by the SP1ICS07Tendermint contract.
	DefaultTrustLevel = ibctm.Fraction{Numerator: 2, Denominator: 3}.ToTendermint()

	// DefaultTrustPeriod is the trust period used by the SP1ICS07Tendermint contract.
	DefaultTrustPeriod = 1209669
)
