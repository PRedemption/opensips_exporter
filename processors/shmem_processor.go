package processors

import (
	"github.com/VoIPGRID/opensips_exporter/opensips"
	"github.com/prometheus/client_golang/prometheus"
)

// shmemProcessor provices metrics about shared memory
// doc: http://www.opensips.org/Documentation/Interface-CoreStatistics-1-11#toc21
// src: https://github.com/OpenSIPS/opensips/blob/1.11/mem/shm_mem.c#L52
type shmemProcessor struct {
	statistics map[string]opensips.Statistic
}

var shmemLabelNames = []string{}
var shmemMetrics = map[string]metric{
	"total_size":     newMetric("shmem", "total_size", "Total size of shared memory available to OpenSIPS processes.", shmemLabelNames, prometheus.GaugeValue),
	"used_size":      newMetric("shmem", "used_size", "Amount of shared memory requested and used by OpenSIPS processes.", shmemLabelNames, prometheus.GaugeValue),
	"real_used_size": newMetric("shmem", "real_used_size", "Amount of shared memory requested by OpenSIPS processes + malloc overhead", shmemLabelNames, prometheus.GaugeValue),
	"max_used_size":  newMetric("shmem", "max_used_size", "Maximum amount of shared memory ever used by OpenSIPS processes.", shmemLabelNames, prometheus.GaugeValue),
	"free_size":      newMetric("shmem", "free_size", "Free memory available. Computed as total_size - real_used_size", shmemLabelNames, prometheus.GaugeValue),
	"fragments":      newMetric("shmem", "fragments", "Total number of fragments in the shared memory.", shmemLabelNames, prometheus.GaugeValue),
}

func init() {
	for metric := range shmemMetrics {
		OpensipsProcessors[metric] = shmemProcessorFunc
	}
	OpensipsProcessors["shmem:"] = shmemProcessorFunc
}

// Describe implements prometheus.Collector.
func (p shmemProcessor) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range shmemMetrics {
		ch <- metric.Desc
	}
}

// Collect implements prometheus.Collector.
func (p shmemProcessor) Collect(ch chan<- prometheus.Metric) {
	for _, s := range p.statistics {
		if s.Module == "shmem" {
			switch s.Name {
			case "total_size":
				ch <- prometheus.MustNewConstMetric(
					shmemMetrics["total_size"].Desc,
					shmemMetrics["total_size"].ValueType,
					s.Value,
				)
			case "used_size":
				ch <- prometheus.MustNewConstMetric(
					shmemMetrics["used_size"].Desc,
					shmemMetrics["used_size"].ValueType,
					s.Value,
				)
			case "real_used_size":
				ch <- prometheus.MustNewConstMetric(
					shmemMetrics["real_used_size"].Desc,
					shmemMetrics["real_used_size"].ValueType,
					s.Value,
				)
			case "max_used_size":
				ch <- prometheus.MustNewConstMetric(
					shmemMetrics["max_used_size"].Desc,
					shmemMetrics["max_used_size"].ValueType,
					s.Value,
				)
			case "free_size":
				ch <- prometheus.MustNewConstMetric(
					shmemMetrics["free_size"].Desc,
					shmemMetrics["free_size"].ValueType,
					s.Value,
				)
			case "fragments":
				ch <- prometheus.MustNewConstMetric(
					shmemMetrics["fragments"].Desc,
					shmemMetrics["fragments"].ValueType,
					s.Value,
				)
			}
		}
	}
}

func shmemProcessorFunc(s map[string]opensips.Statistic) prometheus.Collector {
	return &shmemProcessor{
		statistics: s,
	}
}
