#!/bin/sh

# exit on error
set -e
COMMAND=$@

echo 'Waiting for database to be available...'
maxTries=10 # tries to connect to database
while [ "$maxTries" -gt 0 ] && ! mysql -h "$MYSQL_HOST" -P "$MYSQL_PORT" -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" "$MYSQL_DB" -e 'SHOW TABLES'; do # wait for database to be available
    maxTries=$(($maxTries - 1)) # decrement maxTries
    sleep 3 # wait 3 seconds
done
echo
if [ "$maxTries" -le 0 ]; then
    echo >&2 'error: unable to contact mysql after 10 tries'
    exit 1
fi

exec $COMMAND