CREATE or replace FUNCTION public.refresh_token_validate(IN var_token bytea, IN var_time_now bigint,
                                                         IN var_time_to_compare bigint)
    RETURNS record
    LANGUAGE 'plpgsql'
AS
$BODY$
declare
    var_session_id       bytea;
    var_user_id          bigint;
    var_private_key      bytea;
    var_last_access_time bigint;
    return_record        record;
begin
    return_record = (null::bytea, null::bigint, null::bytea, false);
    select id, user_id, private_key, last_access_time
    into var_session_id, var_user_id, var_private_key, var_last_access_time
    from sessions
    where refresh_token = var_token;
    if var_session_id is null then
        return return_record;
    end if;
    if var_last_access_time < var_time_to_compare then
        delete from sessions where refresh_token = var_token;
        return return_record;
    end if;
    update sessions set last_access_time = var_time_now where refresh_token = var_token;
    return_record = (var_session_id, var_user_id, var_private_key, true);
    return return_record;
end;
$BODY$;

ALTER FUNCTION public.refresh_token_validate(bytea, bigint, bigint)
    OWNER TO test;