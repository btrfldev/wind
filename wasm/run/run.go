package run

import (
	"bytes"
	"context"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"

	"github.com/btrfldev/wind/env"
)

func Invoke(modname string, wasmPath string, env map[string]string) (string, error) {
	ctx := context.Background()

	run := wazero.NewRuntime(ctx)
	defer run.Close(ctx)
	wasi_snapshot_preview1.MustInstantiate(ctx, run)

	err := PrepareRuntime(run, modname, ctx)
	if err != nil {
		return "", err
	}

	wasmFile, err := os.ReadFile(wasmPath)
	if err != nil {
		return "", err
	}

	var stdoutBuf bytes.Buffer
	config := wazero.NewModuleConfig().WithStdout(&stdoutBuf)

	for k, v := range env {
		config = config.WithEnv(k, v)
	}

	_, err = run.InstantiateWithConfig(ctx, wasmFile, config)
	if err != nil {
		return "", err
	}

	return stdoutBuf.String(), nil
}
