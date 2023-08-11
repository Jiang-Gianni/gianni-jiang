drop table if exists status;

create table status(
    id int primary key,
    description text not null
);

insert into status(id, description) values (1, 'Not Started'), (2, 'In Progress'), (3, 'On Hold'), (4, 'Completed');

-- name: GetAllStatus :many
select * from status;