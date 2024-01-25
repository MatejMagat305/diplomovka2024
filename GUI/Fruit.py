
from kivy.uix.widget import Widget


class Fruit(Widget):
    def move(self, new_pos):
        self.pos = new_pos
