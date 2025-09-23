from typing import List, Optional


class TreeNode:
    def __init__(self, val: int = 0, left: Optional["TreeNode"] = None, right: Optional["TreeNode"] = None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def findTarget(self, root: Optional[TreeNode], k: int) -> bool:
        values: List[int] = []
        self._inorder(root, values)
        left, right = 0, len(values) - 1
        while left < right:
            total = values[left] + values[right]
            if total == k:
                return True
            if total < k:
                left += 1
            else:
                right -= 1
        return False

    def _inorder(self, node: Optional[TreeNode], acc: List[int]) -> None:
        if node is None:
            return
        self._inorder(node.left, acc)
        acc.append(node.val)
        self._inorder(node.right, acc)

