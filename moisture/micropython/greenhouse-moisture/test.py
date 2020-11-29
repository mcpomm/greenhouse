import time
from time import sleep

AirValue = 560
WaterValue = 162


def myMap(x, in_min, in_max, out_min, out_max):
    return (x-in_min)*(out_max-out_min)/(in_max-in_min) + out_min


while True:
    moisture_value = 363
    print(moisture_value)
    print('%.2f' % myMap(moisture_value, AirValue, WaterValue, 0, 100), "%")
    sleep(3)
