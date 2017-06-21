package tsdbadapter

import (
	"fmt"
	"github.com/Symantec/scotty/tsdb"
	"strings"
	"time"
)

func (a *Adapter) fetch(
	region,
	accountNumber,
	instanceId string,
	name string,
	start,
	end int64) (tsdb.TimeSeries, error) {
	fsMetric := strings.HasPrefix(name, "fs:")
	entries, err := a.reader.Read(
		computeAssetId(region, accountNumber, instanceId, fsMetric),
		millisToTime(start),
		millisToTime(end))
	if err != nil {
		return nil, err
	}
	var result tsdb.TimeSeries
	for _, entry := range entries {
		val, ok := entry.Values[name]
		if ok {
			result = append(
				result,
				tsdb.TsValue{
					Ts:    float64(entry.Time.Unix()),
					Value: val,
				})
		}
	}
	return result, nil
}

func millisToTime(millis int64) time.Time {
	mils := millis % 1000
	secs := millis / 1000
	return time.Unix(secs, mils*1000*1000)
}

func computeAssetId(
	region, accountNumber, instanceId string, fsMetric bool) string {
	if fsMetric {
		return fmt.Sprintf(
			"arn:aws:ec2:%s:%s:instance/%s:fs//",
			region,
			accountNumber,
			instanceId)
	}
	return fmt.Sprintf(
		"arn:aws:ec2:%s:%s:instance/%s",
		region,
		accountNumber,
		instanceId)
}
