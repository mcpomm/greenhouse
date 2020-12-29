from flask import Flask, url_for
from flask import Response
import time
import board
import adafruit_dht
import random
import os
import logging
import json
import time
import datetime

PORT = 5100
dhtDevice = adafruit_dht.DHT22(board.D17)

app = Flask(__name__)
hid = random.randint(1000, 1000000)


def _get_humidity():
    humidity = {
        "ID":     str(hid),
        "Name":   "Humidity",
        "Value":  "",
        "Unit":   "%",
        "Time":   str(time.mktime(datetime.datetime.today().timetuple())).split('.')[0]
    }
    try:
        humidity["Value"] = str(dhtDevice.humidity)
        print("Humidity: {}% ".format(humidity["Value"]))
        return json.dumps(humidity), 200

    except RuntimeError as error:
        # Errors happen fairly often, DHT's are hard to read, just keep going
        print(error.args[0])
        return error.args[0], 500

    except Exception as error:
        dhtDevice.exit()
        raise error


@app.route('/')
def api_root():
    return 'Humidity sensor DHT11'


@app.route('/humidity', methods=['GET'])
def api_humidity():
    data, code = _get_humidity()
    resp = Response(data, status=code, mimetype='application/json')
    return resp


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=5100, debug=False)
