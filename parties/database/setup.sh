#!/bin/sh
set -e

/etc/init.d/postgresql start
psql -f create_fixture.sql
/etc/init.d/postgresql stop