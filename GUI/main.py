
from kivy.app import App
from game import SnakeGame


class snakeApp(App):
    def build(self):
        game = SnakeGame()
        return game


if __name__ == '__main__':
    snakeApp().run()
