create table ads
(
    id        int auto_increment
        primary key,
    cover     varchar(200) not null,
    title     varchar(200) not null,
    create_dt datetime     not null,
    click     int          not null,
    link      varchar(200) not null,
    active    varchar(10)  not null,
    count     int          not null
)
    collate = utf8mb3_bin;

create table apk
(
    id                       bigint unsigned auto_increment
        primary key,
    title                    varchar(255)                null,
    description              text                        null,
    descriptionHTML          text                        null,
    summary                  text                        null,
    installs                 text                        null,
    minInstalls              int                         null,
    realInstalls             int                         null,
    score                    double                      null,
    ratings                  int                         null,
    reviews                  int                         null,
    histogram                text charset latin1         null,
    price                    int                         null,
    free                     tinyint(1)                  null,
    currency                 varchar(255) charset latin1 null,
    sale                     tinyint(1)                  null,
    saleTime                 varchar(255) charset latin1 null,
    originalPrice            int                         null,
    inAppProductPrice        varchar(255)                null,
    developer                varchar(255)                null,
    developerId              varchar(255)                null,
    developerEmail           varchar(255)                null,
    developerWebsite         text                        null,
    developerAddress         text                        null,
    privacyPolicy            text                        null,
    genre                    text                        null,
    genreId                  text                        null,
    icon                     text                        null,
    headerImage              text                        null,
    screenshots              text                        null,
    video                    text                        null,
    videoImage               text                        null,
    contentRating            text                        null,
    contentRatingDescription text                        null,
    adSupported              tinyint(1)                  null,
    containsAds              tinyint(1)                  null,
    released                 varchar(255)                null,
    updated                  timestamp                   null,
    version                  varchar(255)                null,
    appId                    varchar(255)                null,
    url                      text                        null,
    categories               text                        null,
    isGame                   tinyint(1) default 0        null,
    isHot                    tinyint(1) default 0        null,
    constraint apk_appId_uindex
        unique (appId)
)
    collate = utf8mb4_unicode_ci;

create table apk_version
(
    versionId       bigint unsigned auto_increment
        primary key,
    appId           varchar(255)       not null,
    requiresAndroid varchar(255)       null,
    architecture    varchar(255)       null,
    signature       varchar(255)       null,
    permissions     text               null,
    version         varchar(255)       null,
    versionCode     varchar(255)       null,
    downloadLink    text               null,
    apkPureLink     varchar(255)       null,
    amountStorage   varchar(255)       null,
    apkType         varchar(255)       null,
    updateOn        varchar(255)       null,
    createAt        datetime           null,
    is_downloaded   smallint default 0 null comment '0 - link của apkpure 1 -link của app store '
)
    collate = utf8mb4_unicode_ci;

create table category
(
    id       bigint unsigned auto_increment
        primary key,
    store_id varchar(255)       null,
    name     varchar(255)       null,
    status   smallint default 0 null,
    icon     text               null,
    slug     varchar(255)       null,
    is_game  tinyint  default 0 null
)
    collate = utf8mb4_unicode_ci;

create definer = review@localhost trigger upd_check
    before update
                      on category
                      for each row
BEGIN
           set NEW.slug = REPLACE(LOWER(NEW.name), ' ', '-');
END;

create definer = review@localhost trigger upd_check_2
    before insert
    on category
    for each row
BEGIN
           set NEW.slug = REPLACE(LOWER(NEW.name), ' ', '-');
END;

create table comment
(
    id      bigint unsigned auto_increment
        primary key,
    content text         null,
    apk_id  varchar(255) null
)
    collate = utf8mb4_unicode_ci;

create table feed
(
    id        int          not null
        primary key,
    cover     varchar(200) not null,
    title     varchar(200) not null,
    create_dt datetime     not null,
    click     int          not null,
    link      varchar(200) not null,
    active    varchar(10)  not null,
    count     int          not null
)
    collate = utf8mb3_bin;

create table hashing_algorithms
(
    hashing_algorithms_id bigint unsigned auto_increment
        primary key,
    algorithm_name        varchar(64) null
)
    collate = utf8mb4_unicode_ci;

create table package
(
    package_id       varchar(255)       not null
        primary key,
    status           smallint default 0 null,
    is_crawl_version int      default 0 null
);

