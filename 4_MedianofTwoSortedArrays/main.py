"See  https://leetcode.com/problems/median-of-two-sorted-arrays/description/"
import statistics

class Solution:
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        nums3 = nums1 + nums2
        return statistics.median(nums3)


if __name__ == "__main__":

    for nums1, nums2, answer in [
            ([1, 3], [2], 2),
            ([1, 2], [3, 4], 2.5),
    ]:
        print("\n----\n")
        print(f"Start '{nums1} {nums2}'.")
        result = Solution().findMedianSortedArrays(nums1=nums1, nums2=nums2)
        assert result == answer
    print("\n ---- \n Success!!!!")
