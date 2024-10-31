package component

import (
	"bytes"
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func (c *Component)Invoke(/*compName string, wasmPath string,*/ env_vars map[string]string) (string, error) {
	ctx := context.Background()

	run := wazero.NewRuntime(ctx)
	defer run.Close(ctx)
	wasi_snapshot_preview1.MustInstantiate(ctx, run)

	/*wasmFile, err := os.ReadFile(wasmPath)
	if err != nil {
		return "", err
	}*/

	/* comp, err := cs.memory.Get(compName)
	if err != nil {
		return "", err
	} */

	var stdoutBuf bytes.Buffer
	config := wazero.NewModuleConfig().WithStdout(&stdoutBuf)

	for k, v := range env_vars {
		config = config.WithEnv(k, v)
	}


	_, err := run.InstantiateModule(ctx, *c.CompiledComponent /**comp.CompiledComponent*/, config)
	if err != nil {
		return "", err
	}

	return stdoutBuf.String(), nil
}
