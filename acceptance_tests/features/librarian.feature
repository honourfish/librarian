Feature: Excercise librarian scenarios

    @user
    Scenario: Add and Delete user to and from the library
        Given a library with no users
        And a "Librarian" librarian with username "Jason"
        When a user is added with username "Pete123" and name "Pete"
        Then the user can be retrieved