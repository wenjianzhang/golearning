package pipe_filter

type Request interface {
}

type Response interface {
}

type Filter interface {
	Process(data Request) (Response, error)
}
