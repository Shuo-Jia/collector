package aggregate

import (
	"time"

	"github.com/XiaoMi/pegasus-go-client/idl/base"
	"github.com/pegasus-kv/collector/client"
)

// PartitionStats is a set of metrics retrieved from this partition.
type PartitionStats struct {
	Gpid base.Gpid

	// perfCounter's name -> the value.
	Stats map[string]float64
}

func (s *PartitionStats) update(pc *partitionPerfCounter) {
	s.Stats[pc.name] = pc.value
}

// TableStats has the aggregated metrics for this table.
type TableStats struct {
	TableName string
	AppID     int

	Partitions map[int]*PartitionStats

	// the time when the stats was generated
	Timestamp time.Time

	// The aggregated value of table metrics.
	// perfCounter's name -> the value.
	Stats map[string]float64
}

// ClusterStats is the aggregated metrics for all the tables in this cluster.
type ClusterStats struct {
	Timestamp time.Time

	Stats map[string]float64
}

func newTableStats(info *client.TableInfo) *TableStats {
	tb := &TableStats{
		TableName:  info.TableName,
		AppID:      info.AppID,
		Partitions: make(map[int]*PartitionStats),
		Stats:      make(map[string]float64),
		Timestamp:  time.Now(),
	}
	for i := 0; i < info.PartitionCount; i++ {
		tb.Partitions[i] = &PartitionStats{
			Gpid:  base.Gpid{Appid: int32(info.AppID), PartitionIndex: int32(i)},
			Stats: make(map[string]float64),
		}
	}
	return tb
}

func (tb *TableStats) aggregate() {
	tb.Timestamp = time.Now()
	for _, part := range tb.Partitions {
		for name, value := range part.Stats {
			tb.Stats[name] += value
		}
	}
	extendStats(&tb.Stats)
}

func aggregateCustomStats(elements []string, stats *map[string]float64, resultName string) {
	aggregated := float64(0)
	for _, ele := range elements {
		if v, found := (*stats)[ele]; found {
			aggregated += v
		}
	}
	(*stats)[resultName] = aggregated
}

func extendStats(stats *map[string]float64) {
	var reads = []string{
		"get",
		"multi_get",
		"scan",
	}
	var readQPS []string
	for _, r := range reads {
		readQPS = append(readQPS, r+"_qps")
	}
	var readBytes []string
	for _, r := range reads {
		readBytes = append(readBytes, r+"_bytes")
	}
	aggregateCustomStats(readQPS, stats, "read_qps")
	aggregateCustomStats(readBytes, stats, "read_bytes")

	var writes = []string{
		"put",
		"remove",
		"multi_put",
		"multi_remove",
		"check_and_set",
		"check_and_mutate",
	}
	var writeQPS []string
	for _, w := range writes {
		writeQPS = append(writeQPS, w+"_qps")
	}
	var writeBytes []string
	for _, w := range writes {
		writeBytes = append(writeBytes, w+"_bytes")
	}
	aggregateCustomStats(writeQPS, stats, "write_qps")
	aggregateCustomStats(writeBytes, stats, "write_bytes")
}
