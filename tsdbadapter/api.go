package tsdbadapter

import (
	"github.com/Symantec/chproxy/chreader"
	"github.com/Symantec/scotty/tsdb"
)

type Adapter struct {
	reader chreader.Reader
}

func New(r chreader.Reader) *Adapter {
	return &Adapter{reader: r}
}

func (a *Adapter) Fetch(
	region,
	accountNumber,
	instanceId string,
	name string,
	start,
	end int64) (tsdb.TimeSeries, error) {
	return a.fetch(
		region,
		accountNumber,
		instanceId,
		name,
		start,
		end)
}
