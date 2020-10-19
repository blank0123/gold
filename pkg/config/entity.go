package config

import "time"

type server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type redis struct {
	Admin struct {
		Addr     string
		Password string
		DB       int
	}
}

type mysql struct {
	Driver string
	DSN    string
}

type jwt struct {
	Secret string
}

type app struct {
	PageSize int
}
