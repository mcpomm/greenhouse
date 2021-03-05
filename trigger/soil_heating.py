from flask import Blueprint
from flask import request

soil_heating_bp = Blueprint("soil_heating", __name__)


@soil_heating_bp.route("/trigger/", methods=["POST"])
def trigger():
    print(request.is_json)
    content = request.get_json()
    print(content)
    return "soil heating triggered"