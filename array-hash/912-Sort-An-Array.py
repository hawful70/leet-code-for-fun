class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        self.mergeSort(nums)
        return nums

    def mergeSort(self, nums: List[int]) -> None:
        if len(nums) > 1:
            # split array into 2 parts: left array and right arrays
            n = len(nums)
            mid = n // 2
            left = nums[:mid]
            right = nums[mid:]

            self.mergeSort(left)
            self.mergeSort(right)

            # merge left array and right array into 1 part
            i = j = k = 0
            while i < len(left) and j < len(right):
                if left[i] < right[j]:
                    nums[k] = left[i]
                    i += 1
                else:
                    nums[k] = right[j]
                    j += 1
                k += 1
            while i < len(left):
                nums[k] = left[i]
                i += 1
                k += 1

            while j < len(right):
                nums[k] = right[j]
                j += 1
                k += 1

        def helper(head, tail):
            if head >= tail:
                return
            l, r = head, tail
            # m = (r - l) // 2 + l
            m = random.randint(l, r)
            pivot = nums[m]
            while r >= l:
                while r >= l and nums[l] < pivot:
                    l += 1
                while r >= l and nums[r] > pivot:
                    r -= 1
                if r >= l:
                    nums[l], nums[r] = nums[r], nums[l]
                    l += 1
                    r -= 1
            helper(head, r)
            helper(l, tail)

        helper(0, len(nums)-1)
