package metric

import (
	"fmt"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
)

const (
	MetricNameGnfdSavedBlock       = "Greenfield_saved_block_height"
	MetricNameGnfdProcessedBlock   = "Greenfield_processed_block_height"
	MetricNameIsGnfdInturnRelayer  = "is_Greenfield_inturn_relayer"
	MetricNameGnfdRelayerStartTime = "Greenfield_relayer_start_time" // inturn relayer start time
	MetricNameGnfdRelayerEndTime   = "Greenfield_relayer_end_time"   // inturn relayer end time

	MetricNameBSCSavedBlock       = "BSC_saved_block_height"
	MetricNameBSCProcessedBlock   = "BSC_processed_block_height"
	MetricNameIsBSCInturnRelayer  = "is_BSC_inturn_relayer"
	MetricNameBSCRelayerStartTime = "BSC_relayer_start_time" // inturn relayer start time
	MetricNameBSCRelayerEndTime   = "BSC_relayer_end_time"   // inturn relayer end time

	MetricNameNextSequenceForChannelFromDB    = "next_seq_from_DB_for_channel"
	MetricNameNextSequenceForChannelFromChain = "next_seq_from_chain_for_channel"

	//for tracking perf
	MetricNameTxOnSrcChainTime        = "tx_on_src_chain_time_t1"
	MetricNameTxSavedToDBTime         = "tx_saved_to_db_time_t2"
	MetricNameTxBroadcastTime         = "tx_broadcast_time_t3"
	MetricNameTxFinishCollectVoteTime = "tx_finish_collect_vote_time_t4"
	MetricNameTxPickForAssembleTime   = "tx_picked_for_assemble_time_t5"
	MetricNameTxOnDestTime            = "tx_on_dest_chain_time_t6"
)

type MetricService struct {
	MetricsMap map[string]prometheus.Collector
	cfg        *config.Config
}

