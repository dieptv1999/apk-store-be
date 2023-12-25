alter table apk_version
    add self_click_download int default 0 null;

alter table apk
    add lang varchar(10) default 'vi' null;