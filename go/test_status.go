/*
 * Chaos Agent
 *
 * API is used to create load for other APIs
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TestStatus struct {

	// Name of the test collection that is being executed
	TestCollectionName string `json:"testCollectionName,omitempty"`

	// The total requests sent to the remote API
	RequestsExecuted int32 `json:"requestsExecuted,omitempty"`

	TransactionsPerSecond int32 `json:"transactionsPerSecond,omitempty"`

	// Execution time in seconds
	ExecutionTime int32 `json:"executionTime,omitempty"`

	// CPU usage on the host that the agent is running on
	Cpu float64 `json:"cpu,omitempty"`

	// The number of users being simulated
	SimulatedUsers int32 `json:"simulatedUsers,omitempty"`
}
