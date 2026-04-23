class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        temp = []
        for num in nums:
            if num == val:
                continue
            temp.append(num)
        nums[: len(temp) + 1] = temp
        return len(temp)