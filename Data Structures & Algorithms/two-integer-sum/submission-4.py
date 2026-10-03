class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashmap = {}
        for i, v in enumerate(nums):
            compared_number = target - v
            
            if compared_number in hashmap:
                return [hashmap[compared_number], i]
            hashmap[v]=i
        return []
		# return 
		