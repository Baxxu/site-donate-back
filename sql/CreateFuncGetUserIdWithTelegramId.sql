CREATE OR REPLACE FUNCTION public.get_user_id_with_telegram_id(IN var_telegram_id bigint)
    RETURNS bigint
    LANGUAGE 'plpgsql'
AS
$BODY$
declare
    var_user_id bigint;
begin
    select id into var_user_id from users where telegram_id = var_telegram_id;
    if var_user_id IS NULL then
        insert into users (telegram_id)
        values (var_telegram_id)
        RETURNING id Into var_user_id;
    end if;
    return var_user_id;
end
$BODY$;

ALTER FUNCTION public.get_user_id_with_telegram_id(bigint)
    OWNER TO test;