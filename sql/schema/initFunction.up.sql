CREATE OR REPLACE FUNCTION public.admin_get_users(
	_id integer)
    RETURNS json
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
DECLARE
	_u users;
	_response JSONB;
BEGIN
	SELECT *
	FROM users
	WHERE id = _id
	INTO _u;

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователь с таким id не существует';
	END IF;

	IF _u.role != 'admin' THEN
		RAISE EXCEPTION 'пользователь не является администратором';
	END IF;

	SELECT
		COALESCE(agu.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM ( 
			SELECT u.id, u.username, u.name, u.surname, u.role, u.stretch_id, u.created
			FROM users u
		) ag
	) agu
	INTO _response;

	RETURN _response;
END;
$BODY$;