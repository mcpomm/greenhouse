from flask import Blueprint
from flask import request
from gpio import soil_heating

import os
import _thread

soil_heating_bp = Blueprint("soil_heating", __name__)


@soil_heating_bp.route("/trigger/", methods=["POST"])
def trigger():
    content = request.get_json()
    duration = content["duration"]
    _thread.start_new_thread(soil_heating, (int(duration),))
    return "soil heating triggered"