package component

import (
	"context"
	"fmt"
	"strings"

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
	wasi_snapshot_preview1.MustInstantiate(ctx, run)

	comp, err := run.NewHostModuleBuilder("env").NewFunctionBuilder().WithFunc(func(v uint32) {
		fmt.Println("log_i32 >> ", v)
	}).Export("log_i32").Compile(ctx)

	if err != nil {
		return err
	}

	comp, err = run.CompileModule(ctx, wasmFile)
	if err != nil {
		return err
	}

	err = cs.memory.Put(compName, Component{
		Runtime:           run,
		CompiledComponent: &comp,
	})

	return err
}

func (cs *ComponentStorage) Get(compName string) (Component, error) {
	comp, err := cs.memory.Get(compName)
	if err != nil {
		return Component{}, err
	} else {
		return comp, nil
	}
}

func (cs *ComponentStorage) Has(compName string) (bool) {
	return cs.memory.Has(compName)
}

func (cs *ComponentStorage) List(prefix string) ([]string) {
	list, _ := cs.memory.List(func(k string) bool {
		return strings.HasPrefix(k, prefix)
	})
	return list
}