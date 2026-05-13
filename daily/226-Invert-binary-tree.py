class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        def dfs(node):
            if not node:
                return
            node.left, node.right = node.right, node.left
            dfs(node.left)
            dfs(node.right)
            return node
        return dfs(root)