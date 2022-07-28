# simple-bank-ledger

*Very* simple bank ledger application in Go. Not finished, at all.

The goal is to create an interface that:

- [x] registers new users
- [x] logs in users
- [x] shows logged-in user's balance
- [x] shows logged-in user's transaction history
- [x] records a transaction for a logged-in user
- [x] logs out logged-in user
- [x] just uses local cache on a single client machine (doesn't handle concurrency, isn't persistent, does not check against a database)
- [x] could easily be extended to use a database, etc.
