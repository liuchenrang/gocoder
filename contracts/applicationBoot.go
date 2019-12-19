package contract


type ApplicationBoot interface {
	Boot()
	AfterBoot()
	Shutdown()
}
