Feature: Excercise Senior engineer scenarios

    Scenario: Add a Book to the library
        Given a library with no books
        And a "Senior" librarian with username "Janice"
        When a book is added with title "Harry Potter", author "J.K Rowling" and copies 5
        Then the book can be retrieved
        And the book has 5 copies
