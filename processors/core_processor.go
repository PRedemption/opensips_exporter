package processors

import (
	"github.com/VoIPGRID/opensips_exporter/opensips"
	"github.com/prometheus/client_golang/prometheus"
)

// coreProcessor provides basic metrics,
// doc: http://www.opensips.org/Documentation/Interface-CoreStatistics-1-11#toc1
// src: https://github.com/OpenSIPS/opensips/blob/1.11/core_stats.h
type coreProcessor struct {
	statistics map[string]opensips.Statistic
}

var coreMetrics = map[string]metric{
	"rcv_requests":        newMetric("core", "requests_total", "Total number of received requests by OpenSIPS.", []string{}, prometheus.CounterValue),
	"rcv_replies":         newMetric("core", "replies_total", "Total number of received replies by OpenSIPS.", []string{}, prometheus.CounterValue),
	"fwd_requests":        newMetric("core", "requests", "Number of requests by OpenSIPS.", []string{"kind"}, prometheus.CounterValue),
	"fwd_replies":         newMetric("core", "replies", "Number of received replies by OpenSIPS.", []string{"kind"}, prometheus.CounterValue),
	"drop_requests":       newMetric("core", "requests", "Number of requests by OpenSIPS.", []string{"kind"}, prometheus.CounterValue),
	"drop_replies":        newMetric("core", "replies", "Number of received replies by OpenSIPS.", []string{"kind"}, prometheus.CounterValue),
	"err_requests":        newMetric("core", "requests", "Number of requests by OpenSIPS.", []string{"kind"}, prometheus.CounterValue),
	"err_replies":         newMetric("core", "replies", "Number of received replies by OpenSIPS.", []string{"kind"}, prometheus.CounterValue),
	"bad_URIs_rcvd":       newMetric("core", "bad_URIs_rcvd", "Number of URIs that OpenSIPS failed to parse.", []string{}, prometheus.CounterValue),
	"unsupported_methods": newMetric("core", "unsupported_methods", "Number of non-standard methods encountered by OpenSIPS while parsing SIP methods.", []string{}, prometheus.CounterValue),
	"bad_msg_hdr":         newMetric("core", "bad_msg_hdr", "Number of SIP headers that OpenSIPS failed to parse.", []string{}, prometheus.CounterValue),
	"timestamp":           newMetric("core", "uptime_seconds", "Number of seconds elapsed from OpenSIPS starting.", []string{}, prometheus.CounterValue),
}

func init() {
	for metric := range coreMetrics {
		OpensipsProcessors[metric] = coreProcessorFunc
	}
	OpensipsProcessors["core:"] = coreProcessorFunc
}

func coreProcessorFunc(s map[string]opensips.Statistic) prometheus.Collector {
	return &coreProcessor{
		statistics: s,
	}
}

// Describe implements prometheus.Collector.
func (p coreProcessor) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range coreMetrics {
		ch <- metric.Desc
	}
}

// Collect implements prometheus.Collector.
func (p coreProcessor) Collect(ch chan<- prometheus.Metric) {
	for _, s := range p.statistics {
		if s.Module == "core" {
			switch s.Name {
			case "rcv_requests":
				ch <- prometheus.MustNewConstMetric(
					coreMetrics["rcv_requests"].Desc,
					coreMetrics["rcv_requests"].ValueType,
					s.Value,
				)
			case "rcv_replies":
				ch <- prometheus.MustNewConstMetric(
					coreMetrics["rcv_replies"].Desc,
					coreMetrics["rcv_replies"].ValueType,
					s.Value,
				)
			case "fwd_requests":
				ch <- prometheus.MustNewConstMetric(
					coreMetrics["fwd_requests"].Desc,
					coreMetrics["fwd_requests"].ValueType,
					s.Value,
					"forwarded",
				)
			case "fwd_replies":
				ch <- prometheus.MustNewConstMetric(
					coreMetrics["fwd_replies"].Desc,
					coreMetrics["fwd_replies"].ValueType,
					s.Value,
					"forwarded",
				)
			case "drop_requests":
				ch <- prometheus.MustNewConstMetric(
					coreMetrics["drop_requests"].Desc,
					coreMetrics["drop_requests"].ValueType,
					s.Value,
					"dropped",
				)
			case "drop_replies":
				ch <- prometheus.MustNewConstMetric(
					coreMetrics["drop_replies"].Desc,
					coreMetrics["drop_replies"].ValueType,
					s.Value,
					"dropped",
				)
			case "err_requests":
				ch <- prometheus.MustNewConstMetric(
					coreMetrics["err_requests"].Desc,
					coreMetrics["err_requests"].ValueType,
					s.Value,
					"error",
				)
			case "err_replies":
				ch <- prometheus.MustNewConstMetric(
					coreMetrics["err_replies"].Desc,
					coreMetrics["err_replies"].ValueType,
					s.Value,
					"error",
				)
			case "bad_URIs_rcvd":
				ch <- prometheus.MustNewConstMetric(
					coreMetrics["bad_URIs_rcvd"].Desc,
					coreMetrics["bad_URIs_rcvd"].ValueType,
					s.Value,
				)
			case "unsupported_methods":
				ch <- prometheus.MustNewConstMetric(
					coreMetrics["unsupported_methods"].Desc,
					coreMetrics["unsupported_methods"].ValueType,
					s.Value,
				)
			case "bad_msg_hdr":
				ch <- prometheus.MustNewConstMetric(
					coreMetrics["bad_msg_hdr"].Desc,
					coreMetrics["bad_msg_hdr"].ValueType,
					s.Value,
				)
			case "timestamp":
				ch <- prometheus.MustNewConstMetric(
					coreMetrics["timestamp"].Desc,
					coreMetrics["timestamp"].ValueType,
					s.Value,
				)
			}
		}
	}
}
