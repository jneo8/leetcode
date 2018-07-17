"""See https://leetcode.com/problems/longest-substring-without-repeating-characters/description/"""
from collections import deque

class Solution:
	def lengthOfLongestSubstring(self, s):
		"""
		:type s: str
		:rtype: int
		"""
		if len(s) < 2:
			return len(s)
		dq = deque()
		max_length = 0
		for c in s:
			print(f"{c} {dq}")
			if c in dq:
				while(len(dq) > 0):
					print(f"dq: {dq}")
					if(c == dq.popleft()):
						break
			dq.append(c)
			if len(dq) > max_length:
				max_length = len(dq)

		return max_length

if __name__ == "__main__":

    for params, answer in [
        ("abcabcbb", "abc"),
        ("bbbbb", "b"),
        ("pwwkew", "wke"),
        ("c", "c"),
        ("", ""),
        ("au", "au"),
        ("aab", "ab"),
    ]:
        print("\n----\n")
        print(f"Start '{params}'.")
        result = Solution().lengthOfLongestSubstring(s=params)
        assert result == len(answer)
    print("\n ---- \n Success!!!!")
