package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pingyingqi/config"
	"strings"
	"time"

	"google.golang.org/protobuf/internal/errors"
)

type Wenxin struct {
	serviceName  string
	website      string
	api_key      string
	secret_key   string
	access_token string
}

func NewWenxin(Api_Key string, Secret_Key string) (*Wenxin, error) {
	w := Wenxin{serviceName: "文心一言", website: "https://yiyan.baidu.com/"}
	// 过期自动更新
	var err error
	var exp int
	w.access_token, err, exp = w.getAccessToken()
	if err != nil {
		return &w, err
	}
	// go Ticker
	expire_timer := time.NewTimer(time.Duration(exp) * time.Second)
	defer expire_timer.Stop()

	go func() {
		for {
			<-expire_timer.C
			w.access_token, err, exp = w.getAccessToken()
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
	Messages string `json:"messages"`
	Top_p    int    `json:"top_p"`
}

func (w *Wenxin) Prompt(pro string) (string, error) {
	url := "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions?access_token=" + w.access_token
	myMessage := wenxinMessage{Role: "user", Content: pro}
	// 各种序列化Prompt
	// wenxinBody()
	payload := strings.NewReader(`{"messages":[{"role":"user","content":"Please help me output the result of the code below.\r\nYour output format and tips are here.\r\n- If the code run successfully, output the result in \"Data\" and the \"Code\" is 0.\r\n- It may exists some bug and you need to point it. If it exists bug, output the\r\nbug in \"Data\" and the \"Code\" is 1.\r\n- If the program runs into an infinite loop, the \"Code\" is 3, output partial result in \"Data\".\r\n\r\n` + "`" + `` + "`" + `` + "`" + `json\r\n{\"Code\":0,\"Data\":\"The result\"}\r\n` + "`" + `` + "`" + `` + "`" + `\r\nYou should not output other content except for the json text.\r\nAnd the code is here\r\n` + "`" + `` + "`" + `` + "`" + `golang\r\npackage main;\r\nimport \"fmt\"\r\nfunc main(){\r\n        for true{\r\n                fmt.Println(\"hello\")\r\n        }\r\n}\r\n` + "`" + `` + "`" + `` + "`" + `golang\npackage main;\nimport \"fmt\"\nfunc main(){\n\tfor true{\n\t\tfmt.Println(\"hello\")\n\t}\n}\n` + "`" + `` + "`" + `` + "`" + `"},{"role":"assistant","content":"` + "`" + `` + "`" + `` + "`" + `json\n{\"Code\":3,\"Data\":\"hello\"}\n` + "`" + `` + "`" + `` + "`" + `"}],"top_p":0.5}`)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return ``, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
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
func (w *Wenxin) getAccessToken() (string, error, int) {
	var err error
	url := "https://aip.baidubce.com/oauth/2.0/token?client_id=" + w.api_key + "&client_secret=" + w.secret_key + "&grant_type=client_credentials"
	payload := strings.NewReader(``)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return ``, err, 0
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
		return ``, err, 0
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ``, err, 0
	}
	var ar authResp
	err = json.Unmarshal(body, &ar)

	if err != nil {
		return ``, err, 0
	} else if ar.Errdesp != `` {
		return ``, errors.New(ar.Errdesp), 0
	}
	return ar.AT, nil, ar.Expires_in
}
