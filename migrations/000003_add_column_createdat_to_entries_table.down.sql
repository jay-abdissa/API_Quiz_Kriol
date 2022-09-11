/*Filename: migrations/000002_add_column_createdat_to_entries_table.down.sql*/
ALTER TABLE entries
DROP COLUMN created_at;