from flask import Flask

from executor_soil_watering import soil_watering_bp
from executor_soil_heating import soil_heating_bp
from executor_fan import fan_bp
from executor_air_heating import air_heating_bp
from executor_tank_filling import tank_filling_bp

from gpio import initialize_gpio_pins

initialize_gpio_pins()

app = Flask(__name__)
app.register_blueprint(soil_watering_bp, url_prefix="/soil-watering")
app.register_blueprint(soil_heating_bp, url_prefix="/soil-heating")
app.register_blueprint(fan_bp, url_prefix="/fan")
app.register_blueprint(air_heating_bp, url_prefix="/air-heating")
app.register_blueprint(tank_filling_bp, url_prefix="/tank-filling")


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5400)
