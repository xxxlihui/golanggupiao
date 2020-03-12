create table day_record
(
    day    int,
    code   int,
    high   float,
    low    float,
    close  float,
    amount bigint,
    vol    bigint,
    zt     bool,
    dt     bool,
    dm     bool,
    dr     bool,
    pb     bool,
    stop   bool,
    lb     int
);
create table day_stat
(
    day    int,
    code   int,
    high   float,
    low    float,
    close  float,
    amount bigint,
    vol    bigint,
    zt     int,
    dt     int,
    dm     int,
    dr     int,
    pb     int,
    stop   int
);
create table day_lb
(
    day  int,
    code int,
    lb   int
)