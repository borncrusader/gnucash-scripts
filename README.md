## xtion-xfrms (Transaction Transforms)
Transaction Transforms is financial transactions database for categorizing
common transactions that occur on various credit card statements, bank account
statements etc. The database itself is in `sqlite`, however all the tooling
around it is in `golang`. The tooling comprises of a single binary, the `xtxf`.

If you are interested to contribute, please check the [CONTRIBUTING.md]() file.

## Why did you build this?
I've been managing my expenses, painstakingly if I may add, on Gnucash since
the Fall of 2011. All these years, it was a manual process, taking about 30
minutes per week. It also gave me a good opportunity to go over my expenses and
understand where I'm spending my money on.

## Why Gnucash? Why not Mint, Quicken etc.?
When I started out, I was a poor graduate student from India. Besides, being
the skeptic that I'm, I wasn't confident giving my credentials to Mint.

More importantly, life as a grad student involved dealing with a lot of IOUs
with friends and roommates. Most of that was tracked via Splitwise back in
those days but there was no way for me to keep track of it on my ledger using
Mint. Hence, I went ahead with Gnucash.

## Why Sqlite? Why not CSV (or other format)?


## Why golang? Why not rust (or some other language)?


## How does this work?
## Categories
The available categories are

### Expenses
1. Restaurants
2. Groceries
3. Pharmacy

