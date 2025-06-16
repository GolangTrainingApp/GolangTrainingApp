package server

type WindyAPI interface {
	GetSingleResponseByLatAndLong() (string, error)
}
