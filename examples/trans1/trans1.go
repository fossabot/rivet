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

package main

import (
	"github.com/ennoo/rivet/examples/model"
	"github.com/ennoo/rivet/rivet"
	"github.com/ennoo/rivet/trans/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	rivet.Initialize(true, false, false)
	rivet.ListenAndServe(&rivet.ListenServe{
		Engine:      rivet.SetupRouter(testRouter1),
		DefaultPort: "8081",
	})
}

func testRouter1(engine *gin.Engine) {
	// 仓库相关路由设置
	vRepo := engine.Group("/rivet")
	vRepo.GET("/get", get1)
	vRepo.POST("/post", post1)
	vRepo.POST("/shunt", shunt1)
}

func get1(context *gin.Context) {
	rivet.Request().Call(context, http.MethodGet, "http://localhost:8082", "rivet/get")
}

func post1(context *gin.Context) {
	rivet.Request().Callback(context, http.MethodPost, "http://localhost:8082", "rivet/post", func() *response.Result {
		return &response.Result{ResultCode: response.Success, Msg: "降级处理"}
	})
}

func shunt1(context *gin.Context) {
	rivet.Response().Do(context, func(result *response.Result) {
		var test = new(model.Test)
		if err := context.ShouldBindJSON(test); err != nil {
			result.SayFail(context, err.Error())
		}
		test.Name = "trans1"
		result.SaySuccess(context, test)
	})
}
