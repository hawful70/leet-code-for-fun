class Solution:
    def minDepth(self, root: Optional[TreeNode]) -> int:
        def dfs(node):
            if not node:
                return 0
            left = dfs(node.left)
            right = dfs(node.right)

            if not node.left:
                return 1 + right
            if not node.right:
                return 1 + left

            return 1 + min(left, right)
        return dfs(root)