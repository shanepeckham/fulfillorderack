### FulfillOrder - TACK

A containerised Go swagger API to fulfill orders and commit them to MongoDB

The following environment variables need to be passed to the container:

### ACK Logging
```
ENV TEAMNAME=[YourTeamName]
```
### For Mongo
```
ENV MONGOURL="mongodb://[mongoinstance].[namespace]"
```
### File mount
```
/orders
```
