drop table if exists feedback;

drop sequence if exists feedback_seq;

create sequence feedback_seq;

create table feedback(
    id int primary key default nextval('feedback_seq'),
    project text not null,
    description text not null,
    created_on timestamp default now() not null
);

-- name: GetAllFeedbacks :many
select * from feedback;

-- name: CreateFeedback :one
insert into feedback (project, description) values ($1, $2) returning *;