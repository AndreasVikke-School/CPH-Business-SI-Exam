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

#region Cypher query methods
def detach_delete(db):
    query = """MATCH (all) DETACH DELETE all;"""
    with db.session() as session:
        session.run(query)
        return "Deleted all nodes and edges"
    
def write_csv_to_db(db):
    constraintQuery= """CREATE CONSTRAINT IF NOT EXISTS FOR (book:Book) REQUIRE book.isbn10 IS UNIQUE;"""
    writeQuery = """
        LOAD CSV WITH HEADERS FROM 'https://docs.google.com/spreadsheets/d/e/2PACX-1vQWglf2y9ogpu19WLA1FcLUmOES-Iham-DzQfZNqwqUj4psSbhl2VocVqQfEls8kw/pub?output=csv' as row 
            FOREACH (n IN (CASE WHEN row.description IS NULL THEN [] ELSE row.description END) |
                MERGE(book:Book {isbn10: row.isbn10, title: row.title, description: row.description, amount: toInteger(row.amount)})
                FOREACH (n IN (CASE WHEN row.genre IS NULL THEN [] ELSE row.authors END) |
                    MERGE(genre:Genre {name: row.genre})
                    MERGE (book)-[:GENRE_IS]->(genre))
                
                FOREACH (n IN (CASE WHEN row.authors IS NULL THEN [] ELSE row.authors END) |
                    MERGE(author:Author {name: row.authors})
                    MERGE (book)-[:WRITTEN_BY]->(author))
                    
                FOREACH (n IN (CASE WHEN row.genre IS NULL THEN [] ELSE row.authors END) |
                    MERGE(genre:Genre {name: row.genre})
                    MERGE (book)-[:GENRE_IS]->(genre))

                FOREACH (n IN (CASE WHEN row.published_year IS NULL THEN [] ELSE row.published_year END) |
                    MERGE (year:Year {name: toInteger(row.published_year)})
                    MERGE (book)-[:WRITTEN_IN]->(year))
            );
    """
    with db.session() as session:
        session.run(constraintQuery)
        session.run(writeQuery)
        return "Succes"
    
def get_book_by_title(db, title):
    query = """
    match (b:Book {title: $input})
    MATCH (b)-[:WRITTEN_BY]->(a:Author)
    RETURN b, a LIMIT 1;"""
    with db.session() as session:
        book = session.run(query, input=title)
        return book.single()

def get_book_by_title_with_relationships(db, title):
    query = """
    MATCH (b:Book {title: $input})
    MATCH (b)-[:WRITTEN_BY]->(a:Author)
    MATCH (b)-[:WRITTEN_IN]->(y:Year)
    MATCH (b)-[:GENRE_IS]->(g:Genre)
    RETURN b, a, y, g LIMIT 1;
    """
    with db.session() as session:
        book = session.run(query, input=title)
        return book.single()

def get_all_books(db):
    query = """
    MATCH (b:Book)
    MATCH (b)-[:WRITTEN_BY]->(a:Author)
    MATCH (b)-[:WRITTEN_IN]->(y:Year)
    MATCH (b)-[:GENRE_IS]->(g:Genre)
    RETURN b, a, y, g;
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
    MATCH (b)-[:GENRE_IS]->(g:Genre)
    RETURN b, a, y, g;
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
        
    def WriteCsvToDb(self, request, context):
        detach_delete(self.driver)
        msg = write_csv_to_db(self.driver)
        return book_pb2.BookTitle(title=msg)
        
    def GetBookByTitle(self, request, context):
        book = get_book_by_title_with_relationships(self.driver, request.title)
        
        if book is None:
            return book_pb2.Book(isbn=0, name="", description="", author="", amount=0, year=0, genre="")
        else:
            return book_pb2.Book(isbn=book['b']["isbn10"], name=book['b']["title"], description=book['b']["description"], author=book['a']['name'], amount=book['b']['amount'], year=book['y']['name'], genre=book["g"]["name"])

    def GetBookSimpleByTitle(self, request, context):
        book = get_book_by_title(self.driver, request.title)
        
        if book is None:
            return book_pb2.BookSimple(isbn=0, name="", author="")
        else:
            return book_pb2.BookSimple(isbn=book['b']["isbn10"], name=book['b']["title"], author=book['a']["name"])
        
    def GetAllBooks(self, request, context):
        book = get_all_books(self.driver)
        
        if book is None:
            return book_pb2.BookList([])
        else:
            booksLst = book_pb2.BookList()
            for bk in book:
                bo = book_pb2.Book(isbn=bk['b']["isbn10"], name=bk['b']["title"], description=bk['b']["description"], author=bk['a']['name'], amount=bk['b']['amount'], year=bk['y']['name'], genre=bk['g']['name'])
                booksLst.books.extend([bo])
            return booksLst
        
    def GetBooksBySearch(self, request, context):
        book = search_book(self.driver, request.title)
        
        if book is None:
            return book_pb2.BookList([])
        else:
            booksLst = book_pb2.BookList()
            for bk in book:
                bo = book_pb2.Book(isbn=bk['b']["isbn10"], name=bk['b']["title"], description=bk['b']["description"], author=bk['a']['name'], amount=bk['b']['amount'], year=bk['y']['name'], genre=bk['g']['name'])
                booksLst.books.extend([bo])
            return booksLst
        
    def GetBookRecsAuthor(self, request, context):
        book = recs_author(self.driver, request.title)
        
        if book is None:
            return book_pb2.BookList([])
        else:
            booksLst = book_pb2.BookSimpleList()
            for bk in book:
                bo = book_pb2.BookSimple(isbn=bk['rec']["isbn10"], name=bk['rec']["title"], author=bk['a']["name"], year=bk['y']["name"])
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