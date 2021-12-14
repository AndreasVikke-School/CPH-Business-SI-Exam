import pika
import psycopg2

def main():
    rabbitmqConnection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
    channel = rabbitmqConnection.channel()

    postgresconnection = psycopg2.connect(
    host="localhost",
    database="suppliers",
    user="postgres",
    password="P@ssword!")

    cursor = postgresconnection.cursor()

    channel.queue_declare(queue='BorrowerQueue')

    def callback(ch, method, properties, body):
        print(" [x] Received %r" % body)
        sql = """INSERT INTO loan(loaner_id, entity_id)
                VALUES(%s, %s)"""
        cursor.execute(sql, (body[0], body[1]))


    channel.basic_consume(queue='hello',
                        auto_ack=True,
                        on_message_callback=callback)

if __name__ == '__main__':
    main()





