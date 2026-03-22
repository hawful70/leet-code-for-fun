class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        # if len(lists) == 0 or (len(lists) == 1 and lists[0] == None) :
        #     return None
        # if len(lists) == 1:
        #     return lists[0]

        # firstList = lists[0]
        # result = self.mergeTwoList(firstList, None)
        # for i in range(1, len(lists)):
        #     currentList = lists[i]
        #     result = self.mergeTwoList(result, currentList)
        if not lists:
            return None
        while len(lists) > 1:
            mergeLists = []
            for i in range(0, len(lists), 2):
                list1 = lists[i]
                list2 = lists[i + 1] if i + 1 < len(lists) else None
                mergeLists.append(self.mergeTwoList(list1, list2))
            lists = mergeLists

        return lists[0]
        

    def mergeTwoList(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0)
        curr = dummy
        while list1 and list2:
            if list1.val <= list2.val:
                curr.next = list1
                list1 = list1.next
            else:
                curr.next = list2
                list2 = list2.next
            curr = curr.next
        
        curr.next = list1 if list1 else list2
        return dummy.next