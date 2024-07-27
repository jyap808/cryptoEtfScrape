## Crypto ETF scraper

This service/bot scrapes the Crypto ETF web sites and reports the information on the X account [@ubiqetfbot](https://twitter.com/ubiqetfbot).

This page outlines the methodology of this bot: [Crypto ETF flow bot - Methodology](https://julianyap.com/pages/2024-03-13-1710370430/)

This was historically [btcEtfScrape](https://github.com/jyap808/btcEtfScrape) but it was refactored (with some other clean ups and improvements) to be made more generic and to support multiple ETF asset types.

## Set up

This is largely a personal project so I just use a wrapper script. Set up the variables accordingly.

runme.sh
```
#!/bin/bash

## X
export GOTWI_API_KEY=
export GOTWI_API_KEY_SECRET=
export GOTWI_ACCESS_TOKEN=
export GOTWI_ACCESS_TOKEN_SECRET=

./cryptoEtfScrape -webhookURL https://discord.com/api/webhooks/[SET THIS]
```

## TODO

Further Dockerize the set up.

## License

[MIT License](LICENSE)
