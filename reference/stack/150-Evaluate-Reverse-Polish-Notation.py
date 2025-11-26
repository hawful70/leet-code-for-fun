class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []

        for token in tokens:
            if token in {\+\, \-\, \*\, \/\}:
                a = stack.pop()
                b = stack.pop()
                if token == \+\:
                    stack.append(b + a)
                elif token == \-\:
                    stack.append(b - a)
                elif token == \*\:
                    stack.append(b * a)
                elif token == \/\:
                    # Use integer division and truncate towards zero
                    stack.append(int(b / a))
            else:
                stack.append(int(token))
        
        return stack[-1]