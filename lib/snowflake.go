package lib

import (
	"github.com/bwmarrin/snowflake"
	"os"
)

type Snowflake struct {
	env    Env
	logger Logger
}

func NewSnowflakeService(env Env, logger Logger) Snowflake {
	return Snowflake{
		env:    env,
		logger: logger,
	}
}

func (s Snowflake) GenerateID() (int64, error) {
	n, err := snowflake.NewNode(s.env.NodeNumber)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	return n.Generate().Int64(), nil
}
