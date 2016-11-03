package main

/*
------------------
Build Instructions
------------------
From the project root folder (/docker_engine_api):
export GOPATH=`pwd`
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
godep go install src/example.com/project/module/docker_engine_api/docker_socket_query.go

Developed and tested with docker v1.9.1
*/

import (
	"fmt"
	dc "github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
)

func main() {
	var cli *dc.Client
	var options *types.ContainerListOptions
	var containers []types.Container
	var containerInfo types.ContainerJSON
	//var ip string
	var err error

	// get the docker client handler
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err = dc.NewClient("unix:///var/run/docker.sock", "v1.21", nil,
		defaultHeaders)
	if err != nil {
		panic(err)
	}
	options = new(types.ContainerListOptions)
	options.All = true // no filtering criteria
	containers, err = cli.ContainerList(context.Background(), *options)
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		// get the exposed ports
		fmt.Println("From type Container")
		for _, port := range container.Ports {
			fmt.Printf("Port: %+v\n", port)
			fmt.Println()
		}
		fmt.Println("From type ContainerJSON")
		containerInfo, err = cli.ContainerInspect(context.Background(),
			container.ID)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		nwSettingsBasePorts := containerInfo.NetworkSettings.Ports
		for k, v := range nwSettingsBasePorts {
			fmt.Printf("%s => %+v\n", k, v)
		}
	} // end for each container
}
