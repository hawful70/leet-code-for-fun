class MyStack:

    def __init__(self):
        self.stacks = []

    def push(self, x: int) -> None:
        self.stacks.append(x)

    def pop(self) -> int:
        top = self.stacks[-1]
        self.stacks.pop()
        return top

    def top(self) -> int:
        return self.stacks[-1]

    def empty(self) -> bool:
        return len(self.stacks) == 0
