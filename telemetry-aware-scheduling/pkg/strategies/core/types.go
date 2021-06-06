//Contains the behaviour shared between all Violable and Enforceable strategies from a Telemetry Policy.

package core

import (
	"time"

	"github.com/intel/telemetry-aware-scheduling/telemetry-aware-scheduling/pkg/cache"
)

//Interface describes expected behavior of a specific strategy.
type Interface interface {
	Violated(cache cache.Reader) map[string]interface{}
	StrategyType() string
	Equals(Interface) bool
	GetPolicyName() string
	SetPolicyName(string)
}

//Enforcer registers strategies by type, adds specific strategies to a registry, and Enforces those strategies.
type Enforcer interface {
	RegisterStrategyType(strategy Interface)
	UnregisterStrategyType(strategy Interface)
	RegisteredStrategyTypes() []string
	IsRegistered(string) bool
	AddStrategy(Enforceable)
	RemoveStrategy(Enforceable)
	EnforceRegisteredStrategies(cache.Reader, time.Ticker)
}



type Enforceable interface {
	Interface
	Enforce(enforcer *MetricEnforcer, cache cache.Reader) (int, error)
	Cleanup() error
}

