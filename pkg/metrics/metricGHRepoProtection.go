package metrics

import (
	"cx-installer/pkg/constants"
)

const (
	tagKeyPrefix = "tag_"
	
	dataKeyInstallationCxVersion          = "cx_version"
	dataKeyInstallationCxCommand          = "cx_command"
	dataKeyInstallationCxStatus           = "cx_status"
	dataKeyInstallationCxConfigRepository = "cx_config_repository"
	dataKeyInstallationOpenshiftNamespace = "openshift_namespace"
	
	tagKeyInstallationCxVersion          = tagKeyPrefix + dataKeyInstallationCxVersion
	tagKeyInstallationCxCommand          = tagKeyPrefix + dataKeyInstallationCxCommand
	tagKeyInstallationCxStatus           = tagKeyPrefix + dataKeyInstallationCxStatus
	tagKeyInstallationCxConfigRepository = tagKeyPrefix + dataKeyInstallationCxConfigRepository
	tagKeyInstallationOpenshiftNamespace = tagKeyPrefix + dataKeyInstallationOpenshiftNamespace
	
	metricNameCxInstallation = "cx_installer"
)

type MetricInstallation struct {
	CxCommand          string
	InstallStatus      string
	CxConfigRepository string
	TargetNameSpace    string
}

func (metric MetricInstallation) WriteMetric() {
	fillMetricInstallComponent(metric).WriteMetric()
}

func fillMetricInstallComponent(metric MetricInstallation) (metrics InfluxDB) {
	metrics = InfluxDB{
		MetricName: metricNameCxInstallation,
		MetricsData: map[string]interface{}{
			dataKeyInstallationCxVersion:          constants.CxVersion,
			dataKeyInstallationCxCommand:          metric.CxCommand,
			dataKeyInstallationCxStatus:           metric.InstallStatus,
			dataKeyInstallationCxConfigRepository: metric.CxConfigRepository,
			dataKeyInstallationOpenshiftNamespace: metric.TargetNameSpace,
		},
		Tags: map[string]string{
			tagKeyInstallationCxVersion:          constants.CxVersion,
			tagKeyInstallationCxCommand:          metric.CxCommand,
			tagKeyInstallationCxStatus:           metric.InstallStatus,
			tagKeyInstallationCxConfigRepository: metric.CxConfigRepository,
			tagKeyInstallationOpenshiftNamespace: metric.TargetNameSpace,
		},
	}
	return metrics
}
