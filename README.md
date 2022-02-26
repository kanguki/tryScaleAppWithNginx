docker-compose up  -d --scale api=3

#when you want to clear cache and rebuild with new config
docker-compose up --scale api=5 --no-deps --build

#caveat
I haven't been able to invalidate cache if I scale from 5 -> 3 (the embedded dns resolver still returns a set of 5 hosts, while only 3 are running. might look into that later)
