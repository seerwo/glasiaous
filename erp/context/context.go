package context

import (
	"github.com/seerwo/gasiaous/credential"
	"github.com/seerwo/gasiaous/erp/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}

