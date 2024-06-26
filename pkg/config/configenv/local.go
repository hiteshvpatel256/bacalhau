//nolint:gomnd
package configenv

import (
	"os"
	"runtime"
	"time"

	"github.com/bacalhau-project/bacalhau/pkg/authn"
	"github.com/bacalhau-project/bacalhau/pkg/config/types"
	"github.com/bacalhau-project/bacalhau/pkg/logger"
	"github.com/bacalhau-project/bacalhau/pkg/model"
	"github.com/bacalhau-project/bacalhau/pkg/models"
)

var Local = types.BacalhauConfig{
	Metrics: types.MetricsConfig{
		Libp2pTracerPath: os.DevNull,
		EventTracerPath:  os.DevNull,
	},
	Update: types.UpdateConfig{
		SkipChecks: true,
	},
	Auth: types.AuthConfig{
		Methods: map[string]types.AuthenticatorConfig{
			"ClientKey": {
				Type: authn.MethodTypeChallenge,
			},
		},
	},
	Node: types.NodeConfig{
		NameProvider: "puuid",
		ClientAPI: types.APIConfig{
			Host: "0.0.0.0",
			Port: 1234,
		},
		ServerAPI: types.APIConfig{
			Host: "0.0.0.0",
			Port: 1234,
			TLS:  types.TLSConfiguration{},
		},
		Network: types.NetworkConfig{
			Type: models.NetworkTypeNATS,
			Port: 4222,
		},
		BootstrapAddresses:        []string{},
		DownloadURLRequestTimeout: types.Duration(300 * time.Second),
		VolumeSizeRequestTimeout:  types.Duration(2 * time.Minute),
		NodeInfoStoreTTL:          types.Duration(10 * time.Minute),
		DownloadURLRequestRetries: 3,
		LoggingMode:               logger.LogModeDefault,
		Type:                      []string{"requester"},
		AllowListedLocalPaths:     []string{},
		Labels:                    map[string]string{},
		DisabledFeatures: types.FeatureConfig{
			Engines:    []string{},
			Publishers: []string{},
			Storages:   []string{},
		},
		Libp2p: types.Libp2pConfig{
			SwarmPort:   1235,
			PeerConnect: "none",
		},
		IPFS: types.IpfsConfig{
			Connect:                "",
			PrivateInternal:        true,
			SwarmAddresses:         []string{},
			Profile:                "flatfs",
			SwarmListenAddresses:   []string{"/ip4/0.0.0.0/tcp/0"},
			GatewayListenAddresses: []string{"/ip4/0.0.0.0/tcp/0"},
			APIListenAddresses:     []string{"/ip4/0.0.0.0/tcp/0"},
		},
		Compute:   LocalComputeConfig,
		Requester: LocalRequesterConfig,
		WebUI: types.WebUIConfig{
			Enabled: false,
			Port:    8483,
		},
		StrictVersionMatch: false,
	},
}

var LocalComputeConfig = types.ComputeConfig{
	Capacity: types.CapacityConfig{
		IgnorePhysicalResourceLimits: false,
		TotalResourceLimits: models.ResourcesConfig{
			CPU:    "",
			Memory: "",
			Disk:   "",
			GPU:    "",
		},
		JobResourceLimits: models.ResourcesConfig{
			CPU:    "",
			Memory: "",
			Disk:   "",
			GPU:    "",
		},
		DefaultJobResourceLimits: models.ResourcesConfig{
			CPU:    "500m",
			Memory: "1Gb",
			Disk:   "",
			GPU:    "",
		},
	},
	ExecutionStore: types.JobStoreConfig{
		Type: types.BoltDB,
		Path: "",
	},
	JobTimeouts: types.JobTimeoutConfig{
		JobExecutionTimeoutClientIDBypassList: []string{},
		JobNegotiationTimeout:                 types.Duration(3 * time.Minute),
		MinJobExecutionTimeout:                types.Duration(500 * time.Millisecond),
		MaxJobExecutionTimeout:                types.Duration(model.NoJobTimeout),
		DefaultJobExecutionTimeout:            types.Duration(10 * time.Minute),
	},
	JobSelection: model.JobSelectionPolicy{
		Locality:            model.Anywhere,
		RejectStatelessJobs: false,
		AcceptNetworkedJobs: false,
		ProbeHTTP:           "",
		ProbeExec:           "",
	},
	Logging: types.LoggingConfig{
		LogRunningExecutionsInterval: types.Duration(10 * time.Second),
	},
	ManifestCache: types.DockerCacheConfig{
		Size:      1000,
		Duration:  types.Duration(1 * time.Hour),
		Frequency: types.Duration(1 * time.Hour),
	},
	LogStreamConfig: types.LogStreamConfig{
		ChannelBufferSize: 10,
	},
	LocalPublisher: types.LocalPublisherConfig{
		Address: "127.0.0.1",
		Port:    6001,
	},
	ControlPlaneSettings: types.ComputeControlPlaneConfig{
		InfoUpdateFrequency:     types.Duration(60 * time.Second),
		ResourceUpdateFrequency: types.Duration(30 * time.Second),
		HeartbeatFrequency:      types.Duration(15 * time.Second),
		HeartbeatTopic:          "heartbeat",
	},
}

var LocalRequesterConfig = types.RequesterConfig{
	ExternalVerifierHook: "",
	JobSelectionPolicy: model.JobSelectionPolicy{
		Locality:            model.Anywhere,
		RejectStatelessJobs: false,
		AcceptNetworkedJobs: false,
		ProbeHTTP:           "",
		ProbeExec:           "",
	},
	JobStore: types.JobStoreConfig{
		Type: types.BoltDB,
		Path: "",
	},
	HousekeepingBackgroundTaskInterval: types.Duration(30 * time.Second),
	NodeRankRandomnessRange:            5,
	OverAskForBidsFactor:               3,
	FailureInjectionConfig: model.FailureInjectionRequesterConfig{
		IsBadActor: false,
	},
	EvaluationBroker: types.EvaluationBrokerConfig{
		EvalBrokerVisibilityTimeout:    types.Duration(60 * time.Second),
		EvalBrokerInitialRetryDelay:    types.Duration(1 * time.Second),
		EvalBrokerSubsequentRetryDelay: types.Duration(30 * time.Second),
		EvalBrokerMaxRetryCount:        10,
	},
	Worker: types.WorkerConfig{
		WorkerCount:                  runtime.NumCPU(),
		WorkerEvalDequeueTimeout:     types.Duration(5 * time.Second),
		WorkerEvalDequeueBaseBackoff: types.Duration(1 * time.Second),
		WorkerEvalDequeueMaxBackoff:  types.Duration(30 * time.Second),
	},
	Scheduler: types.SchedulerConfig{
		QueueBackoff:               types.Duration(30 * time.Second),
		NodeOverSubscriptionFactor: 1.5,
	},
	JobDefaults: types.JobDefaults{
		TotalTimeout: types.Duration(30 * time.Minute),
	},
	StorageProvider: types.StorageProviderConfig{
		S3: types.S3StorageProviderConfig{
			PreSignedURLExpiration: types.Duration(30 * time.Minute),
		},
	},
	ControlPlaneSettings: types.RequesterControlPlaneConfig{
		HeartbeatCheckFrequency: types.Duration(30 * time.Second),
		HeartbeatTopic:          "heartbeat",
		NodeDisconnectedAfter:   types.Duration(30 * time.Second),
	},
}
