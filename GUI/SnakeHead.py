
from kivy.uix.widget import Widget
from  values import *
from kivy.vector import Vector


class SnakeHead(Widget):
    orientation = (PLAYER_SIZE, 0)

    def reset_pos(self):
        # positions the player roughly in the middle of the gameboard
        self.pos = \
            [int(WINDOW_WIDTH / 2 - (WINDOW_WIDTH / 2 % PLAYER_SIZE)),
             int(WINDOW_HEIGHT / 2 - (WINDOW_HEIGHT / 2 % PLAYER_SIZE))]
        self.orientation = (PLAYER_SIZE, 0)

    def move(self):
        self.pos = Vector(*self.orientation) + self.pos
