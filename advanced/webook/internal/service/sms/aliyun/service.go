// Copyright 2023 igevin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aliyun

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"strconv"
	"strings"
)

type Service struct {
	client   *dysmsapi.Client
	signName string
}

func NewService(c *dysmsapi.Client, signName string) *Service {
	return &Service{
		client:   c,
		signName: signName,
	}
}

func (s *Service) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	req := dysmsapi.CreateSendSmsRequest()
	req.Scheme = "https"
	// 阿里云多个手机号为字符串逗号间隔
	req.PhoneNumbers = strings.Join(numbers, ",")
	req.SignName = s.signName
	// 传的是 JSON
	argsMap := make(map[string]string, len(args))
	for idx, arg := range args {
		argsMap[strconv.Itoa(idx)] = arg
	}
	// 这意味着，你的模板必须是 你的短信验证码是{0}
	// 你的短信验证码是{code}
	bCode, err := json.Marshal(argsMap)
	if err != nil {
		return err
	}
	req.TemplateParam = string(bCode)
	req.TemplateCode = tplId

	var resp *dysmsapi.SendSmsResponse
	resp, err = s.client.SendSms(req)
	if err != nil {
		return err
	}

	if resp.Code != "OK" {
		return fmt.Errorf("发送失败，code: %s, 原因：%s",
			resp.Code, resp.Message)
	}
	return nil
}
