package config

import (
	"encoding/json"
	"sync"
	"time"
)

// MockConfig will respond with whatever config it's set to do during
// initialization
type MockConfig struct {
	Callbacks                            []func()
	GetAPIKeysErr                        error
	GetAPIKeysVal                        []string
	GetCollectorTypeErr                  error
	GetCollectorTypeVal                  string
	GetInMemoryCollectorCacheCapacityErr error
	GetInMemoryCollectorCacheCapacityVal InMemoryCollectorCacheCapacity
	GetHoneycombAPIErr                   error
	GetHoneycombAPIVal                   string
	GetListenAddrErr                     error
	GetListenAddrVal                     string
	GetPeerListenAddrErr                 error
	GetPeerListenAddrVal                 string
	GetCompressPeerCommunicationsVal     bool
	GetGRPCListenAddrErr                 error
	GetGRPCListenAddrVal                 string
	GetLoggerTypeErr                     error
	GetLoggerTypeVal                     string
	GetHoneycombLoggerConfigErr          error
	GetHoneycombLoggerConfigVal          HoneycombLoggerConfig
	GetLoggingLevelErr                   error
	GetLoggingLevelVal                   string
	GetOtherConfigErr                    error
	// GetOtherConfigVal must be a JSON representation of the config struct to be populated.
	GetOtherConfigVal             string
	GetPeersErr                   error
	GetPeersVal                   []string
	GetRedisHostErr               error
	GetRedisHostVal               string
	GetRedisUsernameErr           error
	GetRedisUsernameVal           string
	GetRedisPasswordErr           error
	GetRedisPasswordVal           string
	GetUseTLSErr                  error
	GetUseTLSVal                  bool
	GetUseTLSInsecureErr          error
	GetUseTLSInsecureVal          bool
	GetSamplerTypeErr             error
	GetSamplerTypeName            string
	GetSamplerTypeVal             interface{}
	GetMetricsTypeErr             error
	GetMetricsTypeVal             string
	GetHoneycombMetricsConfigErr  error
	GetHoneycombMetricsConfigVal  HoneycombMetricsConfig
	GetPrometheusMetricsConfigErr error
	GetPrometheusMetricsConfigVal PrometheusMetricsConfig
	GetSendDelayErr               error
	GetSendDelayVal               time.Duration
	GetBatchTimeoutVal            time.Duration
	GetTraceTimeoutErr            error
	GetTraceTimeoutVal            time.Duration
	GetMaxBatchSizeVal            uint
	GetUpstreamBufferSizeVal      int
	GetPeerBufferSizeVal          int
	SendTickerVal                 time.Duration
	IdentifierInterfaceName       string
	UseIPV6Identifier             bool
	RedisIdentifier               string
	PeerManagementType            string
	DebugServiceAddr              string
	DryRun                        bool
	DryRunFieldName               string
	AddHostMetadataToTrace        bool
	AddRuleReasonToTrace          bool
	EnvironmentCacheTTL           time.Duration
	DatasetPrefix                 string
	QueryAuthToken                string
	GRPCMaxConnectionIdle         time.Duration
	GRPCMaxConnectionAge          time.Duration
	GRPCMaxConnectionAgeGrace     time.Duration
	GRPCTime                      time.Duration
	GRPCTimeout                   time.Duration
	PeerTimeout                   time.Duration
	AdditionalErrorFields         []string
	AddSpanCountToRoot            bool
	CacheOverrunStrategy          string
	CfgMetadata                   []ConfigMetadata

	Mux sync.RWMutex
}

func (m *MockConfig) ReloadConfig() {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	for _, callback := range m.Callbacks {
		callback()
	}
}

func (m *MockConfig) RegisterReloadCallback(callback func()) {
	m.Mux.Lock()
	m.Callbacks = append(m.Callbacks, callback)
	m.Mux.Unlock()
}

func (m *MockConfig) GetAPIKeys() ([]string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetAPIKeysVal, m.GetAPIKeysErr
}

func (m *MockConfig) GetCollectorType() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetCollectorTypeVal, m.GetCollectorTypeErr
}

func (m *MockConfig) GetInMemCollectorCacheCapacity() (InMemoryCollectorCacheCapacity, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetInMemoryCollectorCacheCapacityVal, m.GetInMemoryCollectorCacheCapacityErr
}

func (m *MockConfig) GetHoneycombAPI() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetHoneycombAPIVal, m.GetHoneycombAPIErr
}

func (m *MockConfig) GetListenAddr() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetListenAddrVal, m.GetListenAddrErr
}

func (m *MockConfig) GetPeerListenAddr() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetPeerListenAddrVal, m.GetPeerListenAddrErr
}

func (m *MockConfig) GetCompressPeerCommunication() bool {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetCompressPeerCommunicationsVal
}

func (m *MockConfig) GetGRPCListenAddr() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetGRPCListenAddrVal, m.GetGRPCListenAddrErr
}

func (m *MockConfig) GetLoggerType() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetLoggerTypeVal, m.GetLoggerTypeErr
}

func (m *MockConfig) GetHoneycombLoggerConfig() (HoneycombLoggerConfig, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetHoneycombLoggerConfigVal, m.GetHoneycombLoggerConfigErr
}

func (m *MockConfig) GetLoggingLevel() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetLoggingLevelVal, m.GetLoggingLevelErr
}

