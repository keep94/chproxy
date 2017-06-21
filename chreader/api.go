package chreader

import (
	"github.com/Symantec/scotty/lib/yamlutil"
	"time"
)

// Instance related metric names
const (
	CpuUsedPercentAvg = "cpu:used:percent.avg"
	CpuUsedPercentMax = "cpu:used:percent.max"
	CpuUsedPercentMin = "cpu:used:percent.min"
	// Continue
)

// File system related metric names
const (
	FsSizeBytesAvg = "fs:size:bytes.avg"
	FsSizeBytesMax = "fs:size:bytes.max"
	FsSizeBytesMin = "fs:size:bytes.min"
	// Continue
)

type Entry struct {
	Time   time.Time
	Values map[string]float64
}

type CH interface {
	Fetch(url string) (entries []*Entry, next string, err error)
}

var (
	DefaultCH CH = &chType{}
)

type Config struct {
	ApiKey string `yaml:"apiKey"`
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type configFields Config
	return yamlutil.StrictUnmarshalYAML(unmarshal, (*configFields)(c))
}

func (c *Config) Reset() {
	*c = Config{}
}

type Reader interface {
	Read(assetId string, start, end time.Time) ([]*Entry, error)
}

func NewReader(c Config, ch CH, now func() time.Time) Reader {
	return &chReaderType{
		config: c,
		ch:     ch,
		now:    now}
}
