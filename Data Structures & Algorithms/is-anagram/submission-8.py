class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        hashmap_s = {}
        hashmap_t = {}
        for i, v in enumerate(s):
            hashmap_s[v]= hashmap_s.get(v, 0) +1
        for i, v in enumerate(t):
            hashmap_t[v]= hashmap_t.get(v, 0) +1 
        for k, v in hashmap_s.items():
            if not(k in hashmap_t.keys() and v == hashmap_t.get(k, False)):
                return False
        return True
