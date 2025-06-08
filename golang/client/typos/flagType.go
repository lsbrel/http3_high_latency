package typos

type UsecasesFlags string

const (
	LatencyDisabled UsecasesFlags = "nolatency"
	LatencyEnabled  UsecasesFlags = "latency"
)
