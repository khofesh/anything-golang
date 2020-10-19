```
createuser -P -d gwp
createdb gwp
psql -U gwp -f setup.sql -d gwp
```
