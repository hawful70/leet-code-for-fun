class Solution:
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        results = []
        def dfs(node):
            if not node:
                return
            results.append(node.val)
            dfs(node.left)
            dfs(node.right)
        dfs(root)
        return results