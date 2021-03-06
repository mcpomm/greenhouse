from flask import Blueprint
from flask import request
from gpio import fan_activating

import os
import _thread

fan_bp = Blueprint("fan", __name__)


@fan_bp.route("/trigger/", methods=["POST"])
def trigger():
    content = request.get_json()
    duration = content["duration"]
    _thread.start_new_thread(fan_activating, (int(duration),))
    return "fan triggered"