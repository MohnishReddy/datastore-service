package constants

type ServiceMode string

const (
	TestMode         ServiceMode = "test"
	DataStorePodMode ServiceMode = "data-store"
	LoadBalancerMode ServiceMode = "load-balancer"
)
