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
	AND direction = TRUE