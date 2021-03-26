Feature: Excercise Book CRUD

    Scenario: Add a Book to the library
        Given a library with no books
        When a book is added with title "Harry Potter" and author "J.K Rowling"
        Then the book can be retrieved
