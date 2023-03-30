package config

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
)

type Config struct {
	GreenfieldConfig GreenfieldConfig `json:"greenfield_config"`
	BSCConfig        BSCConfig        `json:"bsc_config"`
	RelayConfig      RelayConfig      `json:"relay_config"`
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
	GRPCAddrs          []string `json:"grpc_addrs"`
	PrivateKey         string   `json:"private_key"`
	BlsPrivateKey      string   `json:"bls_private_key"`
	ChainId            uint64   `json:"chain_id"`
	StartHeight        uint64   `json:"start_height"`
	MonitorChannelList []uint8  `json:"monitor_channel_list"`
	GasLimit           uint64   `json:"gas_limit"`
	ChainIdString      string   `json:"chain_id_string"`
}

type BSCConfig struct {
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

type RelayConfig struct {
	BSCToGreenfieldInturnRelayerTimeout int64  `json:"bsc_to_greenfield_inturn_relayer_timeout"` // in second
	GreenfieldToBSCInturnRelayerTimeout int64  `json:"greenfield_to_bsc_inturn_relayer_timeout"` // in second
	GreenfieldSequenceUpdateLatency     int64  `json:"greenfield_sequence_update_latency"`       // in second
	BSCSequenceUpdateLatency            int64  `json:"bsc_sequence_update_latency"`              // in second
	GreenfieldEventTypeCrossChain       string `json:"greenfield_event_type_cross_chain"`
	BSCCrossChainPackageEventName       string `json:"bsc_cross_chain_package_event_name"`
	CrossChainPackageEventHex           string `json:"cross_chain_package_event_hex"`
	CrossChainContractAddr              string `json:"cross_chain_contract_addr"`
	GreenfieldLightClientContractAddr   string `json:"greenfield_light_client_contract_addr"`
}

func (cfg *BSCConfig) Validate() {
	if len(cfg.RPCAddrs) == 0 {
		panic("provider address of Binance Smart Chain should not be empty")
	}

	if cfg.KeyType == "" {
		panic("key_type Binance Smart Chain should not be empty")
	}
	if cfg.KeyType != KeyTypeLocalPrivateKey && cfg.KeyType != KeyTypeAWSPrivateKey {
		panic(fmt.Sprintf("key_type of Binance Smart Chain only supports %s and %s", KeyTypeLocalPrivateKey, KeyTypeAWSPrivateKey))
	}
	if cfg.KeyType == KeyTypeAWSPrivateKey && cfg.AWSRegion == "" {
		panic("aws_region of Binance Smart Chain should not be empty")
	}
	if cfg.KeyType == KeyTypeAWSPrivateKey && cfg.AWSSecretName == "" {
		panic("aws_secret_name of Binance Smart Chain should not be empty")
	}
	if cfg.KeyType != KeyTypeAWSPrivateKey && cfg.PrivateKey == "" {
		panic("privateKey of Binance Smart Chain should not be empty")
	}
	if cfg.GasLimit == 0 {
		panic("gas_limit of Binance Smart Chain should be larger than 0")
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
	EnableAlert     bool  `json:"enable_alert"`
	EnableHeartBeat bool  `json:"enable_heart_beat"`
	Interval        int64 `json:"interval"`

	Identity       string `json:"identity"`
	TelegramBotId  string `json:"telegram_bot_id"`
	TelegramChatId string `json:"telegram_chat_id"`

	BalanceThreshold     string `json:"balance_threshold"`
	SequenceGapThreshold uint64 `json:"sequence_gap_threshold"`
}

func (cfg *AlertConfig) Validate() {
	if !cfg.EnableAlert {
		return
	}
	if cfg.Interval <= 0 {
		panic("alert interval should be positive")
	}
	balanceThreshold, ok := big.NewInt(1).SetString(cfg.BalanceThreshold, 10)
	if !ok {
		panic("unrecognized balance_threshold")
	}

	if balanceThreshold.Cmp(big.NewInt(0)) <= 0 {
		panic("balance_threshold should be positive")
	}

	if cfg.SequenceGapThreshold <= 0 {
		panic("sequence_gap_threshold should be positive")
	}
}

type DBConfig struct {
	Dialect       string `json:"dialect"`
	KeyType       string `json:"key_type"`
	AWSRegion     string `json:"aws_region"`
	AWSSecretName string `json:"aws_secret_name"`
	Password      string `json:"password"`
	Username      string `json:"username"`
	Url           string `json:"url"`
}

func (cfg *DBConfig) Validate() {
	if cfg.Dialect != DBDialectMysql {
		panic(fmt.Sprintf("only %s supported", DBDialectMysql))
	}
	if cfg.Username == "" || cfg.Password == "" || cfg.Url == "" {
		panic("db config is not correct")
	}
}

func (cfg *Config) Validate() {
	cfg.AdminConfig.Validate()
	cfg.LogConfig.Validate()
	cfg.BSCConfig.Validate()
	cfg.AlertConfig.Validate()
	cfg.DBConfig.Validate()
}

func ParseConfigFromJson(content string) *Config {
	var config Config
	if err := json.Unmarshal([]byte(content), &config); err != nil {
		panic(err)
	}
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
