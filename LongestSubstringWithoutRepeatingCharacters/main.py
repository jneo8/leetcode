class Solution:
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        if s == "":
            return 0
        elif len(s) == len("".join(set(s))):
            return len(s)
        else:
            longest_substring = s[0]
        for idx_1, w1 in enumerate(s):
            print(f"{s} - {idx_1} {w1}")
            tmp_string = ""
            if len(longest_substring) > len(s[idx_1:]):
                break
            for idx_2, w2 in enumerate(s[idx_1:]):
                print(f"w2: {w2}")
                if w2 in tmp_string:
                    ## find longest, break
                    if len(tmp_string) > len(longest_substring):
                        longest_substring = tmp_string
                    break
                else:
                    tmp_string += w2
            if len(tmp_string) > len(longest_substring):
                longest_substring = tmp_string
                print(f"tmp_string: {tmp_string}")
                print(f"longest_substring: {longest_substring}")

        print(f"longest_substring: {longest_substring}")
        return len(longest_substring)



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
