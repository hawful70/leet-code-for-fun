from bisect import bisect_right
from typing import List


class TopVotedCandidate:
    def __init__(self, persons: List[int], times: List[int]):
        self.times = times
        self.leaders = []

        counts = {}
        leader = -1
        max_votes = 0

        for person in persons:
            counts[person] = counts.get(person, 0) + 1
            if counts[person] >= max_votes:
                max_votes = counts[person]
                leader = person
            self.leaders.append(leader)

    def q(self, t: int) -> int:
        idx = bisect_right(self.times, t) - 1
        if idx < 0:
            return -1
        return self.leaders[idx]
