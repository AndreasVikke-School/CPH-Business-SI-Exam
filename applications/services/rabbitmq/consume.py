import pika
import loan_pb2

def main():
    rabbitmqConnection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
    channel = rabbitmqConnection.channel()


    channel.queue_declare(queue='BorrowerQueue')

    def callback(ch, method, properties, body):
        print(" [x] Received %r" % body)
        loan_pb2.Loan.userId = body[0]
        loan_pb2.Loan.entityId = body[1]
        loan_pb2.Loan.status = 0

    channel.basic_consume(queue='BorrowerQueue',
                        auto_ack=True,
                        on_message_callback=callback)

if __name__ == '__main__':
    main()





