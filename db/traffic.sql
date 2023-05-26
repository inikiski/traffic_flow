CREATE TABLE db_traffic ( TIME TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, entrance_flow INT NOT NULL, tunnel_flow INT NOT NULL, export_flow INT NOT NULL, direction BOOLEAN NOT NULL );
CREATE TABLE db_traffic_left ( TIME TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, entrance_flow INT NOT NULL, tunnel_flow INT NOT NULL, export_flow INT NOT NULL );
CREATE TABLE db_traffic_right ( TIME TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, entrance_flow INT NOT NULL, tunnel_flow INT NOT NULL, export_flow INT NOT NULL );
SELECT
	"db_traffic" COMMENT ON COLUMN db_traffic.direction IS 'false表示右,true表示左';
SELECT
	create_hypertable ( 'db_traffic', 'time' );
INSERT INTO "traffic"."db_traffic" ( entrance_flow, tunnel_flow, export_flow, direction )
VALUES
	( 23, 43, 56, TRUE );
INSERT INTO "traffic"."db_traffic" ( entrance_flow, tunnel_flow, export_flow, direction )
VALUES
	( 45, 33, 45, TRUE );
SELECT
	"sum" ( entrance_flow ) AS "entrance_flow",
	"sum" ( tunnel_flow ) AS "tunnel_flow",
	"sum" ( export_flow ) AS "export_flow",
	direction
FROM
	traffic.db_traffic
WHERE
	TIME > TIMESTAMP '2023-05-22 16:48:00'  - INTERVAL '1 minutes'
	AND direction = TRUE;

CREATE
MATERIALIZED VIEW db_traffic_left_daily WITH ( timescaledb.continuous ) AS
SELECT time_bucket(INTERVAL '1 day', time) AS bucket,
	   AVG(entrance_flow)                  AS avg_entrance_flow,
	   AVG(tunnel_flow)                    AS avg_tunnel_flow,
	   AVG(export_flow)                    AS avg_export_flow,
	   MAX(entrance_flow)                  AS max_entrance_flow,
	   MAX(tunnel_flow)                    AS max_tunnel_flow,
	   MAX(export_flow)                    AS max_export_flow,
	   MIN(entrance_flow)                  AS min_entrance_flow,
	   MIN(tunnel_flow)                    AS min_tunnel_flow,
	   MIN(export_flow)                    AS min_export_flow
FROM traffic.db_traffic_left
GROUP BY bucket;

CREATE
MATERIALIZED VIEW db_traffic_right_daily WITH ( timescaledb.continuous ) AS
SELECT time_bucket(INTERVAL '1 day', time) AS bucket,
	   AVG(entrance_flow)                  AS avg_entrance_flow,
	   AVG(tunnel_flow)                    AS avg_tunnel_flow,
	   AVG(export_flow)                    AS avg_export_flow,
	   MAX(entrance_flow)                  AS max_entrance_flow,
	   MAX(tunnel_flow)                    AS max_tunnel_flow,
	   MAX(export_flow)                    AS max_export_flow,
	   MIN(entrance_flow)                  AS min_entrance_flow,
	   MIN(tunnel_flow)                    AS min_tunnel_flow,
	   MIN(export_flow)                    AS min_export_flow
FROM traffic.db_traffic_left
GROUP BY bucket;