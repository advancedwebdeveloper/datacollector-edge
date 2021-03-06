// Copyright 2018 StreamSets Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package process

const (
	DefaultProcessMetricsCaptureInterval = -1
)

type Config struct {
	ProcessMetricsCaptureInterval int64 `toml:"process-metrics-capture-interval"`
}

// NewConfig returns a new Config with default settings.
func NewConfig() Config {
	return Config{
		ProcessMetricsCaptureInterval: DefaultProcessMetricsCaptureInterval,
	}
}
