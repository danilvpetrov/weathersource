DROP SCHEMA IF EXISTS weathersource;
CREATE SCHEMA weathersource;
USE weathersource;
GRANT SELECT,INSERT,UPDATE,DELETE ON weathersource.* TO 'weathersource'@'%' IDENTIFIED BY 'weathersource-password';

--
-- weather_data table stores meteological data (actual and forecast estimates).
--
CREATE TABLE IF NOT EXISTS weather_data (
    data_time                       INTEGER NOT NULL,
    data_type                       VARCHAR(255),
    timezone                        VARCHAR(255),
    latitude                        DECIMAL (10,4),
    longitude                       DECIMAL (10,4),
    apparent_temperature_high       DECIMAL (10,4),
    apparent_temperature_high_time  BIGINT,
    apparent_temperature_low        DECIMAL (10,4),
    apparent_temperature_low_time   BIGINT,
    apparent_temperature_max        DECIMAL (10,4),
    apparent_temperature_max_time   BIGINT,
    apparent_temperature_min        DECIMAL (10,4),
    apparent_temperature_min_time   BIGINT,
    temperature                     DECIMAL (10,4),
    apparent_temperature            DECIMAL (10,4),
    cloud_cover                     DECIMAL (10,4),
    dew_point                       DECIMAL (10,4),
    humidity                        DECIMAL (10,4),
    icon                            VARCHAR(255),
    moon_phase                      DECIMAL (10,4),
    ozone                           DECIMAL (10,4),
    precip_intensity                DECIMAL (10,4),
    precip_intensity_max            DECIMAL (10,4),
    precip_intensity_max_time       BIGINT,
    precip_probability              DECIMAL (10,4),
    precip_type                     VARCHAR(255),
    pressure                        DECIMAL (10,4),
    summary                         VARCHAR(255),
    sunrise_time                    BIGINT,
    sunset_time                     BIGINT,
    temperature_high                DECIMAL (10,4),
    temperature_high_time           BIGINT,
    temperature_low                 DECIMAL (10,4),
    temperature_low_time            BIGINT,
    temperature_max                 DECIMAL (10,4),
    temperature_max_time            BIGINT,
    temperature_min                 DECIMAL (10,4),
    temperature_min_time            BIGINT,
    uv_index                        INT,
    uv_index_time                   BIGINT,
    visibility                      DECIMAL (10,4),
    wind_bearing                    INT,
    wind_gust                       DECIMAL (10,4),
    wind_gust_time                  BIGINT,
    wind_speed                      DECIMAL (10,4),

    CONSTRAINT pk_weather_data PRIMARY KEY (
        data_time,
        data_type,
        timezone,
        latitude,
        longitude
    )
) ROW_FORMAT=COMPRESSED;


FLUSH PRIVILEGES;
