Feature: Excercise Book CRUD

    Scenario: Add a Book to the library
        Given a library with no books
        When a book is added with title "Harry Potter" and author "J.K Rowling"
        Then the book can be retrieved

    Scenario: Update a Book in the library
        Given a library with book titled "Harry Potter" and author "J.K Rowling"
        When the book is updated with new title "Harry Potter and the chamber of secrets"
        Then the updated book can be retrieved
