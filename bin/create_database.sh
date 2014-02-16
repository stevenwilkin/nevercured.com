#!/bin/bash

for FILE in db/*.sql
do
	cat $FILE | sqlite3 db/nevercured.db
done
