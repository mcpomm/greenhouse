from machine import Pin
from time import sleep
import ujson
try:
    import usocket as socket
except:
    import socket

# # ESP32 GPIO 26
# relay = Pin(26, Pin.OUT)

# The reason why the values are reversed (off = 1 and on = 0) is described here: https://arduino.stackexchange.com/questions/54502/5v-relay-module

# Or in short, most relays are delivered in normally closed (NC) mode. This means that the circuit is closed and the relay is activated. To switch to normally open (NO), i.e. to deactivate the relay, the trigger pin must be energized (1 / on). If the relay is then to be reactivated, i.e. the manget lever of the relay is to jump back again, the trigger pin must no longer carry current, i.e. (0 / off).

# Therefore:
# on = normally open(NO)
# and
# off = normally closed(NC)

# The corresponding cable connections from (+) to the relay is as follows:
# common -> to the load (+)
# normally open(NO) -> to the power source (+)

# You can also set the values correctly but then you have to change the cable connections as follows:
# common -> to the consumer (+)
# normally colsed(NC) -> to the power source (+)
# This has the consequence that the consumer is switched on continuously in case of a defective or not connected NodeMCU module/trigger.

# Since the consumer, in this case the watering pump, is to activate only briefly by the trigger, the values were simply swapped.

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
