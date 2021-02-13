FROM arm32v7/python:3

WORKDIR /usr/src/app

RUN apt-get update \
  && apt-get upgrade -y\
  # && apt-get install libgpiod2 \
  # && pip install --upgrade setuptools \
  && pip install RPI.GPIO \
  && pip install adafruit-circuitpython-dht \
  && pip install flask

COPY humidity.py ./

RUN chmod +x ./humidity.py
CMD ["python","./humidity.py"]