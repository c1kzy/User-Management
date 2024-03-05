CREATE TABLE users
(
    id              bigserial PRIMARY KEY,
    first_name      varchar(255) not null,
    last_name       varchar(255) not null,
    email           varchar(255) not null UNIQUE,
    hashed_password varchar(255) not null,
    role            int          not null,
    created_at      timestamp    not null DEFAULT CURRENT_TIMESTAMP,
    updated_at      timestamp    not null DEFAULT CURRENT_TIMESTAMP,
    oID             UUID         not null UNIQUE,
    status          int          not null
);

CREATE TABLE posts
(
    id         bigserial PRIMARY KEY,
    user_id    bigint REFERENCES "users" (id),
    name       varchar(255) not null,
    text       varchar(1000),
    created_at timestamp    not null DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp    not null DEFAULT CURRENT_TIMESTAMP,
    oID        UUID         not null,
    vote_sum   int                   DEFAULT 0,
    status     int          not null
);

CREATE TABLE ratings
(
    id           bigserial PRIMARY KEY,
    from_user_id bigint REFERENCES "users" (id),
    to_post_id   bigint REFERENCES "posts" (id),
    user_vote    int       not null,
    when_voted   timestamp not null DEFAULT CURRENT_TIMESTAMP,
    status       int                DEFAULT 0

);

-- User Procedures
CREATE OR REPLACE PROCEDURE insertUser(
    u_first_name varchar(255),
    u_last_name varchar(255),
    u_email varchar(255),
    u_password varchar(255),
    u_role int,
    u_created_at timestamp,
    u_oid uuid,
    u_status int
)
AS
$$
BEGIN
    INSERT INTO users(first_name, last_name, email, hashed_password, role, created_at, oID, status)
    VALUES (u_first_name, u_last_name, u_email, u_password, u_role, u_created_at, u_oid, u_status);
END
$$ LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE updateUser(
    u_first_name varchar(255),
    u_last_name varchar(255),
    u_email varchar(255),
    u_password varchar(255),
    u_role int,
    u_updated_at timestamp,
    u_status int,
    u_id int
)
AS
$$
BEGIN
    UPDATE users
    SET first_name=u_first_name,
        last_name=u_last_name,
        email=u_email,
        hashed_password=u_password,
        role=u_role,
        updated_at=u_updated_at,
        status=u_status
    WHERE id = u_id;
END
$$ LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE deleteUser(
    u_status int,
    u_id int
)
AS
$$
BEGIN
    UPDATE users SET status=u_status WHERE id = u_id;
END
$$ LANGUAGE plpgsql;

-- Post Procedures

CREATE OR REPLACE PROCEDURE insertPost(
    p_user_id int,
    p_name varchar(255),
    p_text varchar(1000),
    p_created_at timestamp,
    p_oid uuid,
    p_status int
)
AS
$$
BEGIN
    INSERT INTO posts(user_id, name, text, created_at, oID, status)
    VALUES (p_user_id, p_name, p_text, p_created_at, p_oid, p_status);
END
$$ LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE updatePost(
    p_updated_at timestamp,
    p_text varchar(1000),
    p_status int,
    p_id int
)
AS
$$
BEGIN
    UPDATE posts SET updated_at=p_updated_at, text=p_text, status=p_status WHERE id = p_id;
END
$$ LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE deletePost(
    p_status int,
    p_id int
)
AS
$$
BEGIN
    UPDATE posts SET status=p_status WHERE id = p_id;
END
$$ LANGUAGE plpgsql;

--- Vote Procedures

CREATE OR REPLACE PROCEDURE insertVote(
    v_from_user_id int,
    v_to_post_id int,
    v_user_vote int,
    v_when_voted timestamp,
    v_status int
)
AS
$$
BEGIN
    INSERT INTO ratings(from_user_id, to_post_id, user_vote, when_voted, status)
    VALUES (v_from_user_id, v_to_post_id, v_user_vote, v_when_voted, v_status);

END
$$ LANGUAGE plpgsql;


CREATE OR REPLACE PROCEDURE updateVote(
    v_user_vote int,
    v_when_voted timestamp,
    v_status int,
    v_id int
)
AS
$$
BEGIN
    UPDATE ratings
    SET user_vote  = v_user_vote,
        when_voted = v_when_voted,
        status     = v_status
    WHERE id = v_id;

END
$$ LANGUAGE plpgsql;


CREATE OR REPLACE PROCEDURE voteSumUpdater(
    v_post_id int,
    v_user_vote int
)
AS
$$
BEGIN
    UPDATE posts
    SET vote_sum = vote_sum + v_user_vote
    WHERE posts.id = v_post_id;

END
$$ LANGUAGE plpgsql;

