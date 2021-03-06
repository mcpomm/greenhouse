from flask import Blueprint
from flask import request
from gpio import tank_filling

import os
import _thread

tank_filling_bp = Blueprint("tank_filling", __name__)


@tank_filling_bp.route("/trigger/", methods=["POST"])
def trigger():
    content = request.get_json()
    duration = content["duration"]
    _thread.start_new_thread(tank_filling, (int(duration),))
    return "tank filling triggered"