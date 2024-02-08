/*
 * Copyright 2024 The Kmesh Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.

 * Author: LemmyHuang
 * Create: 2021-10-09
 */

package controller_workload

import (
	"kmesh.net/kmesh/pkg/controller_workload/interfaces"
)

var (
	stopCh = make(chan struct{})
	client interfaces.ClientFactory
)

func Start() error {
	var err error

	client, err = config.NewClient()
	if err != nil {
		return err
	}

	return client.Run(stopCh)
}

func Stop() {
	var obj struct{}
	stopCh <- obj
	close(stopCh)
	client.Close()
}
