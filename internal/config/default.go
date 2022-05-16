package config

// Default configs
func Default() Config {
	return Config{
		Nats: Nats{
			Nat1: "",
			Nat2: "",
		},
		Stream: Stream{
			StreamName:  "snapp",
			Subject:     "snapp*",
			SubjectName: "snapp@cab",
		},
		Tests: 10,
	}
}
