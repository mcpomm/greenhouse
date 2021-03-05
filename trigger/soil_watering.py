from flask import Blueprint
from flask import request
from gpio import soil_watering

import os
import _thread

soil_watering_bp = Blueprint("soil_watering", __name__)


@soil_watering_bp.route("/trigger/", methods=["POST"])
def trigger():
    content = request.get_json()
    duration = content["duration"]
    _thread.start_new_thread(soil_watering, (int(duration),))
    return "soil watering triggered"