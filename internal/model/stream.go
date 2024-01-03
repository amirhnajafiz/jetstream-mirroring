package model

// Stream model is for NATS streams
type Stream struct {
	Name        string `koanf:"name"`
	Subject     string `koanf:"subject"`
	SubjectName string `koanf:"subject_name"`
}
