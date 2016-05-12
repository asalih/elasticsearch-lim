# lim as an alternative for monitoring elasticsearch

lim logs node stats and index stats separately. You can see your index stats(search, log, fetch, etc.) within all nodes and keeps calculated datas.

## Index Statistics
![screen shot 2015-02-15 at 22 41 57](screenshots/index_stats.png)

## Node Statistics
![screen shot 2015-02-15 at 22 41 57](screenshots/node_stats.png)

# Installation

Download and build with go. When you build successfuly, just edit ".env" file and run.

- ELASTICSEARCH_SOURCE_URL = Gets the index stats node stats from this url
- ELASTICSEARCH_TARGET_URL = Posts the calculated url here. It uses the other elasticsearch server to keep data, like marvel.
- ELASTICSEARCH_INDEX_ST = Index address for keeping "index" stats.
- ELASTICSEARCH_INDEX_ND = Index address for keeping "node" stats.
- ELASTICSEARCH_TTL = Time to live stored index and node stats.
- INTERVAL_SECOND = Interval for read stats from source.
