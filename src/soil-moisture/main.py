import machine
from machine import Pin
from machine import ADC
import urandom
import ujson
import utime
import ntptime
try:
    import usocket as socket
except:
    import socket

moisture = ADC(0)

FullDry = 566
FullWet = 137

ntptime.settime()
machine.RTC().datetime()


soilmoisture = {
    "ID":     str(urandom.getrandbits(30)),
    "Name":   "Soil Moisture",
    "Value":  "",
    "Unit":   "%",
    "Time":   ""
}

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind(('', 5200))
s.listen(5)


def getPercentage(x, wet, dry):
    return round(float((x - dry) * 100 / (wet - dry)), 2)


def getMoisture():
    readtime = 946684800 + utime.time()
    moisture_value = moisture.read()
    print(moisture_value)
    print('%.2f' % getPercentage(moisture_value, FullWet, FullDry), "%")
    soilmoisture["Time"] = str(readtime)
    soilmoisture["Value"] = str(
        '%.2f' % getPercentage(moisture_value, FullWet, FullDry))
    return ujson.dumps(soilmoisture)


while True:
    conn, addr = s.accept()
    request = conn.recv(1024)
    if 'favicon' in str(request):
        continue
    print('Got a connection from %s' % str(addr))
    request = str(request)
    print('Content = %s' % request)
    response = getMoisture()
    conn.send('HTTP/1.1 200 OK\n')
    conn.send('Content-Type: application/json\n')
    conn.send('Connection: close\n\n')
    conn.sendall(response)
    conn.close()
