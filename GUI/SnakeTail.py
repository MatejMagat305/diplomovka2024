
from kivy.uix.widget import Widget


class SnakeTail(Widget):
    def move(self, new_pos):
        self.pos = new_pos

