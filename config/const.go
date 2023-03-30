package config

const (
	FlagConfigPath          = "config-path"
	FlagConfigType          = "config-type"
	FlagConfigAwsRegion     = "aws-region"
	FlagConfigAwsSecretKey  = "aws-secret-key"
	FlagConfigPrivateKey    = "private-key"
	FlagConfigBlsPrivateKey = "bls-private-key"
	FlagConfigDbPass        = "db-pass"

	DBDialectMysql         = "mysql"
	LocalConfig            = "local"
	AWSConfig              = "aws"
	KeyTypeLocalPrivateKey = "local_private_key"
	KeyTypeAWSPrivateKey   = "aws_private_key"
)
