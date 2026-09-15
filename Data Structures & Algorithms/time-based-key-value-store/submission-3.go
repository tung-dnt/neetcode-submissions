type pair struct {
	value string
	timestamp int
}

type TimeMap struct {
	store map[string][]pair
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.store[key] == nil {
		this.store[key] = make([]pair, 0)
	}
	this.store[key] = append(this.store[key], pair{value,timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	pairs, exists := this.store[key]
	if !exists {
		return ""
	}
	l, r := 0, len(pairs)-1
	res := ""
	for l <= r {
		mid := int(l + (r-l)/2)
		if pairs[mid].timestamp == timestamp {
			return pairs[mid].value
		} else if pairs[mid].timestamp < timestamp {
			res = pairs[mid].value
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}
