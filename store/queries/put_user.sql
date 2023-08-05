insert into users (email, name)
values (?, ?)
on conflict(email)
do update set name = excluded.name