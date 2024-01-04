package config

import "github.com/amirhnajafiz/j-mirror/internal/model"

// Default configs
func Default() Config {
	return Config{
		Clusters: []string{},
		Stream: model.Stream{
			Name:        "snapp",
			Subject:     "snapp*",
			SubjectName: "snapp@cab",
		},
	}
}
