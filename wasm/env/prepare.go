package env

import (
	"context"
	"fmt"

	"github.com/tetratelabs/wazero"
)

func PrepareRuntime(run wazero.Runtime, modname string, ctx context.Context) error {
	_, err := run.NewHostModuleBuilder("env").NewFunctionBuilder().WithFunc(func(v uint32) {
		fmt.Println("log_i32 >> ", v)
	}).Export("log_i32").Instantiate(ctx)

	return err
}
