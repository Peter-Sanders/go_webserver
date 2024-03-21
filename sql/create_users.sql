create table if not exists users (
  id integer not null primary key autoincrement,
  FName text,
  LName text,
  Phone text,
  Email text,
  Time timestamp default current_timestamp,
  Password text,
  Username text
  );

