-- FUNCTION: public.user_create(json, character varying)

-- DROP FUNCTION IF EXISTS public.user_create(json, character varying);

CREATE OR REPLACE FUNCTION public.user_create(
	js json,
	_uid character varying)
    RETURNS boolean
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
DECLARE
	_u users;
BEGIN
	SELECT *
	FROM users
	WHERE email = js->>'email'
	AND username = js->>'username'
	INTO _u;

	IF _u.id ISNULL THEN
		INSERT INTO users(
			uid,
			email,
			username,
			password
		) VALUES(
			_uid,
			js->>'email',
			js->>'username',
			js->>'password'
		);
		RETURN TRUE;
	ELSE
		RAISE EXCEPTION 'пользователь с таким именем уже существует';
	END IF;
END;
$BODY$;

ALTER FUNCTION public.user_create(json, character varying)
    OWNER TO postgres;
