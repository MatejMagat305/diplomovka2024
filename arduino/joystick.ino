const int pinVert = A0;
const int pinHoriz = A1;
const int pinSelect = 2;

void setup() {
  Serial.begin(9600);
  pinMode(pinSelect, INPUT_PULLUP);
}

void loop() {
  int vertValue = analogRead(pinVert);
  int horizValue = analogRead(pinHoriz);
  int selectButton = digitalRead(pinSelect);
  Serial.print("Vertical: ");
  Serial.print(vertValue);
  Serial.print("\tHorizontal: ");
  Serial.print(horizValue);
  Serial.print("\tButton: ");
  Serial.println(selectButton);
  delay(100);
}

