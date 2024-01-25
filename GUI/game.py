

from Fruit import Fruit
from SmartGrid import smartGrid
from SnakeHead import SnakeHead
from SnakeTail import SnakeTail
from  values import *

from kivy.uix.widget import Widget
from kivy.properties import ObjectProperty, NumericProperty, StringProperty
from kivy.core.window import Window
from kivy.clock import Clock
from random import randint

class SnakeGame(Widget):
    head = ObjectProperty(None)
    fruit = ObjectProperty(None)
    score = NumericProperty(0)
    player_size = NumericProperty(PLAYER_SIZE)
    game_over = StringProperty("")

    def __init__(self):
        super(SnakeGame, self).__init__()

        Window.size = (WINDOW_WIDTH, WINDOW_HEIGHT)
        Window.bind(on_key_down=self.key_action)

        if PLAYER_SIZE < 3:
            raise ValueError("Player size should be at least 3 px")

        if WINDOW_HEIGHT < 3 * PLAYER_SIZE or WINDOW_WIDTH < 3 * PLAYER_SIZE:
            raise ValueError("Window size must be at least 3 times larger than player size")

        self.timer = Clock.schedule_interval(self.refresh, GAME_SPEED)
        self.tail = []
        self.restart_game()

    def restart_game(self):
        self.occupied = smartGrid()

        # resets the timer
        self.timer.cancel()
        self.timer = Clock.schedule_interval(self.refresh, GAME_SPEED)

        self.head.reset_pos()
        self.score = 0

        for block in self.tail:
            self.remove_widget(block)

        self.tail = []

        self.tail.append(
            SnakeTail(
                pos=(self.head.pos[0] - PLAYER_SIZE, self.head.pos[1]),
                size=(self.head.size)
            )
        )
        self.add_widget(self.tail[-1])
        self.occupied[self.tail[-1].pos] = True

        self.tail.append(
            SnakeTail(
                pos=(self.head.pos[0] - 2 * PLAYER_SIZE, self.head.pos[1]),
                size=(self.head.size)
            )
        )
        self.add_widget(self.tail[-1])
        self.occupied[self.tail[1].pos] = True

        self.spawn_fruit()

    def refresh(self, dt):
        if not (0 <= self.head.pos[0] < WINDOW_WIDTH) or \
           not (0 <= self.head.pos[1] < WINDOW_HEIGHT):
            self.restart_game()
            return

        if self.occupied[self.head.pos] is True:
            self.restart_game()
            return

        # move the tail
        self.occupied[self.tail[-1].pos] = False
        self.tail[-1].move(self.tail[-2].pos)

        for i in range(2, len(self.tail)):
            self.tail[-i].move(new_pos=(self.tail[-(i + 1)].pos))

        self.tail[0].move(new_pos=self.head.pos)
        self.occupied[self.tail[0].pos] = True

        self.head.move()

        if self.head.pos == self.fruit.pos:
            self.score += 1
            self.tail.append(
                SnakeTail(
                    pos=self.head.pos,
                    size=self.head.size))
            self.add_widget(self.tail[-1])
            self.spawn_fruit()

    def spawn_fruit(self):
        roll = self.fruit.pos
        found = False
        while not found:
            roll = [PLAYER_SIZE *
                    randint(0, int(WINDOW_WIDTH / PLAYER_SIZE) - 1),
                    PLAYER_SIZE *
                    randint(0, int(WINDOW_HEIGHT / PLAYER_SIZE) - 1)]

            if self.occupied[roll] is True or \
                    roll == self.head.pos:
                continue

            found = True

        self.fruit.move(roll)

    def key_action(self, *args):
        """This handles user input
        """

        command = list(args)[3]
        print(args)

        if command == 'w' or command == 'up':
            self.head.orientation = (0, PLAYER_SIZE)
        elif command == 's' or command == 'down':
            self.head.orientation = (0, -PLAYER_SIZE)
        elif command == 'a' or command == 'left':
            self.head.orientation = (-PLAYER_SIZE, 0)
        elif command == 'd' or command == 'right':
            self.head.orientation = (PLAYER_SIZE, 0)
        elif command == 'r':
            self.restart_game()
        elif command == 'q':
            raise SystemExit()

