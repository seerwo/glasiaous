package context

import (
	"github.com/seerwo/glasiaous/credential"
	"github.com/seerwo/glasiaous/erp/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}

