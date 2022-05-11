package config

func Default() Config {
	return Config{
		Nat1:        "",
		Nat2:        "",
		StreamName:  "snapp",
		Subject:     "snapp*",
		SubjectName: "snapp.test",
	}
}
