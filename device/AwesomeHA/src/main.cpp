#include <ArduinoJson.h>
#include <ESP8266WiFi.h>
#include <ESP8266mDNS.h>
#include <WiFiClient.h>
#include <ESP8266WebServer.h>
#include <ElegantOTA.h>
 
const char* ssid = "CentNet";
const char* password = "n00dl3xyz";
const String ALLOWED_URL = "http://localhost:4000";
ESP8266WebServer server(81);
unsigned long ota_progress_millis = 0;
void onOTAStart() {
  // Log OTA start
  Serial.println("OTA update started");
}
void onOTAProgress(size_t current, size_t final) {
  // Log every 1 second
  if (millis() - ota_progress_millis > 1000) {
    ota_progress_millis = millis();
    Serial.printf("OTA Progress Current: %u bytes, Final: %u bytes\n", current, final);
  }
}
void onOTAEnd(bool success) {
  // Log when OTA has finished
  if (success) {
    Serial.println("OTA update finished successfully!");
  } else {
    Serial.println("There was an error during OTA update!");
  }
  // <Add your own code here>
}

void handleWhatCanItDo() {
  server.sendHeader("Access-Control-Allow-Origin", ALLOWED_URL);
  server.send(200, "application/json", R"({"info":"This is a device to test the API integration system for future projects","type":"boolean","ip":")" + WiFi.localIP().toString() + R"(","commands":[{"name":"on","info":"Turns built in LED on"},{"name":"off","info":"Turns built in LED off"},{"name":"blink","info":"Blinks built in LED for 1s"},{"name":"value","info":"Gets data from sensor"}]})");
}

void handleOn(){
  server.sendHeader("Access-Control-Allow-Origin", ALLOWED_URL);
  server.send(200, "text/plain", "ON");
  digitalWrite(LED_BUILTIN, 0);
}

void handleOff(){
  server.sendHeader("Access-Control-Allow-Origin", ALLOWED_URL); // Allow from any origin
  server.send(200, "text/plain", "OFF");
  digitalWrite(LED_BUILTIN, 1);
}

void handleBlink(){
  server.sendHeader("Access-Control-Allow-Origin", ALLOWED_URL); // Allow from any origin
  server.send(200, "text/plain", "BLINK");
  static bool state = false;
  state = !state;
  digitalWrite(LED_BUILTIN, state ? LOW : HIGH);
}

void handleData(){
  String JString = "{\"value\":" + String(analogRead(A0)) + '}';
  server.sendHeader("Access-Control-Allow-Origin", ALLOWED_URL);
  server.send(200, "application/json", JString);
}

void setup() {
  Serial.begin(115200);
  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);
  while (WiFi.waitForConnectResult() != WL_CONNECTED) {
    delay(500);
  }

  pinMode(LED_BUILTIN, OUTPUT);
  MDNS.begin("awesome000");
  server.on("/", handleWhatCanItDo);
  server.on("/on", handleOn);
  server.on("/off", handleOff);
  server.on("/blink", handleBlink);
  server.on("/value", handleData);
  ElegantOTA.begin(&server);
  ElegantOTA.onStart(onOTAStart);
  ElegantOTA.onProgress(onOTAProgress);
  ElegantOTA.onEnd(onOTAEnd);
  server.begin();
}

void loop() {
  MDNS.update();
  server.handleClient();
  ElegantOTA.loop();
}