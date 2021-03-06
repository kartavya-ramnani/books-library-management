### Requirements

#### Introduction

1. Any library member should be able to search books by their title, author, subject category as well by the publication date.

1. Each book will have a unique identification number and other details including a rack number which will help to physically locate the book.

1. There could be more than one copy of a book, and library members should be able to check-out and reserve any copy. We will call each copy of a book, a book item.

1. The system should be able to retrieve information like who took a particular book or what are the books checked-out by a specific library member.

1. There should be a maximum limit (5) on how many books a member can check-out.

1. There should be a maximum limit (10) on how many days a member can keep a book.

1. The system should be able to collect fines for books returned after the due date.

1. Members should be able to reserve books that are not currently available.

1. The system should be able to send notifications whenever the reserved books become available, as well as when the book is not returned within the due date.

1. Each book and member card will have a unique barcode. The system will be able to read barcodes from books and members’ library cards.

#### Actors/Domains

We have three main actors in our system:

- Librarian: Mainly responsible for adding and modifying books, book items, and users. The Librarian can also issue, reserve, and return book items.
- Member: All members can search the catalog, as well as check-out, reserve, renew, and return a book.
- System: Mainly responsible for sending notifications for overdue books, canceled reservations, etc.

#### Operations/Services

Here are the top use cases of the Library Management System:

1. Add/Remove/Edit book: To add, remove or modify a book or book item.
1. Search catalog: To search books by title, author, subject or publication date.
1. Register new account/cancel membership: To add a new member or cancel the membership of an existing member.
1. Check-out book: To borrow a book from the library.
1. Reserve book: To reserve a book which is not currently available.
1. Renew a book: To reborrow an already checked-out book.
1. Return a book: To return a book to the library which was issued to a member.