package bootstrap

// ServiceInfo 当前服务信息
type ServiceInfo struct {
	Name    string
	Version string
	Id      string
}

func NewServiceInfo(name, version, id string) ServiceInfo {
	return ServiceInfo{
		Name:    name,
		Version: version,
		Id:      id,
	}
}
