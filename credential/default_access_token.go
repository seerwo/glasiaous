package credential

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/glasiaous/cache"
	"github.com/seerwo/glasiaous/util"
	"sync"

)

const (
	//AccessTokenURL 获取access_token的接口
	accessTokenURL = "https://onlinedemo.glasiaous.com/SNM_DEMO/api/security/login2"
	//CacheKeyOfficialAccountPrefix 微信公众号cache key前缀
	CacheKeyGaErpPrefix = "ga_erp_"
	//CacheKeyMiniProgramPrefix 小程序cache key前缀
	CacheKeyMiniProgramPrefix = "gowechat_miniprogram_"
)

//DefaultAccessToken 默认AccessToken 获取
type DefaultAccessToken struct {
	appID           string
	appSecret       string
	authCode 		string
	cacheKeyPrefix  string
	cache           cache.Cache
	accessTokenLock *sync.Mutex
}

//NewDefaultAccessToken new DefaultAccessToken
func NewDefaultAccessToken(appID, appSecret, authCode, cacheKeyPrefix string, cache cache.Cache) AccessTokenHandle {
	if cache == nil {
		panic("cache is ineed")
	}
	return &DefaultAccessToken{
		appID:           appID,
		appSecret:       appSecret,
		authCode: 		 authCode,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		accessTokenLock: new(sync.Mutex),
	}
}

//ResAccessToken struct
type ResAccessToken struct {
	util.CommonError
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	SecurityCode string `json:"securityCode"`

	//Data UserInfo `json:"data"`
	//AccessToken string `json:"access_token"`
	//ExpiresIn   int64  `json:"expires_in"`
}

func  (ak *DefaultAccessToken)GetAccessParam(method string, req interface{})(accessParam string, err error){
	m := map[string]string{}

	m["company"] = ak.appID
	m["userID"] = ak.appID
	m["Password"] = ak.appSecret
	m["lang"] = "ja-JP"
	m["userkey"] = "testAPI"

	j, _ := json.Marshal(req)
	m["data"] = string(j)

	//if m["sign"],err = util.ParamSign(m, ak.appSecret); err != nil {
	//	return
	//}
	js, _ := json.Marshal(m)
	accessParam = string(js)
	return accessParam, nil
}

//GetAccessToken 获取access_token,先从cache中获取，没有则从服务端获取
func (ak *DefaultAccessToken) GetAccessToken() (accessToken string, err error) {
	//加上lock，是为了防止在并发获取token时，cache刚好失效，导致从微信服务器上获取到不同token
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()

	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", ak.cacheKeyPrefix, ak.appID)
	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}

	//cache失效，从微信服务器获取
	//var resAccessToken ResAccessToken
	//resAccessToken, err = GetTokenFromServer(ak.appID, ak.appSecret)
	//if err != nil {
	//	return
	//}

	//expires := resAccessToken.ExpiresIn - 1500
	//err = ak.cache.Set(accessTokenCacheKey, resAccessToken.AccessToken, time.Duration(expires)*time.Second)
	//if err != nil {
	//	return
	//}
	//accessToken = resAccessToken.AccessToken
	return
}

//GetTokenFromServer 强制从微信服务器获取token
func GetTokenFromServer(appID, appSecret string) (resAccessToken ResAccessToken, err error) {
	url := fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", accessTokenURL, appID, appSecret)
	var body []byte
	body, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resAccessToken)
	if err != nil {
		return
	}
	if resAccessToken.Status != "" {
		err = fmt.Errorf("get access_token error : errcode=%v , errormsg=%v", resAccessToken.Status, resAccessToken.Messages)
		return
	}
	return
}

