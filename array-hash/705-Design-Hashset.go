const capacity = 10000 // Size of the bucket array

type MyHashSet struct {
	buckets [][]int
}

// NewHashSet initializes a new HashSet
func Constructor() MyHashSet {
	return MyHashSet{
		buckets: make([][]int, capacity),
	}
}

// hash function computes index from the key
func (hs *MyHashSet) hash(key int) int {
	return key % capacity
}

// Add inserts a key into the HashSet if not already present
func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}
	index := this.hash(key)
	this.buckets[index] = append(this.buckets[index], key)
}

// Remove deletes a key from the HashSet if present
func (this *MyHashSet) Remove(key int) {
	index := this.hash(key)
	bucket := this.buckets[index]
	for i, v := range bucket {
		if v == key {
			// Remove item using the "append" technique
			this.buckets[index] = append(bucket[:i], bucket[i+1:]...)
			return
		}
	}
}

// Contains checks if a key exists in the HashSet
func (this *MyHashSet) Contains(key int) bool {
	index := this.hash(key)
	for _, v := range this.buckets[index] {
		if v == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */