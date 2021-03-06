from flask import Blueprint
from flask import request
from gpio import air_heating

import os
import _thread

air_heating_bp = Blueprint("air_heating", __name__)


@air_heating_bp.route("/trigger/", methods=["POST"])
def trigger():
    content = request.get_json()
    duration = content["duration"]
    _thread.start_new_thread(air_heating, (int(duration),))
    return "air heating triggered"