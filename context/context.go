package context

import (
	"github.com/qinyer/sr-client-go/config"
	"github.com/sirupsen/logrus"
)

// Context
type Context struct {
	*config.YouShuConfig
	*logrus.Logger
}

func (c *Context) SetLogger(logger *logrus.Logger) {
	c.Logger = logger
}
