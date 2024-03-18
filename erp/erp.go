package erp

import (
	"github.com/seerwo/glasiaous/credential"
	"github.com/seerwo/glasiaous/erp/config"
	"github.com/seerwo/glasiaous/erp/context"
	"github.com/seerwo/glasiaous/erp/oauth"
	"github.com/seerwo/glasiaous/erp/order"
)

//Erp yangyue relate api
type ERP struct {
	ctx *context.Context
}

//NewOpenPlatform new openplatform
func NewERP(cfg *config.Config) *ERP {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, cfg.AuthCode, credential.CacheKeyGaErpPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &ERP{ctx: ctx}
}

//SetAccessTokenHandle custom access_token get method
func (c *ERP) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	c.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (c *ERP) GetContext() *context.Context {
	return c.ctx
}

//GetAccessToken get access_token
func (c *ERP) GetAccessToken() (string, error) {
	return c.ctx.GetAccessToken()
}

// GetOauth oauth2 web oauth
func (c *ERP) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(c.ctx)
}

// GetOrder get order
func (c *ERP) GetOrder() *order.Order {
	return order.NewOrder(c.ctx)
}
