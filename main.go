package main

import (
	"log"
	"os"

	"ldocker/docker"
	"ldocker/panel"
	"ldocker/terminal"
)

func main() {
	defer func() {
		f, _ := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		log.SetOutput(f)
		err := recover().(error)
		log.Println(err.Error())
	}()

	dockerExecutor := docker.NewExecutor()

	t, err := terminal.NewTerminal(dockerExecutor)
	if err != nil {
		log.Println(err)
	}

	inspectPanel := panel.NewInspectPanel()

	var commandPanel *panel.CommandPanel

	resourcePanel := panel.NewResourcePanel(
		func(resource string) {

			selectedCmd := commandPanel.GetSelectedCommand()

			_, inspectLines := dockerExecutor.ExecuteCommand(
				docker.InspectCommand{
					T: selectedCmd.Type(),
					R: resource,
				},
			)

			inspectPanel.UpdateContent(inspectLines)
		},
	)

	commandPanel = panel.NewCommandPanel(
		func(command docker.ListCommand) {
			resources, lines := dockerExecutor.ExecuteCommand(command)
			resourcePanel.UpdateContent(resources, lines)
		},
	)

	t.AddPanel(commandPanel)
	t.AddPanel(resourcePanel)
	t.AddPanel(inspectPanel)

	initialCmd := commandPanel.GetSelectedCommand()
	resources, lines := dockerExecutor.ExecuteCommand(initialCmd)
	resourcePanel.UpdateContent(resources, lines)

	_, inspect := dockerExecutor.ExecuteCommand(
		docker.InspectCommand{
			T: initialCmd.Type(),
			R: resources[0],
		},
	)
	inspectPanel.UpdateContent(inspect)

	t.Start()
}
