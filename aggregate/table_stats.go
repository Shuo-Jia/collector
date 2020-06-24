package aggregate

var allMetricNames []string = []string{
	"get_qps",
	"multi_get_qps",
	"put_qps",
	"multi_put_qps",
	"remove_qps",
	"multi_remove_qps",
	"incr_qps",
	"check_and_set_qps",
	"check_and_mutate_qps",
	"scan_qps",
	"duplicate_qps",
	"dup_shipped_ops",
	"dup_failed_shipping_ops",
	"recent_read_cu",
	"recent_write_cu",
	"recent_expire_count",
	"recent_filter_count",
	"recent_abnormal_count",
	"recent_write_throttling_delay_count",
	"recent_write_throttling_reject_count",
	"storage_mb",
	"storage_count",
	"rdb_block_cache_hit_rate",
	"rdb_block_cache_mem_usage",
	"rdb_index_and_filter_blocks_mem_usage",
	"rdb_memtable_mem_usage",
	"rdb_estimate_num_keys",
	"rdb_bf_seek_negatives_rate",
	"rdb_bf_point_negatives_rate",
	"rdb_bf_point_false_positive_rate",
	"read_qps",
	"write_qps",
	"backup_request_qps",
	"get_bytes",
	"multi_get_bytes",
	"scan_bytes",
	"put_bytes",
	"multi_put_bytes",
	"check_and_set_bytes",
	"check_and_mutate_bytes",
	"read_bytes",
	"write_bytes",
}
