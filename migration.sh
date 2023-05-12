#!/bin/sh

dbstring=`env |grep -i dbstring | sed -e 's/MIGRATION_DBSTRING=//g' -e 's/"//g'`

migrate -path=migrations -database ${dbstring} $1
