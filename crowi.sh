#!/bin/sh
docker run -p 9200:9200 -e "http.host=0.0.0.0" -e "transport.host=127.0.0.1"  -e ES_JAVA_OPTS='-Xms64m -Xmx64m' -v /home/tesujiro/Crowi/elasticsearchdir:/usr/share/elasticsearch/data -d --name es sotarok/elasticsearch-kuromoji:2.4

#!/bin/sh
docker run --name mongo -v /home/tesujiro/Crowi/mongodir:/data/db -d mongo

#!/bin/sh
docker run --name redis -v /home/tesujiro/Crowi/redisdir:/data -d redis

#!/bin/sh
docker run --name crowi -p 8080:3000 -d \
-e MONGO_URI=mongodb://`sudo docker inspect --format="{{ .NetworkSettings.IPAddress }}" mongo`:27017/mongo \
-e REDIS_URI=redis://`sudo docker inspect --format="{{ .NetworkSettings.IPAddress }}" redis`:6379/redis \
-e ELASTICSEARCH_URI=http://`sudo docker inspect --format="{{ .NetworkSettings.IPAddress }}" es`:9200/es \
-e PASSWORD_SEED=tesujiro \
-e DEBUG='crowi:*' \
bakudankun/crowi

