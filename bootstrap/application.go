package bootstrap

import (
	contract "coder/contracts"
	"coder/providers"
)

var (
	app = NewApplication()
)
func Bootstrap(){
	app.Register(providers.NewDatabase())
	app.Boot()
	app.AfterBoot()
}
func GetApp() *Application {
	return &app
}
type Application struct {
	boots []contract.ApplicationBoot
}

func (a *Application) AfterBoot()  {
	for i:=0; i < len(a.boots); i++ {
		a.boots[i].AfterBoot()
	}
}
func (a *Application) Boot()  {
	for i:=0; i < len(a.boots); i++ {
		a.boots[i].Boot()
	}
}
func (a *Application) Register(boot contract.ApplicationBoot)  {
	a.boots = append(a.boots, boot)
}
func (a *Application) Shutdown()  {
	for i:=0; i < len(a.boots); i++ {
		a.boots[i].Shutdown()
	}
}
func NewApplication() Application {
	return Application{}
}
