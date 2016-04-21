package app

type MyJsonName struct {
	Total struct {
		Completion struct {
			SizeInBytes int `json:"size_in_bytes"`
		} `json:"completion"`
		Docs struct {
			Count   int `json:"count"`
			Deleted int `json:"deleted"`
		} `json:"docs"`
		Fielddata struct {
			Evictions         int `json:"evictions"`
			MemorySizeInBytes int `json:"memory_size_in_bytes"`
		} `json:"fielddata"`
		Flush struct {
			Total             int `json:"total"`
			TotalTimeInMillis int `json:"total_time_in_millis"`
		} `json:"flush"`
		Get struct {
			Current             int `json:"current"`
			ExistsTimeInMillis  int `json:"exists_time_in_millis"`
			ExistsTotal         int `json:"exists_total"`
			MissingTimeInMillis int `json:"missing_time_in_millis"`
			MissingTotal        int `json:"missing_total"`
			TimeInMillis        int `json:"time_in_millis"`
			Total               int `json:"total"`
		} `json:"get"`
		Indexing struct {
			DeleteCurrent        int  `json:"delete_current"`
			DeleteTimeInMillis   int  `json:"delete_time_in_millis"`
			DeleteTotal          int  `json:"delete_total"`
			IndexCurrent         int  `json:"index_current"`
			IndexFailed          int  `json:"index_failed"`
			IndexTimeInMillis    int  `json:"index_time_in_millis"`
			IndexTotal           int  `json:"index_total"`
			IsThrottled          bool `json:"is_throttled"`
			NoopUpdateTotal      int  `json:"noop_update_total"`
			ThrottleTimeInMillis int  `json:"throttle_time_in_millis"`
		} `json:"indexing"`
		Merges struct {
			Current                    int `json:"current"`
			CurrentDocs                int `json:"current_docs"`
			CurrentSizeInBytes         int `json:"current_size_in_bytes"`
			Total                      int `json:"total"`
			TotalAutoThrottleInBytes   int `json:"total_auto_throttle_in_bytes"`
			TotalDocs                  int `json:"total_docs"`
			TotalSizeInBytes           int `json:"total_size_in_bytes"`
			TotalStoppedTimeInMillis   int `json:"total_stopped_time_in_millis"`
			TotalThrottledTimeInMillis int `json:"total_throttled_time_in_millis"`
			TotalTimeInMillis          int `json:"total_time_in_millis"`
		} `json:"merges"`
		Percolate struct {
			Current           int    `json:"current"`
			MemorySize        string `json:"memory_size"`
			MemorySizeInBytes int    `json:"memory_size_in_bytes"`
			Queries           int    `json:"queries"`
			TimeInMillis      int    `json:"time_in_millis"`
			Total             int    `json:"total"`
		} `json:"percolate"`
		QueryCache struct {
			CacheCount        int `json:"cache_count"`
			CacheSize         int `json:"cache_size"`
			Evictions         int `json:"evictions"`
			HitCount          int `json:"hit_count"`
			MemorySizeInBytes int `json:"memory_size_in_bytes"`
			MissCount         int `json:"miss_count"`
			TotalCount        int `json:"total_count"`
		} `json:"query_cache"`
		Recovery struct {
			CurrentAsSource      int `json:"current_as_source"`
			CurrentAsTarget      int `json:"current_as_target"`
			ThrottleTimeInMillis int `json:"throttle_time_in_millis"`
		} `json:"recovery"`
		Refresh struct {
			Total             int `json:"total"`
			TotalTimeInMillis int `json:"total_time_in_millis"`
		} `json:"refresh"`
		RequestCache struct {
			Evictions         int `json:"evictions"`
			HitCount          int `json:"hit_count"`
			MemorySizeInBytes int `json:"memory_size_in_bytes"`
			MissCount         int `json:"miss_count"`
		} `json:"request_cache"`
		Search struct {
			FetchCurrent       int `json:"fetch_current"`
			FetchTimeInMillis  int `json:"fetch_time_in_millis"`
			FetchTotal         int `json:"fetch_total"`
			OpenContexts       int `json:"open_contexts"`
			QueryCurrent       int `json:"query_current"`
			QueryTimeInMillis  int `json:"query_time_in_millis"`
			QueryTotal         int `json:"query_total"`
			ScrollCurrent      int `json:"scroll_current"`
			ScrollTimeInMillis int `json:"scroll_time_in_millis"`
			ScrollTotal        int `json:"scroll_total"`
		} `json:"search"`
		Segments struct {
			Count                       int `json:"count"`
			DocValuesMemoryInBytes      int `json:"doc_values_memory_in_bytes"`
			FixedBitSetMemoryInBytes    int `json:"fixed_bit_set_memory_in_bytes"`
			IndexWriterMaxMemoryInBytes int `json:"index_writer_max_memory_in_bytes"`
			IndexWriterMemoryInBytes    int `json:"index_writer_memory_in_bytes"`
			MemoryInBytes               int `json:"memory_in_bytes"`
			NormsMemoryInBytes          int `json:"norms_memory_in_bytes"`
			StoredFieldsMemoryInBytes   int `json:"stored_fields_memory_in_bytes"`
			TermVectorsMemoryInBytes    int `json:"term_vectors_memory_in_bytes"`
			TermsMemoryInBytes          int `json:"terms_memory_in_bytes"`
			VersionMapMemoryInBytes     int `json:"version_map_memory_in_bytes"`
		} `json:"segments"`
		Store struct {
			SizeInBytes          int `json:"size_in_bytes"`
			ThrottleTimeInMillis int `json:"throttle_time_in_millis"`
		} `json:"store"`
		Suggest struct {
			Current      int `json:"current"`
			TimeInMillis int `json:"time_in_millis"`
			Total        int `json:"total"`
		} `json:"suggest"`
		Translog struct {
			Operations  int `json:"operations"`
			SizeInBytes int `json:"size_in_bytes"`
		} `json:"translog"`
		Warmer struct {
			Current           int `json:"current"`
			Total             int `json:"total"`
			TotalTimeInMillis int `json:"total_time_in_millis"`
		} `json:"warmer"`
	} `json:"total"`
}
