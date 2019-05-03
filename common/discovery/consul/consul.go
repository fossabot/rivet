/*
 * Copyright (c) 2019. ENNOO - All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package consul

import (
	"fmt"
	"github.com/ennoo/rivet/common/util/env"
	"github.com/ennoo/rivet/common/util/file"
	"github.com/ennoo/rivet/common/util/string"
	"github.com/ennoo/rivet/dolphin/http/request"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.uber.org/zap"
	"os"
	"strings"
	"viewhigh.com/dams/common/model"
)

var ServiceID = xid.New().String()

func Health(r *gin.Engine) {
	// 仓库相关路由设置
	vRepo := r.Group("/health")
	vRepo.GET("/check", health)
}

// 调用此方法注册 consul
//
// consulUrl：consul 注册地址，包括端口号（优先通过环境变量 CONSUL_URL 获取）
//
// serviceName：注册到 consul 的服务名称（优先通过环境变量 SERVICE_NAME 获取）
func ConsulRegister(consulUrl string, serviceName string) {
	defer func() {
		zap.S().Info("register consul start")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
			os.Exit(0)
		}
	}()
	consulRegister(consulUrl, serviceName)
}

func consulRegister(consulUrl string, serviceName string) {
	hosts, err := file.ReadFileByLine("/etc/hostname")
	if nil != err {
		panic(err)
	}
	zap.S().Info("serviceID = ", ServiceID)
	containerID := str.Trim(hosts[0])
	zap.S().Info("containerID = ", containerID)
	restJsonHandler := request.RestJsonHandler{
		Header:  nil,
		Cookies: nil,
		RestHandler: request.RestHandler{
			RemoteServer: strings.Join([]string{
				"http://",
				env.GetEnvDafult(env.ConsulUrl, consulUrl),
				"/v1/agent/service/register"}, ""),
			Uri: "",
			Param: model.ConsulRegister{
				ID:                ServiceID,
				Name:              env.GetEnvDafult(env.ServiceName, serviceName),
				Address:           containerID,
				Port:              80,
				EnableTagOverride: false,
				Check: model.ConsulCheck{
					DeregisterCriticalServiceAfter: "1m",
					HTTP:                           strings.Join([]string{"http://", containerID, "/health/check"}, ""),
					Interval:                       "10s"}},
			Values: nil}}
	body, err := restJsonHandler.Put()
	if nil != err {
		zap.S().Error(err.Error())
	}
	zap.S().Info("register result = ", string(body))
}