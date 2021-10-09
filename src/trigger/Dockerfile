FROM arm32v7/python:3.8.7
WORKDIR /usr/src/app
COPY . .
RUN pip3 install -r requirements.txt
CMD ["gunicorn", "--bind", "0.0.0.0:5400", "wsgi:app"]