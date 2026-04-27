class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if not root:
            return []
        results = []
        queue = deque([root])
        while queue:
            len_size = len(queue)
            level = []
            for _ in range(len_size):
                node = queue.popleft()
                level.append(node.val)

                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            results.append(level)
        return results