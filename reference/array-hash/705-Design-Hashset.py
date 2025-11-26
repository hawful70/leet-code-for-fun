class MyHashSet:
    def __init__(self):
        self.bucket_size = 769
        self.buckets = [[] for _ in range(self.bucket_size)]

    def _hash(self, key: int) -> int:
        return key % self.bucket_size

    def add(self, key: int) -> None:
        index = self._hash(key)
        if key not in self.buckets[index]:
            self.buckets[index].append(key)

    def remove(self, key: int) -> None:
        index = self._hash(key)
        if key in self.buckets[index]:
            self.buckets[index].remove(key)

    def contains(self, key: int) -> bool:
        index = self._hash(key)
        return key in self.buckets[index]