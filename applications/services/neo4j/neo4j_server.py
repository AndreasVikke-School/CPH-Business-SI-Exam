from concurrent import futures
import grpc
import neo4j
import book_pb2
import book_pb2_grpc
import time
import threading
import logging
from neo4j import GraphDatabase
import os

#region Cypherquery methods
def get_book_by_title(db, title):
    query = """match (b:Book {title: $input}) RETURN b as book LIMIT 1;"""
    with db.session() as session:
        book = session.run(query, input=title)
        return book.single()[0]

def get_book_by_title_with_more(db, title):
    query = """
    MATCH (b:Book {title: $input})
    MATCH (b)-[:WRITTEN_BY]->(a:Author)
    MATCH (b)-[:WRITTEN_IN]->(y:Year)
    RETURN b, a, y LIMIT 1;
    """
    with db.session() as session:
        book = session.run(query, input=title)
        return book.single()

def get_all_books(db):
    query = """
    MATCH (b:Book)
    MATCH (b)-[:WRITTEN_BY]->(a:Author)
    MATCH (b)-[:WRITTEN_IN]->(y:Year)
    RETURN b, a, y;
    """
    lst = []
    with db.session() as session:
        books = session.run(query)
        for item in books:
            lst.append(item)
        return lst
    
def search_book(db, title):
    query = """
    MATCH (b:Book)
    WHERE b.title =~ $input
    MATCH (b)-[:WRITTEN_BY]->(a:Author)
    MATCH (b)-[:WRITTEN_IN]->(y:Year)
    RETURN b, a, y;
    """
    lst = []
    with db.session() as session:
        books = session.run(query, input="(?i).*" + title + ".*")
        for item in books:
            lst.append(item)
        return lst
    
def recs_author(db, title):
    query = """
    MATCH (b:Book) 
    WHERE b.title =~ $input
    MATCH (b)-[:WRITTEN_IN]->(y:Year)
    MATCH (b)-[:WRITTEN_BY]->(a:Author)<-[:WRITTEN_BY]-(rec:Book)
    RETURN rec, a, y;
    """
    lst = []
    with db.session() as session:
        books = session.run(query, input=title)
        for item in books:
            lst.append(item)
        return lst
    
def recs_year(db, title):
    query = """
    MATCH (b:Book) 
    WHERE b.title =~ $input
    MATCH (b)-[:WRITTEN_BY]->(a:Author)
    MATCH (b)-[:WRITTEN_IN]->(y:Year)<-[:WRITTEN_IN]-(rec:Book)
    RETURN rec, y;
    """
    lst = []
    with db.session() as session:
        books = session.run(query, input=title)
        for item in books:
            lst.append(item)
        return lst
#endregion

#region GRPC Methods
class Neo4jServicer(book_pb2_grpc.BookServiceServicer):
    def __init__(self):
        uri = "bolt://localhost:7687"
        username = "neo4j"
        password = "123"
        self.driver = GraphDatabase.driver(os.getenv('NEOIP', uri), auth=(os.getenv('NEOUSER', username), os.getenv('NEOPASS', password)))
        
    def close(self):
        self.driver.close()
        
    def GetBookByTitle(self, request, context):
        book = get_book_by_title_with_more(self.driver, request.title)
        
        if book is None:
            return book_pb2.Book(id=0, name="", description="", author="", amount=0, year=0)
        else:
            print(book['a']["name"])
            return book_pb2.Book(id=book['b']["id"], name=book['b']["title"], description=book['b']["description"], author=book['a']['name'], amount=30, year=book['y']['name'])
        
    def GetBookSimpleByTitle(self, request, context):
        book = get_book_by_title(self.driver, request.title)
        
        if book is None:
            return book_pb2.Book(id=0, name="", amount=0)
        else:
            return book_pb2.Book(id=book["id"], name=book["title"], description=book["description"])
        
    def GetAllBooks(self, request, context):
        book = get_all_books(self.driver)
        
        if book is None:
            return book_pb2.BookList([])
        else:
            booksLst = book_pb2.BookList()
            for bk in book:
                bo = book_pb2.Book(id=bk['b']["id"], name=bk['b']["title"], description=bk['b']["description"], author=bk['a']['name'], amount=30, year=bk['y']['name'])
                booksLst.books.extend([bo])
            return booksLst
        
    def GetBooksBySearch(self, request, context):
        book = search_book(self.driver, request.title)
        
        if book is None:
            return book_pb2.BookList([])
        else:
            booksLst = book_pb2.BookList()
            for bk in book:
                bo = book_pb2.Book(id=bk['b']["id"], name=bk['b']["title"], description=bk['b']["description"], author=bk['a']['name'], amount=30, year=bk['y']['name'])
                booksLst.books.extend([bo])
            return booksLst
        
    def GetBookRecsAuthor(self, request, context):
        book = recs_author(self.driver, request.title)
        
        if book is None:
            return book_pb2.BookList([])
        else:
            booksLst = book_pb2.BookSimpleList()
            for bk in book:
                bo = book_pb2.BookSimple(id=bk['rec']["id"], name=bk['rec']["title"], author=bk['a']["name"], year=bk['y']["name"])
                booksLst.books.extend([bo])
            return booksLst
#endregion 
   
#region Caller method and Main method        
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    book_pb2_grpc.add_BookServiceServicer_to_server(
        Neo4jServicer(), server
    )
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()
    
if __name__ == '__main__':
    logging.basicConfig()
    serve()
#endregion