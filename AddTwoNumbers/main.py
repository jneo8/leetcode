# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        l3_list = []
        buffer_v = 0
        while l1 is not None or l2 is not None or buffer_v > 0:
            l1_val = l1.val if l1 is not None else 0
            l2_val = l2.val if l2 is not None else 0
            l3_val = l1_val + l2_val + buffer_v
            buffer_v = 0
            print(l3_val)
            if l3_val >= 10:
                buffer_v += 1
                l3_val -= 10
            l3_list.append(ListNode(x=l3_val))
            if l1 is not None:
                l1 = l1.next
            if l2 is not None:
                l2 = l2.next

        for l3_1, l3_2 in zip(l3_list[:-1], l3_list[1:]):
            l3_1.next = l3_2
        print(l3_list)
        return l3_list[0]
