# Technical Decisions

## SQLite

Vitalynq will use SQLite to store data locally.

Reasons:

- local storage suited to a single-user CLI application;
- a single file that is easy to back up;
- no server to administer;
- good transaction support.

The code will use `database/sql` to stay close to the standard library.

The planned SQLite driver is `modernc.org/sqlite` because it works without CGO and simplifies local installation.

This decision can be revisited if a concrete technical constraint appears.
