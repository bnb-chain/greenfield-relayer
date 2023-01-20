package config

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	DBDialectMysql         = "mysql"
	LocalConfig            = "local"
	AWSConfig              = "aws"
	KeyTypeLocalPrivateKey = "local_private_key"
	KeyTypeAWSPrivateKey   = "aws_private_key"
)

var (
	EventTypeCrossChain                = "cosmos.crosschain.v1.EventCrossChain"
	CrossChainPackageEventName         = "CrossChainPackage"
	CrossChainPackageEventHash         = common.HexToHash("0x64998dc5a229e7324e622192f111c691edccc3534bbea4b2bd90fbaec936845a")
	CrossChainContractAddr             = common.HexToAddress("0x2078fEEF78BD06AcDdb46619a681327aaEeeAE20")
	InscriptionLightClientContractAddr = common.HexToAddress("0x71b750F84B4d1d72C17EcEba811fA6E4C8c9CfdC")
)
