FROM python:3.7-buster

LABEL org.opencontainers.image.source https://github.com/AndreasVikke-School/cph-business-si-exam

WORKDIR /app
ENV NEOIP="neo4j://neo4j.neo4j-si"
ENV NEOUSER="neo4j"
ENV NEOPASS="P@ssword!"

COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

COPY . .

CMD [ "python3", "neo4j_server.py"]