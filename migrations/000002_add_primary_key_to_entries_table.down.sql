/*Filename: migrations/000002_add_primary_key_to_table_entries.down.sql*/
ALTER TABLE entries
DROP COLUMN id;
