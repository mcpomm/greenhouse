from machine import Pin
from time import sleep
import ujson
try:
    import usocket as socket
except:
    import socket

# # ESP32 GPIO 26
# relay = Pin(26, Pin.OUT)

off = 1
on = 0

# ESP8266 GPIO 5
relay = Pin(5, Pin.OUT)
relay.value(off)

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind(('', 5400))
s.listen(5)

print('Initial GPIO pin value %s' % relay.value())

response = {
    "pinValue": relay.value()
}


def triggerPump(pumpDuration):
    print("Start pump.")
    startPump()
    _logCurrentPumpValue()
    sleep(pumpDuration)
    print("Stop pump.")
    stopPump()
    _logCurrentPumpValue()
    return ujson.dumps(response)


def startPump():
    relay.value(on)


def stopPump():
    relay.value(off)


def _logCurrentPumpValue():
    print('Current GPIO pin value %s' % relay.value())


while True:
    conn, addr = s.accept()
    request = conn.recv(1024)
    if 'favicon' in str(request):
        continue
    print('Got a connection from %s' % str(addr))
    request = str(request)
    print('Content = %s' % request)
    split = request.split(
        "\"pumpDuration\"")
    pumpDuration = split[-1].replace(":", "").replace("\"",
                                                      "").replace("'", "").replace("}", "").replace(" ", "")
    print(int(pumpDuration))
    res = triggerPump(int(pumpDuration))
    conn.send('HTTP/1.1 200 OK\n')
    conn.send('Content-Type: application/json\n')
    conn.send('Connection: close\n\n')
    conn.sendall(res)
    conn.close()
