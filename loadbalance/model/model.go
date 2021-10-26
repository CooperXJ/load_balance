package model

// EtcdConfig  is configuration for EtcdServer
type EtcdConfig struct {
	EtcdEndpoints      []string `json:"etcd_endpoints"`
	TlsDisabled        bool     `json:"tls_disabled"`
	EtcdTLSCertPath    string   `json:"etcd_tls_cert_path" `
	EtcdTLSCertKeyPath string   `json:"etcd_tls_cert_key_path"`
	EtcdTLSCAPath      string   `json:"etcd_tls_ca_path"`
}

type ServiceDefMap struct {
	ServiceEndpoint string `json:"service_endpoint"`
	ServiceName     string `json:"service_name"`
	ServiceHost     string `json:"service_host"`
}
