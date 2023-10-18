package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	GreenfieldConfig GreenfieldConfig `json:"greenfield_config"`
	BSCConfig        BSCConfig        `json:"bsc_config"`
	RelayConfig      RelayConfig      `json:"relay_config"`
	VotePoolConfig   VotePoolConfig   `json:"vote_pool_config"`
	LogConfig        LogConfig        `json:"log_config"`
	AdminConfig      AdminConfig      `json:"admin_config"`
	AlertConfig      AlertConfig      `json:"alert_config"`
	DBConfig         DBConfig         `json:"db_config"`
}

type AdminConfig struct {
	Port uint16 `json:"port"`
}

func (cfg *AdminConfig) Validate() {
	if cfg.Port <= 0 || cfg.Port > 65535 {
		panic("port should be within (0, 65535]")
	}
}

type GreenfieldConfig struct {
	KeyType            string   `json:"key_type"`
	AWSRegion          string   `json:"aws_region"`
	AWSSecretName      string   `json:"aws_secret_name"`
	AWSBlsSecretName   string   `json:"aws_bls_secret_name"`
	RPCAddrs           []string `json:"rpc_addrs"`
	PrivateKey         string   `json:"private_key"`
	BlsPrivateKey      string   `json:"bls_private_key"`
	ChainId            uint64   `json:"chain_id"`
	StartHeight        uint64   `json:"start_height"`
	MonitorChannelList []uint8  `json:"monitor_channel_list"`
	GasLimit           int64    `json:"gas_limit"`
	FeeAmount          int64    `json:"fee_amount"`
	ChainIdString      string   `json:"chain_id_string"`
	UseWebsocket       bool     `json:"use_websocket"`
}

func (cfg *GreenfieldConfig) Validate() {
	if len(cfg.RPCAddrs) == 0 {
		panic("provider address of Greenfield should not be empty")
	}

	if cfg.KeyType == "" {
		panic("key_type Greenfield should not be empty")
	}
	if cfg.KeyType != KeyTypeLocalPrivateKey && cfg.KeyType != KeyTypeAWSPrivateKey {
		panic(fmt.Sprintf("key_type of Greenfield only supports %s and %s", KeyTypeLocalPrivateKey, KeyTypeAWSPrivateKey))
	}
	if cfg.KeyType == KeyTypeAWSPrivateKey && cfg.AWSRegion == "" {
		panic("aws_region of Greenfield should not be empty")
	}
	if cfg.KeyType == KeyTypeAWSPrivateKey && cfg.AWSSecretName == "" {
		panic("aws_secret_name of Greenfield should not be empty")
	}
	if cfg.KeyType != KeyTypeAWSPrivateKey && cfg.PrivateKey == "" {
		panic("privateKey of Greenfield should not be empty")
	}
	if cfg.GasLimit <= 0 || cfg.FeeAmount <= 0 {
		panic("gas_limit and fee_amount should be larger than 0")
	}
}

type BSCConfig struct {
	OpBNB                     bool     `json:"op_bnb"`
	KeyType                   string   `json:"key_type"`
	AWSRegion                 string   `json:"aws_region"`
	AWSSecretName             string   `json:"aws_secret_name"`
	RPCAddrs                  []string `json:"rpc_addrs"`
	PrivateKey                string   `json:"private_key"`
	GasLimit                  uint64   `json:"gas_limit"`
	GasPrice                  uint64   `json:"gas_price"`
	NumberOfBlocksForFinality uint64   `json:"number_of_blocks_for_finality"`
	StartHeight               uint64   `json:"start_height"`
	ChainId                   uint64   `json:"chain_id"`
}

func (cfg *BSCConfig) Validate() {
	if len(cfg.RPCAddrs) == 0 {
		panic("provider address of BNB Smart Chain should not be empty")
	}

	if cfg.KeyType == "" {
		panic("key_type BNB Smart Chain should not be empty")
	}
	if cfg.KeyType != KeyTypeLocalPrivateKey && cfg.KeyType != KeyTypeAWSPrivateKey {
		panic(fmt.Sprintf("key_type of BNB Smart Chain only supports %s and %s", KeyTypeLocalPrivateKey, KeyTypeAWSPrivateKey))
	}
	if cfg.KeyType == KeyTypeAWSPrivateKey && cfg.AWSRegion == "" {
		panic("aws_region of BNB Smart Chain should not be empty")
	}
	if cfg.KeyType == KeyTypeAWSPrivateKey && cfg.AWSSecretName == "" {
		panic("aws_secret_name of BNB Smart Chain should not be empty")
	}
	if cfg.KeyType != KeyTypeAWSPrivateKey && cfg.PrivateKey == "" {
		panic("privateKey of BNB Smart Chain should not be empty")
	}
	if cfg.GasLimit == 0 {
		panic("gas_limit of BNB Smart Chain should be larger than 0")
	}
	if cfg.NumberOfBlocksForFinality < 2 || cfg.NumberOfBlocksForFinality > 21 {
		panic("NumberOfBlocksForFinality should be [2, 21]")
	}
}

func (cfg *BSCConfig) IsOpCrossChain() bool {
	return cfg.OpBNB
}

