class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        results = []
        def dfs(node):
            if not node:
                return
            dfs(node.left)
            results.append(node.val)
            dfs(node.right)
        dfs(root)
        return results