func NewMetricService(config *config.Config) *MetricService {
	ms := make(map[string]prometheus.Collector, 0)

	// Greenfield
	gnfdSavedBlockMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameGnfdSavedBlock,
		Help: "Saved block height for Greenfield in Database",
	})
	ms[MetricNameGnfdSavedBlock] = gnfdSavedBlockMetric
	prometheus.MustRegister(gnfdSavedBlockMetric)

	gnfdProcessedBlockMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameGnfdProcessedBlock,
		Help: "Processed block height for Greenfield in Database",
	})
	ms[MetricNameGnfdProcessedBlock] = gnfdProcessedBlockMetric
	prometheus.MustRegister(gnfdProcessedBlockMetric)

	gnfdIsInturnRelayerMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameIsGnfdInturnRelayer,
		Help: "Whether relayer is inturn to relay transaction from BSC to Greenfield",
	})
	ms[MetricNameIsGnfdInturnRelayer] = gnfdIsInturnRelayerMetric
	prometheus.MustRegister(gnfdIsInturnRelayerMetric)

	// Greenfield relayer(BSC -> Greenfield) relay interval metrics
	gnfdRelayerStartTimeMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameGnfdRelayerStartTime,
		Help: "inturn gnfd relayer start time or out-turn relayer previous start time",
	})
	ms[MetricNameGnfdRelayerStartTime] = gnfdRelayerStartTimeMetric
	prometheus.MustRegister(gnfdRelayerStartTimeMetric)

	gnfdRelayerEndTimeMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameGnfdRelayerEndTime,
		Help: "inturn gnfd relayer end time or out-turn relayer previous end time",
	})
	ms[MetricNameGnfdRelayerEndTime] = gnfdRelayerEndTimeMetric
	prometheus.MustRegister(gnfdRelayerEndTimeMetric)

	// BSC
	bscSavedBlockMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameBSCSavedBlock,
		Help: "Saved block height for BSC in Database",
	})
	ms[MetricNameBSCSavedBlock] = bscSavedBlockMetric
	prometheus.MustRegister(bscSavedBlockMetric)

	bscProcessedBlockMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameBSCProcessedBlock,
		Help: "Processed block height for BSC in Database",
	})
	ms[MetricNameBSCProcessedBlock] = bscProcessedBlockMetric
	prometheus.MustRegister(bscProcessedBlockMetric)

	bscIsInturnRelayerMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameIsBSCInturnRelayer,
		Help: "Whether relayer is inturn to relay transaction from Greenfield to BSC",
	})
	ms[MetricNameIsBSCInturnRelayer] = bscIsInturnRelayerMetric
	prometheus.MustRegister(bscIsInturnRelayerMetric)

	// BSC relayer(Greenfield -> BSC) relay interval metrics
	bscRelayerStartTimeMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameBSCRelayerStartTime,
		Help: "inturn BSC relayer start time or out-turn relayer previous start time",
	})
	ms[MetricNameBSCRelayerStartTime] = bscRelayerStartTimeMetric
	prometheus.MustRegister(bscRelayerStartTimeMetric)

	bscRelayerEndTimeMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameBSCRelayerEndTime,
		Help: "inturn BSC relayer end time or out-turn relayer previous end time",
	})
	ms[MetricNameBSCRelayerEndTime] = bscRelayerEndTimeMetric
	prometheus.MustRegister(bscRelayerEndTimeMetric)

	// register greenfield oracle channel
	nextSeqFromDB := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_%d", MetricNameNextSequenceForChannelFromDB, 0),
		Help: fmt.Sprintf("Next delivery sequence read from DB for channel %d", 0),
	})
	ms[fmt.Sprintf("%s_%d", MetricNameNextSequenceForChannelFromDB, 0)] = nextSeqFromDB
	prometheus.MustRegister(nextSeqFromDB)

	nextSeqFromChain := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_%d", MetricNameNextSequenceForChannelFromChain, 0),
		Help: fmt.Sprintf("Next delivery sequence read from chain for channel %d", 0),
	})
	ms[fmt.Sprintf("%s_%d", MetricNameNextSequenceForChannelFromChain, 0)] = nextSeqFromChain
	prometheus.MustRegister(nextSeqFromChain)

	// register gnfd -> bsc channels
	for _, c := range config.GreenfieldConfig.MonitorChannelList {
		nextSeqFromDB = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: fmt.Sprintf("%s_%d", MetricNameNextSequenceForChannelFromDB, c),
			Help: fmt.Sprintf("Next delivery sequence retreived from DB for channel %d", c),
		})
		ms[fmt.Sprintf("%s_%d", MetricNameNextSequenceForChannelFromDB, c)] = nextSeqFromDB
		prometheus.MustRegister(nextSeqFromDB)

		nextSeqFromChain = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: fmt.Sprintf("%s_%d", MetricNameNextSequenceForChannelFromChain, c),
			Help: fmt.Sprintf("Next delivery sequence retreived from chain for channel %d", c),
		})
		ms[fmt.Sprintf("%s_%d", MetricNameNextSequenceForChannelFromChain, c)] = nextSeqFromChain
		prometheus.MustRegister(nextSeqFromChain)
	}

	// tracking perf
	TxOnSrcChainTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameTxOnSrcChainTime,
		Help: MetricNameTxOnSrcChainTime,
	})
	ms[MetricNameTxOnSrcChainTime] = TxOnSrcChainTime
	prometheus.MustRegister(TxOnSrcChainTime)

	TxSavedToDBTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameTxSavedToDBTime,
		Help: MetricNameTxSavedToDBTime,
	})
	ms[MetricNameTxSavedToDBTime] = TxSavedToDBTime
	prometheus.MustRegister(TxSavedToDBTime)

	TxBroadcastTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameTxBroadcastTime,
		Help: MetricNameTxBroadcastTime,
	})
	ms[MetricNameTxBroadcastTime] = TxBroadcastTime
	prometheus.MustRegister(TxBroadcastTime)

	TxFinishCollectVoteTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameTxFinishCollectVoteTime,
		Help: MetricNameTxFinishCollectVoteTime,
	})
	ms[MetricNameTxFinishCollectVoteTime] = TxFinishCollectVoteTime
	prometheus.MustRegister(TxFinishCollectVoteTime)

	TxPickForAssembleTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameTxPickForAssembleTime,
		Help: MetricNameTxPickForAssembleTime,
	})
	ms[MetricNameTxPickForAssembleTime] = TxPickForAssembleTime
	prometheus.MustRegister(TxPickForAssembleTime)

	TxOnDestTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: MetricNameTxOnDestTime,
		Help: MetricNameTxOnDestTime,
	})
	ms[MetricNameTxOnDestTime] = TxOnDestTime
	prometheus.MustRegister(TxOnDestTime)

	return &MetricService{
		MetricsMap: ms,
		cfg:        config,
	}
}

func (m *MetricService) Start() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(fmt.Sprintf(":%d", m.cfg.AdminConfig.Port), nil)
	if err != nil {
		panic(err)
	}
}

func (m *MetricService) SetGnfdSavedBlockHeight(height uint64) {
	m.MetricsMap[MetricNameGnfdSavedBlock].(prometheus.Gauge).Set(float64(height))
}

func (m *MetricService) SetGnfdProcessedBlockHeight(height uint64) {
	m.MetricsMap[MetricNameGnfdProcessedBlock].(prometheus.Gauge).Set(float64(height))
}

