CREATE TABLE gologs (                         
    id serial not null primary key,           
    time Timestamptz not null default  now(), 
    record Jsonb not null2
);