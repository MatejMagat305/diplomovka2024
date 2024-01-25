
from  values import *


class smartGrid:
    def __init__(self):
        self.grid = [[False for _ in range(WINDOW_HEIGHT)]
                     for _ in range(WINDOW_WIDTH)]

    def __getitem__(self, coords):
        return self.grid[coords[0]][coords[1]]

    def __setitem__(self, coords, value):
        self.grid[coords[0]][coords[1]] = value