func (m *MetricService) SetBSCSavedBlockHeight(height uint64) {
	m.MetricsMap[MetricNameBSCSavedBlock].(prometheus.Gauge).Set(float64(height))
}

func (m *MetricService) SetBSCProcessedBlockHeight(height uint64) {
	m.MetricsMap[MetricNameBSCProcessedBlock].(prometheus.Gauge).Set(float64(height))
}

func (m *MetricService) SetBSCInturnRelayerMetrics(isInturn bool, start, end uint64) {
	m.setIsBSCInturnRelayer(isInturn)
	m.setBSCInturnRelayerStartTime(start)
	m.setBSCInturnRelayerEndTime(end)
}

func (m *MetricService) setIsBSCInturnRelayer(isInturn bool) {
	var flag float64
	if isInturn {
		flag = 1
	}
	m.MetricsMap[MetricNameIsBSCInturnRelayer].(prometheus.Gauge).Set(flag)
}

func (m *MetricService) setBSCInturnRelayerStartTime(start uint64) {
	m.MetricsMap[MetricNameBSCRelayerStartTime].(prometheus.Gauge).Set(float64(start))
}

func (m *MetricService) setBSCInturnRelayerEndTime(end uint64) {
	m.MetricsMap[MetricNameBSCRelayerEndTime].(prometheus.Gauge).Set(float64(end))
}

func (m *MetricService) SetGnfdInturnRelayerMetrics(isInturn bool, start, end uint64) {
	m.setIsGnfdInturnRelayer(isInturn)
	m.setGnfdInturnRelayerStartTime(start)
	m.setGnfdInturnRelayerEndTime(end)
}

func (m *MetricService) setIsGnfdInturnRelayer(isInturn bool) {
	var flag float64
	if isInturn {
		flag = 1
	}
	m.MetricsMap[MetricNameIsGnfdInturnRelayer].(prometheus.Gauge).Set(flag)
}

func (m *MetricService) setGnfdInturnRelayerStartTime(start uint64) {
	m.MetricsMap[MetricNameGnfdRelayerStartTime].(prometheus.Gauge).Set(float64(start))
}

func (m *MetricService) setGnfdInturnRelayerEndTime(end uint64) {
	m.MetricsMap[MetricNameGnfdRelayerEndTime].(prometheus.Gauge).Set(float64(end))
}

func (m *MetricService) SetNextSequenceForChannelFromDB(channel uint8, seq uint64) {
	m.MetricsMap[fmt.Sprintf("%s_%d", MetricNameNextSequenceForChannelFromDB, channel)].(prometheus.Gauge).Set(float64(seq))
}

func (m *MetricService) SetNextSequenceForChannelFromChain(channel uint8, seq uint64) {
	m.MetricsMap[fmt.Sprintf("%s_%d", MetricNameNextSequenceForChannelFromChain, channel)].(prometheus.Gauge).Set(float64(seq))
}

func (m *MetricService) SetTxOnSrcChainTimeT1(ts uint64) {
	m.MetricsMap[MetricNameTxOnSrcChainTime].(*prometheus.GaugeVec).With(prometheus.Labels{"timestamp": strconv.FormatInt(int64(ts), 10)}).Set(0)
}
func (m *MetricService) SetTxSavedToDBTimeT2(ts uint64) {
	m.MetricsMap[MetricNameTxSavedToDBTime].(*prometheus.GaugeVec).With(prometheus.Labels{"timestamp": strconv.FormatInt(int64(ts), 10)}).Set(0)
}
func (m *MetricService) SetxBroadcastTimeT3(ts uint64) {
	m.MetricsMap[MetricNameTxBroadcastTime].(*prometheus.GaugeVec).With(prometheus.Labels{"timestamp": strconv.FormatInt(int64(ts), 10)}).Set(0)
}
func (m *MetricService) SetTxFinishCollectVoteTimeT4(ts uint64) {
	m.MetricsMap[MetricNameTxFinishCollectVoteTime].(*prometheus.GaugeVec).With(prometheus.Labels{"timestamp": strconv.FormatInt(int64(ts), 10)}).Set(0)
}
func (m *MetricService) SetTxPickForAssembleTimeT5(ts uint64) {
	m.MetricsMap[MetricNameTxPickForAssembleTime].(*prometheus.GaugeVec).With(prometheus.Labels{"timestamp": strconv.FormatInt(int64(ts), 10)}).Set(0)
}
func (m *MetricService) SetTxOnDestTimeT6(ts uint64) {
	m.MetricsMap[MetricNameTxOnDestTime].(*prometheus.GaugeVec).With(prometheus.Labels{"timestamp": strconv.FormatInt(int64(ts), 10)}).Set(0)
}
