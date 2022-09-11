/*Filename: migrations/000001_create_entries_table.up.sql*/
CREATE TABLE IF NOT EXISTS entries(
    kriol text NOT NULL,
    english text NOT NULL
);