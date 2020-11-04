package metrics

import (
	"cx-installer/pkg/versions"
	"strconv"
)

const (
	tagKeyPrefix = "tag_"
	
	dataKeyCliVersion          = "cli_version"
	dataKeyCliCommand          = "cli_command"
	dataKeyGHProtectionActive  = "gh_protection_active"
	dataKeyGHPullRequestActive = "gh_pr_active"
	dataKeyGHStatusCheckActive = "gh_status_check_active"
	
	tagKeyCliVersion          = tagKeyPrefix + dataKeyCliVersion
	tagKeyCliCommand          = tagKeyPrefix + dataKeyCliCommand
	tagKeyGHProtectionActive  = tagKeyPrefix + dataKeyGHProtectionActive
	tagKeyGHPullRequestActive = tagKeyPrefix + dataKeyGHPullRequestActive
	tagKeyGHStatusCheckActive = tagKeyPrefix + dataKeyGHStatusCheckActive
	
	measurementName = "gh_branch_protection"
)

type MetricGHRepositoryProtection struct {
	CliCommand          string
	GhProtectionActive  bool
	GhPullrequestActive bool
	GhStatusCheckActive bool
}

func (metric MetricGHRepositoryProtection) WriteMetric() {
	fillMetricInstallComponent(metric).WriteMetric()
}

func fillMetricInstallComponent(metric MetricGHRepositoryProtection) (metrics InfluxDB) {
	metrics = InfluxDB{
		MetricName: measurementName,
		MetricsData: map[string]interface{}{
			dataKeyCliVersion:          versions.VersionFromGit,
			dataKeyCliCommand:          metric.CliCommand,
			dataKeyGHProtectionActive:  metric.GhProtectionActive,
			dataKeyGHPullRequestActive: metric.GhPullrequestActive,
			dataKeyGHStatusCheckActive: metric.GhStatusCheckActive,
		},
		Tags: map[string]string{
			tagKeyCliVersion:          versions.VersionFromGit,
			tagKeyCliCommand:          metric.CliCommand,
			tagKeyGHProtectionActive:  strconv.FormatBool(metric.GhProtectionActive),
			tagKeyGHPullRequestActive: strconv.FormatBool(metric.GhPullrequestActive),
			tagKeyGHStatusCheckActive: strconv.FormatBool(metric.GhStatusCheckActive),
		},
	}
	return metrics
}
