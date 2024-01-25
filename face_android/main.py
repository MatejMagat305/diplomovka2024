from kivy.app import App
from kivy.uix.boxlayout import BoxLayout
from kivy.uix.image import Image
from kivy.clock import Clock
import cv2
from kivy.core.window import Window

class CameraApp(App):
    def build(self):
        self.img = Image()
        layout = BoxLayout(orientation='vertical')
        layout.add_widget(self.img)

        # Vytvorte kameru OpenCV
        self.capture = cv2.VideoCapture(0)

        # Určte veľkosť obrazu z kamery
        Window.size = (self.capture.get(3), self.capture.get(4))

        # Aktualizujte obraz každé 1/30 sekundy (30 fps)
        Clock.schedule_interval(self.update, 1.0 / 30.0)

        return layout

    def update(self, dt):
        # Zachyťte obraz z kamery
        ret, frame = self.capture.read()

        if ret:
            # Konvertujte obraz na formát Kivy a zobrazte ho
            buf1 = cv2.flip(frame, 0)
            buf = buf1.tostring()
            texture1 = Texture.create(size=(frame.shape[1], frame.shape[0]), colorfmt='bgr')
            texture1.blit_buffer(buf, colorfmt='bgr', bufferfmt='ubyte')
            self.img.texture = texture1

            # Detekcia tvárí pomocou OpenCV
            gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
            face_cascade = cv2.CascadeClassifier(cv2.data.haarcascades + 'haarcascade_frontalface_default.xml')
            faces = face_cascade.detectMultiScale(gray, scaleFactor=1.1, minNeighbors=5, minSize=(30, 30))

            # Zobrazte obdĺžniky okolo detekovaných tvárí
            for (x, y, w, h) in faces:
                cv2.rectangle(frame, (x, y), (x+w, y+h), (0, 255, 0), 2)

            cv2.imshow('Face Detection', frame)

    def on_stop(self):
        # Ukončite zachytenie kamery pri zatvorení aplikácie
        self.capture.release()
        cv2.destroyAllWindows()

if __name__ == '__main__':
    CameraApp().run()
