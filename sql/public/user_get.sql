-- FUNCTION: public.user_get(integer)

-- DROP FUNCTION IF EXISTS public.user_get(integer);

CREATE OR REPLACE FUNCTION public.user_get(
	_idu integer)
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
	WHERE id = _idu
	INTO _u;

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователя с таким id не существует';
	END IF;

	SELECT
		COALESCE(ug.s, '[]')
	FROM
    (
        SELECT json_agg(ag.*)::JSONB s
        FROM (
            SELECT u.id, u.uid, u.email, u.username, u.password, u.created_at, u.updated_at
            FROM users u
            WHERE u.id = _idu
        ) ag
    ) ug
    INTO _response;

    RETURN _response;
END;
$BODY$;

ALTER FUNCTION public.user_get(integer)
    OWNER TO postgres;
