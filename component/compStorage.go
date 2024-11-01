package component

import (
	"context"
	"fmt"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"

	"github.com/btrfldev/wind/storage"
)



type ComponentStorage struct {
	memory storage.MemoryStore[string, Component]
}

func NewComponentStorage() *ComponentStorage {
	return &ComponentStorage{
		memory: *storage.NewMemoryStore[string, Component](),
	}
}

func (cs *ComponentStorage) Register(ctx context.Context, compName string, wasmFile []byte) (err error) {
	run := wazero.NewRuntime(ctx)
	//defer run.Close(ctx)
	wasi_snapshot_preview1.MustInstantiate(ctx, run)

	comp, err := run.NewHostModuleBuilder("env").NewFunctionBuilder().WithFunc(func(v uint32) {
		fmt.Println("log_i32 >> ", v)
	}).Export("log_i32").Compile(ctx)

	if err != nil {
		return err
	}

	comp, err = run.CompileModule(ctx, wasmFile)
	if err!=nil{
		return nil
	}

	err = cs.memory.Put(compName, Component{
		Runtime: run,
		CompiledComponent: &comp,
	})

	return err
}

func (cs *ComponentStorage) Get(compName string) (Component, error) {
	comp, err := cs.memory.Get(compName)
	if err!=nil{
		return Component{}, err
	} else {
		return comp, nil
	}
}