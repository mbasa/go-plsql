#!/bin/sh
GO_DB="go"
PG_MOD="`pwd`/libGoPgFunc.so"
psql $GO_DB --set=MOD=\'$PG_MOD\' -f install.sql
