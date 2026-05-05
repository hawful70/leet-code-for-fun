class Node:
    def __init__(self, key, val):
        self.key = key
        self.val = val
        self.freq = 1
        self.prev = self.next = None

class DLinkedList:
    def __init__(self):
        self.head = Node(0, 0)   # dummy
        self.tail = Node(0, 0)   # dummy
        self.head.next = self.tail
        self.tail.prev = self.head
        self.size = 0

    def add_front(self, node):
        nxt = self.head.next
        self.head.next = node
        node.prev, node.next = self.head, nxt
        nxt.prev = node
        self.size += 1

    def remove(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev
        self.size -= 1

    def remove_last(self):
        if self.size == 0:
            return None
        last = self.tail.prev
        self.remove(last)
        return last


class LFUCache:
    def __init__(self, capacity: int):
        self.cap = capacity
        self.min_freq = 0
        self.key_to_node = {}                          # key → Node
        self.freq_to_list = defaultdict(DLinkedList)   # freq → DLinkedList

    def _promote(self, node):
        """Move a node up one frequency bucket."""
        self.freq_to_list[node.freq].remove(node)
        if self.freq_to_list[node.freq].size == 0 and node.freq == self.min_freq:
            self.min_freq += 1          # bucket empty + was the minimum → raise floor
        node.freq += 1
        self.freq_to_list[node.freq].add_front(node)

    def get(self, key: int) -> int:
        if key not in self.key_to_node:
            return -1
        node = self.key_to_node[key]
        self._promote(node)
        return node.val

    def put(self, key: int, value: int) -> None:
        if self.cap == 0:
            return

        if key in self.key_to_node:
            node = self.key_to_node[key]
            node.val = value
            self._promote(node)
        else:
            if len(self.key_to_node) >= self.cap:
                # Evict LFU node (tail of the minimum-frequency list)
                evicted = self.freq_to_list[self.min_freq].remove_last()
                del self.key_to_node[evicted.key]

            new_node = Node(key, value)
            self.key_to_node[key] = new_node
            self.freq_to_list[1].add_front(new_node)
            self.min_freq = 1           # new nodes always start at freq = 1