create table review
(
    id                   bigint unsigned auto_increment
        primary key,
    reviewId             text charset latin1        null,
    userName             text                       null,
    userImage            text                       null,
    content              text                       null,
    score                int                        null,
    thumbsUpCount        int                        null,
    reviewCreatedVersion varchar(32)                null,
    at                   varchar(255)               null,
    replyContent         text                       null,
    repliedAt            varchar(255)               null,
    appVersion           varchar(32) charset latin1 null,
    appId                varchar(255)               null
)
    collate = utf8mb4_unicode_ci;

create index review_appId_index
    on review (appId);

create table user
(
    id            bigint unsigned auto_increment
        primary key,
    user_name     varchar(100)  not null,
    user_code     varchar(50)   not null,
    date_of_birth date          null,
    email         varchar(500)  null,
    full_name     varchar(500)  null,
    first_name    varchar(500)  null,
    last_name     varchar(500)  null,
    avatar        varchar(1000) null,
    gender        int           null,
    phone_number  varchar(11)   not null,
    type          varchar(50)   not null,
    status        int           not null,
    two_mfa       int           null,
    partner       varchar(200)  null,
    created_at    timestamp     not null,
    updated_at    timestamp     not null,
    updated_by    varchar(50)   null,
    version       varchar(20)   null,
    constraint user_name
        unique (user_name)
)
    collate = utf8mb4_unicode_ci;

create table user_login_data
(
    user_id                    bigint auto_increment
        primary key,
    password_hash              varchar(255) null,
    password_salt              varchar(255) null,
    hashing_algorithms_id      int          null,
    confirmation_token         varchar(255) null,
    token_generation_time      timestamp    null,
    password_recovery_token    varchar(255) null,
    recovery_token_time        timestamp    null,
    email_validation_status_id int          null,
    created_at                 timestamp    not null,
    updated_at                 timestamp    not null
)
    collate = utf8mb4_unicode_ci;

