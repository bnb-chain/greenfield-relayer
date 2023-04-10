package metric

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/bnb-chain/greenfield-relayer/config"
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

	MetricNameNextSendSequenceForChannel    = "next_send_seq_for_channel"
	MetricNameNextReceiveSequenceForChannel = "next_receive_seq_for_channel"
)

type MetricService struct {
	MetricsMap map[string]prometheus.Metric
	cfg        *config.Config
}

func NewMetricService(config *config.Config) *MetricService {
	ms := make(map[string]prometheus.Metric, 0)

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
	nextSendOracleSeq := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_%d", MetricNameNextSendSequenceForChannel, 0),
		Help: fmt.Sprintf("Next Send Oracle sequence"),
	})
	ms[fmt.Sprintf("%s_%d", MetricNameNextSendSequenceForChannel, 0)] = nextSendOracleSeq
	prometheus.MustRegister(nextSendOracleSeq)

	nextReceiveOracleSeq := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_%d", MetricNameNextReceiveSequenceForChannel, 0),
		Help: fmt.Sprintf("Next Delivery Oracle sequence"),
	})
	ms[fmt.Sprintf("%s_%d", MetricNameNextReceiveSequenceForChannel, 0)] = nextReceiveOracleSeq
	prometheus.MustRegister(nextReceiveOracleSeq)

	// register gnfd -> bsc channels
	for _, c := range config.GreenfieldConfig.MonitorChannelList {
		nextSendSeq := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: fmt.Sprintf("%s_%d", MetricNameNextSendSequenceForChannel, c),
			Help: fmt.Sprintf("Next Send sequence for channel %d", c),
		})
		ms[fmt.Sprintf("%s_%d", MetricNameNextSendSequenceForChannel, c)] = nextSendSeq
		prometheus.MustRegister(nextSendSeq)

		nextReceiveSeq := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: fmt.Sprintf("%s_%d", MetricNameNextReceiveSequenceForChannel, c),
			Help: fmt.Sprintf("Next delivery sequence for channel %d", c),
		})
		ms[fmt.Sprintf("%s_%d", MetricNameNextReceiveSequenceForChannel, c)] = nextReceiveSeq
		prometheus.MustRegister(nextReceiveSeq)
	}

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

func (m *MetricService) SetNextSendSequenceForChannel(channel uint8, seq uint64) {
	m.MetricsMap[fmt.Sprintf("%s_%d", MetricNameNextSendSequenceForChannel, channel)].(prometheus.Gauge).Set(float64(seq))
}

func (m *MetricService) SetNextReceiveSequenceForChannel(channel uint8, seq uint64) {
	m.MetricsMap[fmt.Sprintf("%s_%d", MetricNameNextReceiveSequenceForChannel, channel)].(prometheus.Gauge).Set(float64(seq))
}