type RelayConfig struct {
	BSCToGreenfieldInturnRelayerTimeout int64  `json:"bsc_to_greenfield_inturn_relayer_timeout"` // in second
	GreenfieldToBSCInturnRelayerTimeout int64  `json:"greenfield_to_bsc_inturn_relayer_timeout"` // in second
	GreenfieldSequenceUpdateLatency     int64  `json:"greenfield_sequence_update_latency"`       // in second
	BSCSequenceUpdateLatency            int64  `json:"bsc_sequence_update_latency"`              // in second
	CrossChainContractAddr              string `json:"cross_chain_contract_addr"`
	GreenfieldLightClientContractAddr   string `json:"greenfield_light_client_contract_addr"`
	RelayerHubContractAddr              string `json:"relayer_hub_contract_addr"`
}

func (cfg *RelayConfig) Validate() {
	if cfg.BSCToGreenfieldInturnRelayerTimeout <= 0 {
		panic("bsc_to_greenfield_inturn_relayer_timeout should be larger than 0")
	}
	if cfg.GreenfieldToBSCInturnRelayerTimeout <= 0 {
		panic("greenfield_to_bsc_inturn_relayer_timeout should be larger than 0")
	}
	if cfg.GreenfieldSequenceUpdateLatency <= 0 {
		panic("greenfield_sequence_update_latency should be larger than 0")
	}
	if cfg.BSCSequenceUpdateLatency <= 0 {
		panic("bsc_sequence_update_latency should be larger than 0")
	}
}

type VotePoolConfig struct {
	BroadcastIntervalInMillisecond int64 `json:"broadcast_interval_in_millisecond"`
	VotesBatchMaxSizePerInterval   int64 `json:"votes_batch_max_size_per_interval"`
	QueryIntervalInMillisecond     int64 `json:"query_interval_in_millisecond"`
}

func (cfg *VotePoolConfig) Validate() {
	if cfg.BroadcastIntervalInMillisecond <= 0 {
		panic("broadcast_interval_in_millisecond should be larger than 0")
	}
	if cfg.VotesBatchMaxSizePerInterval <= 0 {
		panic("votes_batch_max_size_per_interval should be larger than 0")
	}
	if cfg.QueryIntervalInMillisecond <= 0 {
		panic("query_interval_in_millisecond should be larger than 0")
	}
}

type LogConfig struct {
	Level                        string `json:"level"`
	Filename                     string `json:"filename"`
	MaxFileSizeInMB              int    `json:"max_file_size_in_mb"`
	MaxBackupsOfLogFiles         int    `json:"max_backups_of_log_files"`
	MaxAgeToRetainLogFilesInDays int    `json:"max_age_to_retain_log_files_in_days"`
	UseConsoleLogger             bool   `json:"use_console_logger"`
	UseFileLogger                bool   `json:"use_file_logger"`
	Compress                     bool   `json:"compress"`
}

func (cfg *LogConfig) Validate() {
	if cfg.UseFileLogger {
		if cfg.Filename == "" {
			panic("filename should not be empty if use file logger")
		}
		if cfg.MaxFileSizeInMB <= 0 {
			panic("max_file_size_in_mb should be larger than 0 if use file logger")
		}
		if cfg.MaxBackupsOfLogFiles <= 0 {
			panic("max_backups_off_log_files should be larger than 0 if use file logger")
		}
	}
}

type AlertConfig struct {
	Identity       string `json:"identity"`
	TelegramBotId  string `json:"telegram_bot_id"`
	TelegramChatId string `json:"telegram_chat_id"`
}

type DBConfig struct {
	Dialect       string `json:"dialect"`
	KeyType       string `json:"key_type"`
	AWSRegion     string `json:"aws_region"`
	AWSSecretName string `json:"aws_secret_name"`
	Password      string `json:"password"`
	Username      string `json:"username"`
	Url           string `json:"url"`
	MaxIdleConns  int    `json:"max_idle_conns"`
	MaxOpenConns  int    `json:"max_open_conns"`
}

func (cfg *DBConfig) Validate() {
	if cfg.Dialect != DBDialectMysql && cfg.Dialect != DBDialectSqlite3 {
		panic(fmt.Sprintf("only %s and %s supported", DBDialectMysql, DBDialectSqlite3))
	}
	if cfg.Dialect == DBDialectMysql && (cfg.Username == "" || cfg.Url == "") {
		panic("db config is not correct, missing username and/or url")
	}
	if cfg.MaxIdleConns == 0 || cfg.MaxOpenConns == 0 {
		panic("db connections is not correct")
	}
}

func (cfg *Config) Validate() {
	cfg.AdminConfig.Validate()
	cfg.LogConfig.Validate()
	cfg.GreenfieldConfig.Validate()
	cfg.BSCConfig.Validate()
	cfg.RelayConfig.Validate()
	cfg.VotePoolConfig.Validate()
	cfg.DBConfig.Validate()
}

func ParseConfigFromJson(content string) *Config {
	var config Config
	if err := json.Unmarshal([]byte(content), &config); err != nil {
		panic(err)
	}
	config.Validate()
	return &config
}

func ParseConfigFromFile(filePath string) *Config {
	bz, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(bz, &config); err != nil {
		panic(err)
	}
	config.Validate()
	return &config
}
