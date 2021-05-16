package confighandler

import (
	"fmt"
	"restoros/configurationmanager"
)

type ConfigInitHandler struct {
	Manager     configurationmanager.IManager
	RepoManager configurationmanager.IRepositoryManager
}

func (handler *ConfigInitHandler) Handle(args []string) error {
	if handler.Manager.IsConfigurationInitialized() {
		return fmt.Errorf("already initialized")
	}

	err := handler.RepoManager.Initialize()
	if err != nil {
		return err
	}

	initialConfiguration := &configurationmanager.RestorosConfiguration{}
	return handler.Manager.Write(initialConfiguration)
}
