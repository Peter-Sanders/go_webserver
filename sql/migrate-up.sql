create table if not exists user (
  id integer not null primary key autoincrement,
  fname text,
  lname text,
  phone text,
  email text,
  time timestamp default current_timestamp,
  password text,
  username text
  );


create table if not exists content_type (
  id integer not null primary key autoincrement,
  typ_cd text,
  cre_ts timestamp default current_timestamp
);


create table if not exists content (
  id integer not null primary key autoincrement,
  user_id integer not null,
  content_type_id integer not null,
  content text,
  images text,
  time timestamp default current_timestamp,
  foreign key (user_id) references user(id),
  foreign key (content_type_id) references content_type(id)
  );

