package component

import (
	"bytes"
	"context"
	"fmt"

	"github.com/tetratelabs/wazero"
)

type Component struct {
	Runtime           wazero.Runtime
	CompiledComponent *wazero.CompiledModule
}

func (c *Component) Invoke(env_vars map[string]string) (string, error) {
	ctx := context.Background()
	var stdLogBuf bytes.Buffer
	config := wazero.NewModuleConfig().WithStdout(&stdLogBuf).WithStderr(&stdLogBuf)

	for k, v := range env_vars {
		config = config.WithEnv(k, v)
	}

	_, err := c.Runtime.InstantiateModule(ctx, *c.CompiledComponent, config)
	if err != nil {
		return "", err
	}


	log := stdLogBuf.String()
	if log != "" {
		fmt.Println(log)
	}

	return log, nil
}
