package glasiaous

import (
	"github.com/seerwo/glasiaous/cache"
	"github.com/seerwo/glasiaous/erp"
	"github.com/seerwo/glasiaous/erp/config"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// Gasiaous struct
type Gasiaous struct {
	cache cache.Cache
}

// NewGasiaous init
func NewYangmatou() *Gasiaous {
	return &Gasiaous{}
}

//SetCache  set cache
func (c *Gasiaous) SetCache(cahce cache.Cache) {
	c.cache = cahce
}

//GetOfficialAccount get glasiaous
func (c *Gasiaous) GetGasiaous(cfg *config.Config) *erp.ERP {
	if cfg.Cache == nil {
		cfg.Cache = c.cache
	}
	return erp.NewERP(cfg)
}