FROM arm32v7/python:3

WORKDIR /usr/src/app

RUN apt-get update \
  && apt-get upgrade -y\
  # && apt-get install libgpiod2 \
  # && pip install --upgrade setuptools \
  && pip install RPI.GPIO \
  && pip install adafruit-circuitpython-dht \
  && pip install flask \
  && pip install gunicorn

COPY humidity.py ./
COPY wsgi.py ./

CMD ["gunicorn", "--bind", "0.0.0.0:5100", "wsgi:app"]