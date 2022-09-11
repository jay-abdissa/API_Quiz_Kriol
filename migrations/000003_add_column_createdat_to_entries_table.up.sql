/*Filename: migrations/000002_add_column_createdat_to_entries_table.up.sql*/
ALTER TABLE entries
ADD COLUMN created_at timestamp (0) with time zone NOT NULL DEFAULT NOW(); 