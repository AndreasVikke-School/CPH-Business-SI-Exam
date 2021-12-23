FROM python:3.8

LABEL org.opencontainers.image.source https://github.com/AndreasVikke-School/cph-business-si-exam

WORKDIR /app
ENV RABBITIP="rabbitmq.rabbitmq-si"
ENV RABBITUSER="rabbitmq"
ENV RABBITPASS="P@ssword!"
ENV POSTGRESIP="postgres-service.services-si"

COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

COPY . .

CMD ["python3", "consume.py"]