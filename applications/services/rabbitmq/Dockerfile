FROM python:3.8

RUN pip install pika

LABEL org.opencontainers.image.source https://github.com/AndreasVikke-School/cph-business-si-exam

WORKDIR /app

COPY . .

CMD ["python3", "consume.py"]