create
definer = review@localhost function remove_accents(str text) returns text deterministic sql security invoker no sql
BEGIN

    SET str = REPLACE(str, 'Š', 'S');
    SET str = REPLACE(str, 'š', 's');
    SET str = REPLACE(str, 'Ð', 'Dj');
    SET str = REPLACE(str, 'Ž', 'Z');
    SET str = REPLACE(str, 'ž', 'z');
    SET str = REPLACE(str, 'À', 'A');
    SET str = REPLACE(str, 'Á', 'A');
    SET str = REPLACE(str, 'Ả', 'A');
    SET str = REPLACE(str, 'Ạ', 'A');
    SET str = REPLACE(str, 'Ắ', 'A');
    SET str = REPLACE(str, 'Ằ', 'A');
    SET str = REPLACE(str, 'Ặ', 'A');
    SET str = REPLACE(str, 'Ẳ', 'A');
    SET str = REPLACE(str, 'Â', 'A');
    SET str = REPLACE(str, 'Ầ', 'A');
    SET str = REPLACE(str, 'Ấ', 'A');
    SET str = REPLACE(str, 'Ã', 'A');
    SET str = REPLACE(str, 'Ä', 'A');
    SET str = REPLACE(str, 'Å', 'A');
    SET str = REPLACE(str, 'Æ', 'A');
    SET str = REPLACE(str, 'Ç', 'C');
    SET str = REPLACE(str, 'È', 'E');
    SET str = REPLACE(str, 'É', 'E');
    SET str = REPLACE(str, 'Ẹ', 'E');
    SET str = REPLACE(str, 'Ê', 'E');
    SET str = REPLACE(str, 'Ệ', 'E');
    SET str = REPLACE(str, 'Ề', 'E');
    SET str = REPLACE(str, 'Ế', 'E');
    SET str = REPLACE(str, 'Ë', 'E');
    SET str = REPLACE(str, 'Ì', 'I');
    SET str = REPLACE(str, 'Í', 'I');
    SET str = REPLACE(str, 'Ị', 'I');
    SET str = REPLACE(str, 'Î', 'I');
    SET str = REPLACE(str, 'Ï', 'I');
    SET str = REPLACE(str, 'Ñ', 'N');
    SET str = REPLACE(str, 'Ò', 'O');
    SET str = REPLACE(str, 'Ó', 'O');
    SET str = REPLACE(str, 'Ọ', 'O');
    SET str = REPLACE(str, 'Ô', 'O');
    SET str = REPLACE(str, 'Ộ', 'O');
    SET str = REPLACE(str, 'Ố', 'O');
    SET str = REPLACE(str, 'Ồ', 'O');
    SET str = REPLACE(str, 'Õ', 'O');
    SET str = REPLACE(str, 'Ơ', 'O');
    SET str = REPLACE(str, 'Ờ', 'O');
    SET str = REPLACE(str, 'Ớ', 'O');
    SET str = REPLACE(str, 'Ợ', 'O');
    SET str = REPLACE(str, 'Ö', 'O');
    SET str = REPLACE(str, 'Ø', 'O');
    SET str = REPLACE(str, 'Ủ', 'U');
    SET str = REPLACE(str, 'Ù', 'U');
    SET str = REPLACE(str, 'Ú', 'U');
    SET str = REPLACE(str, 'Ụ', 'U');
    SET str = REPLACE(str, 'Ư', 'U');
    SET str = REPLACE(str, 'Û', 'U');
    SET str = REPLACE(str, 'Ü', 'U');
    SET str = REPLACE(str, 'Ý', 'Y');
    SET str = REPLACE(str, 'Þ', 'B');
    SET str = REPLACE(str, 'Đ', 'D');
    SET str = REPLACE(str, 'ß', 'Ss');
    SET str = REPLACE(str, 'à', 'a');
    SET str = REPLACE(str, 'ả', 'a');
    SET str = REPLACE(str, 'á', 'a');
    SET str = REPLACE(str, 'ạ', 'a');
    SET str = REPLACE(str, 'ắ', 'a');
    SET str = REPLACE(str, 'ằ', 'a');
    SET str = REPLACE(str, 'ặ', 'a');
    SET str = REPLACE(str, 'ẳ', 'a');
    SET str = REPLACE(str, 'â', 'a');
    SET str = REPLACE(str, 'ấ', 'a');
    SET str = REPLACE(str, 'ầ', 'a');
    SET str = REPLACE(str, 'ã', 'a');
    SET str = REPLACE(str, 'ä', 'a');
    SET str = REPLACE(str, 'å', 'a');
    SET str = REPLACE(str, 'æ', 'a');
    SET str = REPLACE(str, 'ç', 'c');
    SET str = REPLACE(str, 'è', 'e');
    SET str = REPLACE(str, 'é', 'e');
    SET str = REPLACE(str, 'ẹ', 'e');
    SET str = REPLACE(str, 'ê', 'e');
    SET str = REPLACE(str, 'ệ', 'e');
    SET str = REPLACE(str, 'ề', 'e');
    SET str = REPLACE(str, 'ế', 'e');
    SET str = REPLACE(str, 'ë', 'e');
    SET str = REPLACE(str, 'ì', 'i');
    SET str = REPLACE(str, 'í', 'i');
    SET str = REPLACE(str, 'ị', 'i');
    SET str = REPLACE(str, 'î', 'i');
    SET str = REPLACE(str, 'ï', 'i');
    SET str = REPLACE(str, 'ð', 'o');
    SET str = REPLACE(str, 'ñ', 'n');
    SET str = REPLACE(str, 'ò', 'o');
    SET str = REPLACE(str, 'ó', 'o');
    SET str = REPLACE(str, 'ơ', 'o');
    SET str = REPLACE(str, 'ờ', 'o');
    SET str = REPLACE(str, 'ớ', 'o');
    SET str = REPLACE(str, 'ợ', 'o');
    SET str = REPLACE(str, 'ọ', 'o');
    SET str = REPLACE(str, 'ô', 'o');
    SET str = REPLACE(str, 'ộ', 'o');
    SET str = REPLACE(str, 'ố', 'o');
    SET str = REPLACE(str, 'ồ', 'o');
    SET str = REPLACE(str, 'õ', 'o');
    SET str = REPLACE(str, 'ö', 'o');
    SET str = REPLACE(str, 'ø', 'o');
    SET str = REPLACE(str, 'ủ', 'u');
    SET str = REPLACE(str, 'ù', 'u');
    SET str = REPLACE(str, 'ú', 'u');
    SET str = REPLACE(str, 'ụ', 'u');
    SET str = REPLACE(str, 'ư', 'u');
    SET str = REPLACE(str, 'û', 'u');
    SET str = REPLACE(str, 'ý', 'y');
    SET str = REPLACE(str, 'ý', 'y');
    SET str = REPLACE(str, 'þ', 'b');
    SET str = REPLACE(str, 'ÿ', 'y');
    SET str = REPLACE(str, 'ƒ', 'f');
    SET str = REPLACE(str, 'đ', 'd');


RETURN str;
END;

