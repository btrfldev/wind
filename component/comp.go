package component

import (
	"context"
	"fmt"

	"github.com/tetratelabs/wazero"
)

type Component struct {
	Name              string
	CompiledComponent wazero.CompiledModule
}

type ComponentStorage struct {
	MemoryStore stor
}

func (cs *ComponentStorage)NewComponentStorage() {
	
}

func (cs *Components) Register(run wazero.Runtime, compName string, wasmFile []byte, ctx context.Context) (err error) {
	_, err = run.NewHostModuleBuilder("env").NewFunctionBuilder().WithFunc(func(v uint32) {
		fmt.Println("log_i32 >> ", v)
	}).Export("log_i32").Instantiate(ctx)

	if err != nil {
		return err
	}

	comp, err := run.CompileModule(ctx, wasmFile)

	cs[compName] = Component{
		Name:              compName,
		CompiledComponent: comp,
	}

	return err
}
