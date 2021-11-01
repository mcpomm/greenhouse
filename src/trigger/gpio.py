from gpiozero.pins.native import NativeFactory
from gpiozero import LED
from time import sleep
import logging

logging.basicConfig(
    format="%(asctime)s %(levelname)-8s %(message)s",
    level=logging.INFO,
    datefmt="%Y-%m-%d %H:%M:%S",
)

factory = NativeFactory()

soil_watering_pin = LED(21, pin_factory=factory)  # /soil-watering/trigger
soil_heating_pin = LED(20, pin_factory=factory)  # /soil-heating/trigger
fan_activating_pin = LED(16, pin_factory=factory)  # /fan/trigger
air_heating_pin = LED(12, pin_factory=factory)  # /air-heating/trigger
tank_filling_pin = LED(25, pin_factory=factory)  # /tank-filling/trigger


def initialize_gpio_pins():
    logging.info("Initialize GPIO pins on the board.")

    soil_watering_pin.on()
    soil_heating_pin.on()
    fan_activating_pin.on()
    air_heating_pin.on()
    tank_filling_pin.on()


def soil_watering(trigger_duration):
    _do_trigger(soil_watering_pin, "soil watering", trigger_duration)


def soil_heating(trigger_duration):
    _do_trigger(soil_heating_pin, "soil heating", trigger_duration)


def fan_activating(trigger_duration):
    _do_trigger(fan_activating_pin, "fan", trigger_duration)


def air_heating(trigger_duration):
    _do_trigger(air_heating_pin, "air heating", trigger_duration)


def tank_filling(trigger_duration):
    _do_trigger(tank_filling_pin, "tank filling", trigger_duration)


def _do_trigger(pin, action, duration):
    logging.info(f"Activate {action} for {duration} seconds")
    logging.info(f"Start {action}")
    pin.off()
    sleep(duration)
    pin.on()
    logging.info(f"Stop {action}")
