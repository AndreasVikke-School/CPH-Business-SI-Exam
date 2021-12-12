# pip install neo4j to install package 
from neo4j import GraphDatabase

class HelloWorldExample:

    def __init__(self, uri, user, password):
        self.driver = GraphDatabase.driver(uri, auth=(user, password))

    def close(self):
        self.driver.close()

    def print_greeting(self, message):
        with self.driver.session() as session:
            # greeting = session.write_transaction(self._create_and_return_greeting, message)
            greeting = session.write_transaction(self._create_books)
            print(message)

    @staticmethod
    def _create_and_return_greeting(tx, message):
        result = tx.run("CREATE (a:Greeting) "
                        "SET a.message = $message "
                        "RETURN a.message + ', from node ' + id(a)", message=message)
        return result.single()[0]
    
    @staticmethod
    def _create_books(tx):
        result = tx.run("""
                        LOAD CSV WITH HEADERS FROM 'file:///BookCSV.csv' as row FIELDTERMINATOR ';'
                        FOREACH (n IN (CASE WHEN row.description IS NULL THEN [] ELSE row.description END) |
                            MERGE(book:Book {title: row.title, description: row.description})

                        FOREACH (n IN (CASE WHEN row.authors IS NULL THEN [] ELSE row.authors END) |
                            merge(author:Author {name: row.authors})
                            MERGE (book)-[:WRITTEN_BY]->(author))

                        FOREACH (n IN (CASE WHEN row.published_year IS NULL THEN [] ELSE row.published_year END) |
                            MERGE (year:Year {name: toInteger(row.published_year)})
                            MERGE (book)-[:WRITTEN_IN]->(year))
                        )"""
        )


if __name__ == "__main__":
    greeter = HelloWorldExample("bolt://localhost:7687", "neo4j", "123")
    greeter.print_greeting("hello, world")
    greeter.close()