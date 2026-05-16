class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        self.maxDiameter = 0

        def dfs(node):
            if not node:
                return 0
            left = dfs(node.left)
            right = dfs(node.right)

            self.maxDiameter = max(self.maxDiameter, left + right)
            return 1 + max(left, right)
        dfs(root)
        return self.maxDiameter