"""See https://leetcode.com/problems/two-sum/description/"""


class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        idx = 0
        for x in nums.copy():
            nums.remove(x)
            for y in nums:
                if x + y == target:
                    return [idx, nums.index(y) + idx + 1]
            idx += 1


if __name__ == "__main__":
    print(
        Solution().twoSum(nums=[3, 2, 4], target=6)
    )
    print(
        Solution().twoSum(nums=[2, 7, 11, 15], target=9)
    )
