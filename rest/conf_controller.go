// Copyright 2019 GitBitEx.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/pavelkrolevets/gitbitex-spot/service"
	"net/http"
)

// GET /configs
func GetConfigs(ctx *gin.Context) {
	configs, err := service.GetConfigs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, newMessageVo(err))
		return
	}

	m := map[string]string{}
	for _, config := range configs {
		m[config.Key] = config.Value
	}

	ctx.JSON(http.StatusOK, m)
}
