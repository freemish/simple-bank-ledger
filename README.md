# simple-bank-ledger

*Very* simple bank ledger application in Go. Not finished, at all.

The goal is to create an interface that:

- registers new users
- logs in users
- shows logged-in user's balance
- shows logged-in user's transaction history
- records a transaction for a logged-in user
- logs out logged-in user
- just uses local cache on a single client machine (doesn't handle concurrency, isn't persistent, does not check against a database)
- could easily be extended to use a database, etc.
