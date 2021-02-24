ur-last-fm
==========

little cli tools to push some from your last-fm account to postgresql


How to Use
------------
Get an api key from last-fm here https://www.last.fm/api/account/create

Create a ~/.ur-last-fm
```
~/.ur-last-fm.json
{
    "token": "aaaaaaaaaa",
    "secret": "aaaaaaa",
    "username": "username",
    "password": "password",
    "conn": "postgresql://lastfm:password@127.0.0.1:8081/lastfm?sslmode=disable",
    "scrap": "getUserInfo:username"
}
```

if you have docker install you can run a postgresql db with the scripts scripts/init_db_debug.sh

Run ur-last-fm to scrap recenttracks
```
ur-last-fm -scrap getAllRecentTracks:username
```

debug mode will print the json body and SQL query
```
ur-last-fm -scrap getAllRecentTracks:username -debug
```
