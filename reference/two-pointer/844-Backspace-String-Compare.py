class Solution:
    def backspaceCompare(self, s: str, t: str) -> bool:
        
        def stack_string_compare(string):
            stack = []
            for char in string:
                if char != "#":
                    stack.append(char)
                elif stack:
                    stack.pop()
            return stack
        
        return stack_string_compare(s) == stack_string_compare(t)