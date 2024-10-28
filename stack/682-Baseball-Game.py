class Solution:
    def calPoints(self, operations: List[str]) -> int:
        ans = []
        for operation in operations:
            if operation.isdigit() or operation.lstrip('-').isdigit():
                ans.append(int(operation))
            elif operation.isalpha() and operation == 'D':
                current_score = ans[-1] * 2
                ans.append(current_score)
            elif operation.isalpha() and operation == 'C':
                ans.pop()
            else:
                if len(ans) >= 2:
                    current_score = ans[-1] + ans[-2]
                    ans.append(current_score)
        
        return 0 if len(ans) == 0 else sum(ans)
