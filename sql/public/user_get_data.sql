-- FUNCTION: public.user_get_data(character varying, character varying)

-- DROP FUNCTION IF EXISTS public.user_get_data(character varying, character varying);

CREATE OR REPLACE FUNCTION public.user_get_data(
	_username character varying,
	_password character varying)
    RETURNS json
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(ug.s, '[]')
	FROM
    (
        SELECT json_agg(ag.*)::JSONB s
        FROM (
            SELECT u.id, u.uid, u.email, u.username, u.password
            FROM users u
            WHERE u.username = _username AND u.password = _password
        ) ag
    ) ug
    INTO _response;

    RETURN _response;
END;
$BODY$;

ALTER FUNCTION public.user_get_data(character varying, character varying)
    OWNER TO postgres;
