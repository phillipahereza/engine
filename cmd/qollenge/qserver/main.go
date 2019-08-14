package main

import (
	"elbix.dev/engine/cmd/cerulean"
	"elbix.dev/engine/pkg/cli"
	"elbix.dev/engine/pkg/grpcgw"
	"elbix.dev/engine/pkg/initializer"
	"elbix.dev/engine/pkg/log"
)

func main() {
	ctx := cli.Context()
	if err := cerulean.InitializeConfig(ctx, true); err != nil {
		log.Fatal("Dependency injection failed", log.Err(err))
	}

	defer initializer.Initialize(ctx)()

	grpcgw.Serve(ctx)

}
