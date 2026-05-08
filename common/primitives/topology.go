package primitives

// ServiceTopology describes how Temporal services are distributed across processes.
// Components can take shortcuts when all default services are co-located in this process.
type ServiceTopology int

const (
	// ServiceTopologyDistributed indicates that services are split across processes.
	ServiceTopologyDistributed ServiceTopology = iota
	// ServiceTopologySingleProcess indicates that all default services
	// (frontend, history, matching, worker) are running in this process.
	ServiceTopologySingleProcess
)

// NewServiceTopology returns ServiceTopologySingleProcess when the given set
// contains all default Temporal services, otherwise ServiceTopologyDistributed.
func NewServiceTopology(serviceNames map[ServiceName]struct{}) ServiceTopology {
	required := []ServiceName{
		FrontendService,
		HistoryService,
		MatchingService,
		WorkerService,
	}
	for _, name := range required {
		if _, ok := serviceNames[name]; !ok {
			return ServiceTopologyDistributed
		}
	}
	return ServiceTopologySingleProcess
}
