package prometheusrules

import (
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"

	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
)

func init() {
	PrometheusRuleFactories = append(PrometheusRuleFactories, NewSystemSidekiqPrometheusRuleFactory)
}

type SystemSidekiqPrometheusRuleFactory struct {
}

func NewSystemSidekiqPrometheusRuleFactory() PrometheusRuleFactory {
	return &SystemSidekiqPrometheusRuleFactory{}
}

func (s *SystemSidekiqPrometheusRuleFactory) Type() string {
	return "system-sidekiq"
}

func (s *SystemSidekiqPrometheusRuleFactory) PrometheusRule() *monitoringv1.PrometheusRule {
	options, err := systemOptions()
	if err != nil {
		panic(err)
	}
	return component.NewSystem(options).SystemSidekiqPrometheusRules()
}
