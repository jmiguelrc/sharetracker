# Sharetracker

CLI util to track the current value of all the open positions in your share portefolio.

## Usage

```
$ sharetracker

Total buy value:  €4,709.95
Current Value:    €7,004.55
Total Net Value:  €6,362.06
Gross Profit:     €2,294.60
Net Profit:       €1,652.11
YTD Gross Profit: €994.60
```

## Setup

## Install

Install with: `go install github.com/jmiguelrc/sharetracker@latest`

Put all your open positions in $HOME/.sharetracker. Example:

```
positions:
- ticker: CSSPX.MI
  numShares: 10
  buyPrice: 200.0
  date: '25-01-2017'

- ticker: IWDA.AS
  numShares: 20
  buyPrice: 43.895
  date: '15-06-2017'

- ticker: IWDA.AS
  numShares: 20
  buyPrice: 59.35
  date: '02-01-2019'

- ticker: IWDA.AS
  numShares: 10
  buyPrice: 64.505
  date: '22-05-2020'
```

[Yahoo Finance](https://finance.yahoo.com/) is used to fetch the current market price of each position. Please ensure the `ticker` exists in [Yahoo Finance](https://finance.yahoo.com/).