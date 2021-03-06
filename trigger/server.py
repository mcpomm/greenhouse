from flask import Flask

from soil_watering import soil_watering_bp
from soil_heating import soil_heating_bp

from gpio import initialize_gpio_pins

initialize_gpio_pins()

app = Flask(__name__)
app.register_blueprint(soil_watering_bp, url_prefix="/soil-watering")
app.register_blueprint(soil_heating_bp, url_prefix="/soil-heating")


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5400)
