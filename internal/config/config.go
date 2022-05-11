package config

type Config struct {
	Nat1        string `koanf:"nat1"`
	Nat2        string `koanf:"nat2"`
	StreamName  string `koanf:"stream_name"`
	Subject     string `koanf:"subject"`
	SubjectName string `koanf:"subject_name"`
}

func Load() Config {
	return Default()
}
