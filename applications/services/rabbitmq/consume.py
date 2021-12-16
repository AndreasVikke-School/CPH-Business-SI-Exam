import json
import pika
import grpc
import loan_pb2
import loan_pb2_grpc

def main():
    credentials = pika.PlainCredentials("rabbitmq", "P@ssword!")
    parameters = pika.ConnectionParameters('localhost', 5672, '/', credentials)
    rabbitmqConnection = pika.BlockingConnection(parameters)
    channel = rabbitmqConnection.channel()


    channel.queue_declare(queue='LoanQueue')
                        
    chan = grpc.insecure_channel('localhost:50051')
    stub = loan_pb2_grpc.LoanServiceStub(chan)
    def callback(ch, method, properties, body):
        print(" [x] Received %r" % json.loads(body))
        data = json.loads(body)
        loan = loan_pb2.Loan(
            userId = data['userId'],
            entityId = data['entityId'],
            status = 0
        )
        try:
            response = stub.CreateLoan(loan)
            print("Created Loan for user with id: {0} on entity with id: {1}".format(response.userId, response.entityId))
        except:
            print("User not found")
    
    channel.basic_consume(queue='LoanQueue',
                        auto_ack=True,
                        on_message_callback=callback)
    
    channel.start_consuming()
    chan.close()


if __name__ == '__main__':
    main()





