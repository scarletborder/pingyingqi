package provider

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"pingyingqi/config"
	aiprovider "pingyingqi/models/AiProvider"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Wenxin struct {
	serviceName  string
	website      string
	api_key      string
	secret_key   string
	access_token string
}

func NewWenxin(Api_Key string, Secret_Key string) (*Wenxin, error) {
	w := Wenxin{serviceName: "文心一言", website: "https://yiyan.baidu.com/", api_key: config.EnvCfg.WenxinApiKey, secret_key: config.EnvCfg.WenxinSecretKey}
	// 过期自动更新
	var err error
	var exp int
	w.access_token, exp, err = w.getAccessToken()
	if err != nil {
		return &w, err
	}
	// go Ticker
	expire_timer := time.NewTimer(time.Duration(exp) * time.Second)
	defer expire_timer.Stop()

	go func() {
		for {
			<-expire_timer.C
			w.access_token, exp, err = w.getAccessToken()
			if err != nil {
				return
			}
			expire_timer.Reset(time.Duration(exp) * time.Second)
		}
	}()
	return &w, err
}

func (w *Wenxin) GetServiceName() string {
	return w.serviceName
}

// 作为wenxinJson中Messages(切片)成员的元素类型
type wenxinMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type wenxinBody struct {
	Messages []wenxinMessage `json:"messages"`
	Top_p    float32         `json:"top_p"`
}

type wenxinResp struct {
	Result string `json:"result"`
	ErrMsg string `json:"error_msg"`
}

func (w *Wenxin) Prompt(pro string) (string, error) {
	url := "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions?access_token=" + w.access_token
	payloadData := wenxinBody{Messages: []wenxinMessage{{Role: "user", Content: pro}}, Top_p: 0.3}
	payloadJson, _ := json.Marshal(&payloadData)
	payload := strings.NewReader(string(payloadJson))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return ``, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return ``, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ``, err
	}
	var wenxinresp wenxinResp
	err = json.Unmarshal(body, &wenxinresp)
	if err != nil {
		return ``, err
	}
	if wenxinresp.ErrMsg != `` {
		return ``, errors.New(wenxinresp.ErrMsg)
	}
	return wenxinresp.Result, nil
}

type authResp struct {
	Expires_in int    `json:"expires_in"`
	AT         string `json:"access_token"`
	Errdesp    string `json:"error_description"`
}

/** 登录
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * @return string 鉴权签名信息（Access Token）
 * @return error
 */
func (w *Wenxin) getAccessToken() (string, int, error) {
	var err error
	logrus.Debugln("access", w.api_key, w.secret_key)
	url := "https://aip.baidubce.com/oauth/2.0/token?client_id=" + w.api_key + "&client_secret=" + w.secret_key + "&grant_type=client_credentials"
	payload := strings.NewReader(``)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return ``, 0, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	var res *http.Response
	for tryi := -1; tryi < config.EnvCfg.MaxRetryTime; tryi++ {
		res, err = client.Do(req)
		if err == nil && res.StatusCode == 200 {
			break
		}
	}
	if err != nil {
		return ``, 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ``, 0, err
	}
	var ar authResp
	err = json.Unmarshal(body, &ar)

	if err != nil {
		return ``, 0, err
	} else if ar.Errdesp != `` {
		return ``, 0, errors.New("Request succeed, but " + ar.Errdesp)
	}
	return ar.AT, ar.Expires_in, nil
}

func init() {
	if config.EnvCfg.WenxinApiKey != `` && config.EnvCfg.WenxinSecretKey != `` && config.EnvCfg.DefaultProvider != `nil` {
		Wenxin, err := NewWenxin(config.EnvCfg.WenxinApiKey, config.EnvCfg.WenxinSecretKey)
		if err != nil {
			aiprovider.FailOnAiProviderCreate("wenxin", err)
		} else {
			aiprovider.AiHelper.Provider = append(aiprovider.AiHelper.Provider, Wenxin)
		}
	}
}
