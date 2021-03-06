Feature: Excercise Senior librarian scenarios

    @book
    Scenario: Add number of copies of a new Book to the library
        Given a library with no books
        And a "Senior" librarian with username "Janice"
        When a book is added with title "Harry Potter", author "J.K Rowling" and copies 5
        Then the book can be retrieved
        And the book has 5 copies

    @book
    Scenario: Add number of copies of an existing Book to the library
        Given a library with book "Harry Potter", author "J.K Rowling" and copies 5
        And a "Senior" librarian with username "Janice"
        When 5 more copies of the book are added
        Then the book can be retrieved
        And the book has 10 copies

    @book
    Scenario: Remove number of copies of an existing Book from the library
        Given a library with book "Harry Potter", author "J.K Rowling" and copies 5
        And a "Senior" librarian with username "Janice"
        When 3 copies of the book are removed
        Then the book can be retrieved
        And the book has 2 copies