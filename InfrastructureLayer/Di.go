package InfrastructureLayer

type ConnectToBd struct {
	Re Repositorys
}

func NewUserService(Rep Repositorys) *ConnectToBd {
	return &ConnectToBd{Re: Rep}
}
