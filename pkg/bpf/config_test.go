/*
 * Copyright 2023 The Kmesh Authors.
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

 * Author: Min Chen
 * Create: 2024-01-23
 */

package bpf

import "testing"

func TestConfig_ParseConfig(t *testing.T) {
	type fields struct {
		BpfFsPath        string
		Cgroup2Path      string
		EnableKmesh      bool
		EnableMda        bool
		BpfVerifyLogSize int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "valid Mda config",
			fields: fields{
				EnableMda: true,
			},
			wantErr: false,
		},
		// {
		// 	name: "valid Kmesh config",
		// 	fields: fields{
		// 		EnableKmesh: true,

		// 	},
		// },
		{
			name: "none of EnableKmesh or EnableMda",
			fields: fields{
				EnableKmesh: false,
				EnableMda:   false,
			},
			wantErr: true,
		},
		{
			name: "both EnableKmesh and EnableMda",
			fields: fields{
				EnableKmesh: true,
				EnableMda:   true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BpfFsPath:        tt.fields.BpfFsPath,
				Cgroup2Path:      tt.fields.Cgroup2Path,
				EnableKmesh:      tt.fields.EnableKmesh,
				EnableMda:        tt.fields.EnableMda,
				BpfVerifyLogSize: tt.fields.BpfVerifyLogSize,
			}
			if err := c.ParseConfig(); (err != nil) != tt.wantErr {
				t.Errorf("ParseConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
