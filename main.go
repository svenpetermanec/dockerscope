package main

import (
	"log"
	"os"

	"ldocker/docker"
	"ldocker/panel"
	"ldocker/terminal"
)

func main() {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")

	defer func() {
		err := recover().(error)
		log.Println(err.Error())
	}()

	dockerExecutor := docker.NewExecutor()

	t, err := terminal.NewTerminal(dockerExecutor)
	if err != nil {
		log.Println(err)
	}

	// w, h := screen.Size()
	inspectPanel := panel.NewInspectPanel(100, 1, 80)

	var commandPanel *panel.CommandPanel

	resourcePanel := panel.NewResourcePanel(
		15, 1, 80, func(resource string) {

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
		1, 1, 10, func(command docker.ListCommand) {
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
