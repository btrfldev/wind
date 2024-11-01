package component

import (
	"bytes"
	"context"
	//"context"

	"github.com/tetratelabs/wazero"
	//"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

type Component struct {
	Runtime wazero.Runtime
	CompiledComponent *wazero.CompiledModule
}

func (c *Component)Invoke(/*compName string, wasmPath string,*/ env_vars map[string]string) (string, error) {
	/*ctx := context.Background()

	run := wazero.NewRuntime(ctx)
	defer run.Close(ctx)
	wasi_snapshot_preview1.MustInstantiate(ctx, run)*/

	/*wasmFile, err := os.ReadFile(wasmPath)
	if err != nil {
		return "", err
	}*/

	/* comp, err := cs.memory.Get(compName)
	if err != nil {
		return "", err
	} */

	var stdLogBuf bytes.Buffer
	config := wazero.NewModuleConfig().WithStdout(&stdLogBuf).WithStderr(&stdLogBuf)

	for k, v := range env_vars {
		config = config.WithEnv(k, v)
	}

	_, err := c.Runtime.InstantiateModule(context.Background(), *c.CompiledComponent, config)
	if err != nil {
		return "", err
	}
	
	//_, err := run.InstantiateModule(ctx, *c.CompiledComponent /**comp.CompiledComponent*/, config)


	return stdLogBuf.String(), nil
}
