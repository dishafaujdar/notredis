''' a way to minimize the data movement required during scale-ups or scale-downs, 
and this is where Consistent Hashing fits in and minimizes the required data transfer.'''

import hashlib
import bisect

class ConsistentHashing:
    def __init__(self):
        self.node = []

    def hash(self,key):
        return int(hashlib.sha256(key.encode('utf-8')).hexdigest(), 16)

    def add_server(self,server):
        server_hash = self.hash(server)

        bisect.insort(self.node, (server_hash, server))

    def get_server(self,key):
        file_hash = self.hash(key)

        index = bisect.bisect_left(self.node, (file_hash,""))

        if index == len(self.node):
            index = 0

        return self.node[index][1]


ring = ConsistentHashing()
ring.add_server("server1")
ring.add_server("server2")
ring.add_server("server3")

print(ring.get_server("key1"))
print(ring.get_server("key2"))
print(ring.get_server("key3"))
print(ring.get_server("key4"))
print(ring.get_server("key5"))
print(ring.get_server("key6"))
print(ring.get_server("key7"))