func (m *MockConfig) GetOtherConfig(name string, iface interface{}) error {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	err := json.Unmarshal([]byte(m.GetOtherConfigVal), iface)
	if err != nil {
		return err
	}
	return m.GetOtherConfigErr
}

func (m *MockConfig) GetPeers() ([]string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetPeersVal, m.GetPeersErr
}

func (m *MockConfig) GetRedisHost() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetRedisHostVal, m.GetRedisHostErr
}

func (m *MockConfig) GetRedisUsername() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetRedisUsernameVal, m.GetRedisUsernameErr
}

func (m *MockConfig) GetRedisPassword() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetRedisPasswordVal, m.GetRedisPasswordErr
}

func (m *MockConfig) GetUseTLS() (bool, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetUseTLSVal, m.GetUseTLSErr
}

func (m *MockConfig) GetUseTLSInsecure() (bool, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetUseTLSInsecureVal, m.GetUseTLSInsecureErr
}

func (m *MockConfig) GetMetricsType() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetMetricsTypeVal, m.GetMetricsTypeErr
}

func (m *MockConfig) GetHoneycombMetricsConfig() (HoneycombMetricsConfig, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetHoneycombMetricsConfigVal, m.GetHoneycombMetricsConfigErr
}

func (m *MockConfig) GetPrometheusMetricsConfig() (PrometheusMetricsConfig, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetPrometheusMetricsConfigVal, m.GetPrometheusMetricsConfigErr
}

func (m *MockConfig) GetSendDelay() (time.Duration, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetSendDelayVal, m.GetSendDelayErr
}

func (m *MockConfig) GetBatchTimeout() time.Duration {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetBatchTimeoutVal
}

func (m *MockConfig) GetTraceTimeout() (time.Duration, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetTraceTimeoutVal, m.GetTraceTimeoutErr
}

func (m *MockConfig) GetMaxBatchSize() uint {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetMaxBatchSizeVal
}

// TODO: allow per-dataset mock values
func (m *MockConfig) GetSamplerConfigForDataset(dataset string) (interface{}, string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetSamplerTypeVal, m.GetSamplerTypeName, m.GetSamplerTypeErr
}

// GetAllSamplerRules returns all dataset rules, including the default
func (m *MockConfig) GetAllSamplerRules() (map[string]interface{}, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	v := map[string]interface{}{"dataset1": m.GetSamplerTypeVal}
	return v, m.GetSamplerTypeErr
}

func (m *MockConfig) GetUpstreamBufferSize() int {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetUpstreamBufferSizeVal
}

func (m *MockConfig) GetPeerBufferSize() int {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.GetPeerBufferSizeVal
}

func (m *MockConfig) GetIdentifierInterfaceName() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.IdentifierInterfaceName, nil
}

func (m *MockConfig) GetUseIPV6Identifier() (bool, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.UseIPV6Identifier, nil
}

func (m *MockConfig) GetRedisIdentifier() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.RedisIdentifier, nil
}

func (m *MockConfig) GetSendTickerValue() time.Duration {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.SendTickerVal
}

func (m *MockConfig) GetPeerManagementType() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.PeerManagementType, nil
}

func (m *MockConfig) GetDebugServiceAddr() (string, error) {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.DebugServiceAddr, nil
}

func (m *MockConfig) GetIsDryRun() bool {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.DryRun
}

func (m *MockConfig) GetDryRunFieldName() string {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.DryRunFieldName
}

func (m *MockConfig) GetAddHostMetadataToTrace() bool {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.AddHostMetadataToTrace
}

func (m *MockConfig) GetAddRuleReasonToTrace() bool {
	m.Mux.RLock()
	defer m.Mux.RUnlock()

	return m.AddRuleReasonToTrace
}

func (f *MockConfig) GetEnvironmentCacheTTL() time.Duration {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	return f.EnvironmentCacheTTL
}

func (f *MockConfig) GetDatasetPrefix() string {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	return f.DatasetPrefix
}

func (f *MockConfig) GetQueryAuthToken() string {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	return f.QueryAuthToken
}

func (f *MockConfig) GetGRPCMaxConnectionIdle() time.Duration {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	return f.GRPCMaxConnectionIdle
}

func (f *MockConfig) GetGRPCMaxConnectionAge() time.Duration {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	return f.GRPCMaxConnectionAge
}

func (f *MockConfig) GetGRPCMaxConnectionAgeGrace() time.Duration {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	return f.GRPCMaxConnectionAgeGrace
}

func (f *MockConfig) GetGRPCTime() time.Duration {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	return f.GRPCTime
}

func (f *MockConfig) GetGRPCTimeout() time.Duration {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	return f.GRPCTimeout
}

func (f *MockConfig) GetPeerTimeout() time.Duration {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	return f.PeerTimeout
}

func (f *MockConfig) GetAdditionalErrorFields() []string {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	return f.AdditionalErrorFields
}

func (f *MockConfig) GetAddSpanCountToRoot() bool {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	return f.AddSpanCountToRoot
}

func (f *MockConfig) GetCacheOverrunStrategy() string {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	return f.CacheOverrunStrategy
}

func (f *MockConfig) GetConfigMetadata() []ConfigMetadata {
	f.Mux.RLock()
	defer f.Mux.RUnlock()

	return f.CfgMetadata
}
