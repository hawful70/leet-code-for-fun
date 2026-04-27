class Solution:
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        results = []
        def dfs(node):
            if not node:
                return
            dfs(node.left)
            dfs(node.right)
            results.append(node.val)
        dfs(root)
        return results