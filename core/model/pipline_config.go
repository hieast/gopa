/*
Copyright 2016 Medcl (m AT medcl.net)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package model

import (
	"encoding/json"
	"github.com/infinitbyte/gopa/core/errors"
	"github.com/infinitbyte/gopa/core/persist"
	"github.com/infinitbyte/gopa/core/util"
	"time"
)

// JointConfig configs for each joint
type JointConfig struct {
	JointName  string                 `json:"joint" config:"joint"`           //the joint name
	Parameters map[string]interface{} `json:"parameters" config:"parameters"` //kv parameters for this joint
	Enabled    bool                   `json:"enabled" config:"enabled"`
}

// PipelineConfig config for each pipeline, a pipeline may have more than one joints
type PipelineConfig struct {
	ID            string         `json:"id,omitempty" index:"id"`
	Phrase        Phrase         `json:"phrase,omitempty" config:"phrase"`
	Name          string         `json:"name,omitempty" config:"name"`
	StartJoint    *JointConfig   `json:"start,omitempty" config:"start"`
	ProcessJoints []*JointConfig `json:"process,omitempty" config:"process"`
	EndJoint      *JointConfig   `json:"end,omitempty" config:"end"`
	Created       *time.Time     `json:"created,omitempty"`
	Updated       *time.Time     `json:"updated,omitempty"`
	Tags          []string       `json:"tags,omitempty" config:"tags"`
}

const PipelineConfigBucket = "PipelineConfig"

func GetPipelineConfig(id string) (*PipelineConfig, error) {
	if id == "" {
		return nil, errors.New("empty id")
	}
	b, err := persist.GetValue(PipelineConfigBucket, []byte(id))
	if err != nil {
		panic(err)
	}
	if len(b) > 0 {
		v := PipelineConfig{}
		err = json.Unmarshal(b, &v)
		return &v, err
	}
	return nil, errors.Errorf("not found, %s", id)
}

func CreatePipelineConfig(cfg *PipelineConfig) error {
	time := time.Now().UTC()
	cfg.ID = util.GetUUID()
	cfg.Created = &time
	cfg.Updated = &time
	b, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	return persist.AddValue(PipelineConfigBucket, []byte(cfg.ID), b)
}

func UpdatePipelineConfig(id string, cfg *PipelineConfig) error {
	time := time.Now().UTC()
	cfg.ID = id
	cfg.Updated = &time
	b, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	return persist.AddValue(PipelineConfigBucket, []byte(cfg.ID), b)
}

func DeletePipelineConfig(id string) error {
	return persist.DeleteKey(PipelineConfigBucket, []byte(id))
}
