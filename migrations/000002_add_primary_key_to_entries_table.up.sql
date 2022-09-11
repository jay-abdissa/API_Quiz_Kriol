/*Filename: migrations/000002_add_primary_key_to_table_entries.up.sql*/
ALTER TABLE entries
ADD COLUMN id serial NOT NULL PRIMARY KEY;