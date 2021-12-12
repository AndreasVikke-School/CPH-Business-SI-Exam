FROM python:3.7-alpine

LABEL org.opencontainers.image.source https://github.com/AndreasVikke-School/cph-business-si-exam

WORKDIR /app

COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

COPY . .

CMD [ "python3", "main.py"]