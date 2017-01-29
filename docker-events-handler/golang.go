package main

import (
"encoding/json"
"fmt"
"io"
"log"

"github.com/docker/engine-api/client"
"github.com/docker/engine-api/types"
"github.com/docker/engine-api/types/events"
"golang.org/x/net/context"
	"github.com/docker/engine-api/types/filters"
)


func main() {

	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.24", nil, nil)
	//cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	// more about list containers https://docs.docker.com/engine/reference/commandline/ps/
	options := types.ContainerListOptions{}
	containers, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		panic(err)
	}

	for _, c := range containers {
		fmt.Printf("%s - %s\n", c.Names[0], c.ID)
	}

	// more about events here: http://docs-stage.docker.com/v1.10/engine/reference/commandline/events/
	var filters = filters.NewArgs()

	// events to register a container
	filters.Add("event", "exec_create")
	filters.Add("event", "exec_start")
	filters.Add("event", "create")
	filters.Add("event", "start")

	// events to unregister a container
	filters.Add("event", "destroy")
	filters.Add("event", "die")
	filters.Add("event", "kill")
	filters.Add("event", "stop")

	body, err := cli.Events(context.Background(), types.EventsOptions{ Filters: filters })
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(body)
	for {
		var event events.Message
		err := dec.Decode(&event)
		if err != nil && err == io.EOF {
			break
		}

		log.Println(event)
	}
}