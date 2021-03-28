Feature: Excercise librarian scenarios

    @user
    Scenario: Add and Delete user to and from the library
        Given a library with no users
        And a "Librarian" librarian with username "Jason"
        When a user is added with username "Pete123" and name "Pete"
        Then the user can be retrieved

    @user @book
    Scenario: Checkout/in book from the library
        Given a library with user username "Peter123" and name "Pete"
        And a library with book "Harry Potter", author "J.K Rowling" and copies 5
        And a "Librarian" librarian with username "Jason"
        When the user checks out the book
        Then the book can be retrieved
        And the checked out copies of the book is 1
        When the user checks in the book
        Then the book can be retrieved
        And the checked out copies of the book is 0