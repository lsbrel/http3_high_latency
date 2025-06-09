package typos

type ServerFlag string

const (
	HTTP1 ServerFlag = "http1"
	HTTP2 ServerFlag = "http2"
	HTTP3 ServerFlag = "http3"
)
